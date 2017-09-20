package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func betterScan() {
	f, _ := os.Open("./urls.txt")

	scnr := bufio.NewScanner(f)

	scnr.Scan()
	firstLine := scnr.Text()
	fmt.Println(firstLine)
	f.Close()
}

func main() {
	betterScan()
	if true {
		return
	}

	bts, err := ioutil.ReadFile("./urls.txt")
	if err != nil {
		fmt.Println("not cool", err)
		return
	}

	urlStrings := string(bts)
	urls := strings.Split(urlStrings, "\n")

	resp, err := http.Get(urls[0])
	if err != nil {
		fmt.Printf("Could not get %v domain: %v\n", urls[0], err)
		return
	}

	bts, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Could not read response body: %v\n", err)
		return
	}

	respBody := string(bts)
	hasHello := strings.Contains(respBody, "facebook")
	if hasHello {
		fmt.Println("We have facebook in the URL")
	}
}

//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//

// func main() {
// 	urlBytes, err := ioutil.ReadFile("./urls.txt")
// 	if err != nil {
// 		panic(err)
// 	}

// 	urlStr := string(urlBytes)

// 	urls := strings.Split(urlStr, "\n")
// 	for _, url := range urls {
// 		testDomain(url, "hello")
// 	}
// }

// func testDomain(url, target string) {
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		return
// 	}
// 	defer resp.Body.Close()

// 	bts, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		return
// 	}

// 	if strings.Contains(string(bts), "hello") {
// 		fmt.Println(url)
// 	}
// }
