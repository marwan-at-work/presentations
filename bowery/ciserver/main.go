package main

import (
	"context"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"

	"golang.org/x/oauth2"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/pkg/jsonmessage"
	"github.com/docker/docker/pkg/term"
	"github.com/pkg/errors"

	docker "github.com/docker/docker/client"

	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"

	"github.com/docker/docker/pkg/archive"
	"github.com/google/go-github/github"
)

var logs = map[string][]string{}

type logger struct {
	id  string
	url string
}

func (l *logger) Write(p []byte) (int, error) {
	logs[l.id] = append(logs[l.id], string(p))

	fmt.Printf("%s\n", p)

	return len(p), nil
}

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/webhook", webhookHandler)
	http.HandleFunc("/logs", logHandler)

	fmt.Println("👂", " on 3000")
	http.ListenAndServe(":3000", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello!")
}

func webhookHandler(w http.ResponseWriter, r *http.Request) {
	payload, err := github.ValidatePayload(r, []byte("supersecret"))
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	event, err := github.ParseWebHook(github.WebHookType(r), payload)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	switch pre := event.(type) {
	case *github.PullRequestEvent:
		action := pre.GetAction()
		if action == "opened" || action == "reopened" || action == "synchronize" {
			id := pre.PullRequest.Head.GetSHA()
			link := os.Getenv("NGROK_URL")
			if link == "" {
				link = "https://50a593ad.ngrok.io"
			}

			url := fmt.Sprintf("%s/logs?id=%s", link, id)
			l := &logger{id, url}
			fmt.Println("building pr")
			err = buildPR(pre, l)
			status := "success"
			if err != nil {
				status = "error"
			}

			reportStatus(status, pre, url)
		}
	default:
		fmt.Fprintf(w, "Uninteresting event: %T", event)
	}
}

func buildPR(pre *github.PullRequestEvent, l *logger) error {
	status := "pending"
	reportStatus(status, pre, l.url)

	pathToRepo, err := cloneRepo(pre, l)
	if err != nil {
		return err
	}
	defer os.RemoveAll(pathToRepo)

	err = buildImage(pathToRepo, l)

	return err
}

func cloneRepo(pre *github.PullRequestEvent, l *logger) (string, error) {
	repoName := pre.PullRequest.Head.Repo.GetName()
	tempDir, err := ioutil.TempDir("", repoName)
	if err != nil {
		return "", err
	}

	cloneURL := pre.PullRequest.Head.Repo.GetCloneURL()
	branchName := pre.PullRequest.Head.GetRef()

	repo, err := git.PlainClone(tempDir, false, &git.CloneOptions{
		URL:               cloneURL,
		Progress:          l,
		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
		ReferenceName:     plumbing.ReferenceName(fmt.Sprintf("refs/heads/%v", branchName)),
	})
	if err != nil {
		return "", err
	}

	wt, err := repo.Worktree()
	if err != nil {
		return "", err
	}

	commitSha := pre.PullRequest.Head.GetSHA()
	err = wt.Checkout(&git.CheckoutOptions{
		Hash: plumbing.NewHash(commitSha),
	})
	if err != nil {
		return "", err
	}

	return tempDir, nil
}

func buildImage(repoDir string, l *logger) error {
	c, err := docker.NewEnvClient()
	if err != nil {
		return err
	}

	buildCtx, err := archive.TarWithOptions(repoDir, &archive.TarOptions{
		Compression: archive.Uncompressed,
	})
	if err != nil {
		return errors.Wrap(err, "could not create build context")
	}
	defer buildCtx.Close()

	ctx := context.Background()
	resp, err := c.ImageBuild(ctx, buildCtx, types.ImageBuildOptions{
		Dockerfile: "./Dockerfile",
		Tags:       []string{"built_by_ci"},
	})
	if err != nil {
		return errors.Wrap(err, "could not send image build request")
	}

	fd, isTerm := term.GetFdInfo(l)
	err = jsonmessage.DisplayJSONMessagesStream(resp.Body, l, fd, isTerm, nil)

	return err
}

func reportStatus(status string, pre *github.PullRequestEvent, link string) {
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")})
	tc := oauth2.NewClient(oauth2.NoContext, ts)
	c := github.NewClient(tc)

	owner := pre.PullRequest.Head.Repo.Owner.GetName()
	if owner == "" {
		owner = pre.PullRequest.Head.Repo.Owner.GetLogin()
	}

	repoName := pre.PullRequest.Head.Repo.GetName()
	ref := pre.PullRequest.Head.GetSHA()

	_, _, err := c.Repositories.CreateStatus(context.Background(), owner, repoName, ref, &github.RepoStatus{
		State:     &status,
		TargetURL: &link,
	})

	if err != nil {
		fmt.Println("could not report status to github", err)
	}
}

func logHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	logLines, ok := logs[id]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "id not found in memory")
		return
	}

	// TODO: can parse once outside.
	t := template.New("logs")
	t = template.Must(t.ParseFiles("./logs.html"))

	dt := &dataTemplate{logLines}

	err := t.Lookup("logs.html").Execute(w, dt)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
	}
}

type dataTemplate struct {
	Logs []string
}
