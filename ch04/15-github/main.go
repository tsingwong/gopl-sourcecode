/*
 * @Description:
 * @Author: Tsingwong
 * @Date: 2021-05-23 18:48:28
 * @LastEditors: Tsingwong
 * @LastEditTime: 2021-05-23 19:12:54
 */
package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// See https://developer.github.com/v3/search/#search-issues.
const IssuesURL = "https://api.github.com/search/issues"

type IssuesSearchResult struct {
	TotalCount int `json: "total_count"`
	Items      []*Issue
}

type Issue struct {
	Number   int
	HTMLURL  string `json: "html_url"`
	Title    string
	State    string
	User     string
	CreateAt time.Time `json:"create_at"`
	Body     string
}

type User struct {
	Login   string
	HTMLURL string `json: "html_url"`
}

func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}

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
