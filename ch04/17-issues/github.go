/*
 * @Description:
 * @Author: Tsingwong
 * @Date: 2021-05-25 06:58:34
 * @LastEditors: Tsingwong
 * @LastEditTime: 2021-05-25 07:13:33
 */
package main

import (
	"fmt"
	"time"
)

const APIURL = "https://github.com/"
const IssueURL = APIURL + "search/issues"

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string
}

func (i Issue) CacheURL() string {
	return fmt.Sprintf("/issues/%d", i.Number)
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}
