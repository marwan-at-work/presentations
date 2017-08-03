package main

import (
	"context"
	"fmt"
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

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/webhook", webhookHandler)

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
			fmt.Println("building pr")
			err = buildPR(pre)
			fmt.Println("pr ", err)
			status := "success"
			if err != nil {
				status = "error"
			}

			reportStatus(status, pre)
		}
	default:
		fmt.Fprintf(w, "Uninteresting event: %T", event)
		return
	}

	for key, vals := range r.Header {
		fmt.Printf("key: %v - vals: %v\n", key, vals)
	}
}

func buildPR(pre *github.PullRequestEvent) error {
	status := "pending"
	reportStatus(status, pre)

	pathToRepo, err := cloneRepo(pre)
	if err != nil {
		return err
	}
	defer os.RemoveAll(pathToRepo)

	err = buildImage(pathToRepo)

	return err
}

func cloneRepo(pre *github.PullRequestEvent) (string, error) {
	repoName := pre.PullRequest.Head.Repo.GetName()
	tempDir, err := ioutil.TempDir("", repoName)
	if err != nil {
		return "", err
	}

	cloneURL := pre.PullRequest.Head.Repo.GetCloneURL()
	branchName := pre.PullRequest.Head.GetRef()

	repo, err := git.PlainClone(tempDir, false, &git.CloneOptions{
		URL:               cloneURL,
		Progress:          os.Stdout,
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

func buildImage(repoDir string) error {
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

	fd, isTerm := term.GetFdInfo(os.Stdout)
	err = jsonmessage.DisplayJSONMessagesStream(resp.Body, os.Stdout, fd, isTerm, nil)

	return err
}

func reportStatus(status string, pre *github.PullRequestEvent) {
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")})
	tc := oauth2.NewClient(oauth2.NoContext, ts)
	c := github.NewClient(tc)

	owner := pre.PullRequest.Head.Repo.Owner.GetName()
	repoName := pre.PullRequest.Head.Repo.GetName()
	ref := pre.PullRequest.Head.GetSHA()

	ss, _, err := c.Repositories.CreateStatus(context.Background(), owner, repoName, ref, &github.RepoStatus{
		State: &status,
	})

	fmt.Println(ss, "and", err)
}
