/*
 * @Description:
 * @Author: Tsingwong
 * @Date: 2021-05-28 14:26:47
 * @LastEditors: Tsingwong
 * @LastEditTime: 2021-05-28 14:39:11
 */
package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		links, err := findLinks(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "findlinks2: %v \n", err)
			continue
		}
		for _, link := range links {
			fmt.Println(link)
		}
	}

}

func findLinks(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	return visit(nil, doc), nil
}

// 遍历 HTML 节点树，从每个 a 标签的 href 属性获得 link，将这些 links 存入 slice 里，并返回这个 slice
func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	if n.Type == html.ElementNode && (n.Data == "img" || n.Data == "script") {
		for _, a := range n.Attr {
			if a.Key == "src" {
				links = append(links, a.Val)
			}
		}
	}
	if n.Type == html.ElementNode && n.Data == "link" {
		var isStyleSheet bool
		for _, a := range n.Attr {
			if isStyleSheet && a.Key == "href" {
				links = append(links, a.Val)
			}
			if a.Key == "src" && a.Val == "stylesheet" {
				isStyleSheet = true
			}
		}
	}
	// for c := n.FirstChild; c != nil; c = c.NextSibling {
	// 	links = visit(links, c)
	// }
	if n.FirstChild != nil {
		links = visit(links, n.FirstChild)
	}
	if n.NextSibling != nil {
		links = visit(links, n.NextSibling)
	}
	return links
}
