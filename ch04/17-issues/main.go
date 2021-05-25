/*
 * @Description:
 * @Author: Tsingwong
 * @Date: 2021-05-25 06:56:58
 * @LastEditors: Tsingwong
 * @LastEditTime: 2021-05-25 22:36:40
 */
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"strings"
)

func search(query []string) {
	result, err := SearchIssues(query)
	if err != nil {
		log.Fatal(err)
	}
	for _, item := range result.Items {
		format := "#%-5d %9.9s %.55s\n"
		fmt.Printf(format, item.Number, item.User.Login, item.Title)
	}
}

func read(owner, repo, number string) {
	issue, err := GetIssue(owner, repo, number)
	if err != nil {
		log.Fatal(err)
	}
	body := issue.Body
	if body == "" {
		body = "<empty>\n"
	}
	fmt.Printf("repo: %s/%s\nnumber: %s\nuser: %s\ntitle: %s\n\n%s",
		owner, repo, number, issue.User.Login, issue.Title, body)
}

func edit(owner string, repo string, number string) {
	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = "vim"
	}
	editorPath, err := exec.LookPath(editor)
	if err != nil {
		log.Fatal(err)
	}
	tempfile, err := ioutil.TempFile("", "issue_crud")
	if err != nil {
		log.Fatal(err)
	}
	defer tempfile.Close()
	defer os.Remove(tempfile.Name())

	issue, err := GetIssue(owner, repo, number)
	if err != nil {
		log.Fatal(err)
	}

	encoder := json.NewEncoder(tempfile)
	err = encoder.Encode(map[string]string{
		"title": issue.Title,
		"state": issue.State,
		"body":  issue.Body,
	})
	if err != nil {
		log.Fatal(err)
	}

	cmd := &exec.Cmd{
		Path:   editorPath,
		Args:   []string{editor, tempfile.Name()},
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	tempfile.Seek(0, 0)
	fields := make(map[string]string)
	if err = json.NewDecoder(tempfile).Decode(&fields); err != nil {
		log.Fatal(err)
	}

	_, err = EditIssue(owner, repo, number, fields)
	if err != nil {
		log.Fatal(err)
	}
}

func close_(owner string, repo string, number string) {
	_, err := EditIssue(owner, repo, number, map[string]string{"state": "closed"})
	if err != nil {
		log.Fatal(err)
	}
}

func open(owner string, repo string, number string) {
	_, err := EditIssue(owner, repo, number, map[string]string{"state": "open"})
	if err != nil {
		log.Fatal(err)
	}
}

var usage string = `usage:
 search QUERY
 [read|edit|close|open] OWNER REPO ISSUE_NUMBER
 `

func usageDie() {
	fmt.Fprintln(os.Stderr, usage)
	os.Exit(1)
}

func main() {
	if len(os.Args) < 2 {
		usageDie()
	}
	cmd := os.Args[1]
	args := os.Args[2:]

	if cmd == "search" {
		if len(args) < 1 {
			usageDie()
		}
		search(args)
		os.Exit(0)
	}
	if len(args) != 3 {
		usageDie()
	}
	owner, repo, number := args[0], args[1], args[2]
	switch cmd {
	case "read":
		read(owner, repo, number)
	case "edit":
		edit(owner, repo, number)
	case "close":
		close_(owner, repo, number)
	case "open":
		open(owner, repo, number)
	}
}

func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}
	//!-
	// For long-term stability, instead of http.Get, use the
	// variant below which adds an HTTP request header indicating
	// that only version 3 of the GitHub API is acceptable.
	//
	//   req, err := http.NewRequest("GET", IssuesURL+"?q="+q, nil)
	//   if err != nil {
	//       return nil, err
	//   }
	//   req.Header.Set(
	//       "Accept", "application/vnd.github.v3.text-match+json")
	//   resp, err := http.DefaultClient.Do(req)
	//!+

	// We must close resp.Body on all execution paths.
	// (Chapter 5 presents 'defer', which makes this simpler.)
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}

// remember to close the response body.
func get(url string) (*http.Response, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("can't get %s: %s", url, resp.Status)
	}
	return resp, nil
}

func GetIssue(owner string, repo string, number string) (*Issue, error) {
	url := strings.Join([]string{APIURL, "repos", owner, repo, "issues", number}, "/")
	resp, err := get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var issue Issue
	if err := json.NewDecoder(resp.Body).Decode(&issue); err != nil {
		return nil, err
	}
	return &issue, nil
}

func GetIssues(owner string, repo string) ([]Issue, error) {
	url := strings.Join([]string{APIURL, "repos", owner, repo, "issues"}, "/")
	resp, err := get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var issues []Issue
	if err := json.NewDecoder(resp.Body).Decode(&issues); err != nil {
		return nil, err
	}
	return issues, nil
}

func EditIssue(owner, repo, number string, fields map[string]string) (*Issue, error) {
	buf := &bytes.Buffer{}
	encoder := json.NewEncoder(buf)
	err := encoder.Encode(fields)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	url := strings.Join([]string{APIURL, "repos", owner, repo, "issues", number}, "/")
	req, err := http.NewRequest("PATCH", url, buf)
	req.SetBasicAuth(os.Getenv("GITHUB_USER"), os.Getenv("GITHUB_PASS"))
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to edit issue: %s", resp.Status)
	}
	var issue Issue
	if err = json.NewDecoder(resp.Body).Decode(&issue); err != nil {
		return nil, err
	}
	return &issue, nil

}
