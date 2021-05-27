/*
 * @Description:
 * @Author: Tsingwong
 * @Date: 2021-05-26 22:16:00
 * @LastEditors: Tsingwong
 * @LastEditTime: 2021-05-27 14:16:50
 */
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

var freq = make(map[string]int)

func main() {
	// s := `<ul><li><a href="foo">Foo</a><li><a href="/bar/baz">BarBaz</a></ul>`
	// doc, err := html.Parse(strings.NewReader(s))
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
	// for tag, count := range freq {
	// 	fmt.Printf("%4d %s\n", count, tag)
	// }
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
		freq[string(n.Data)]++
		links = visit(links, n.FirstChild)
	}
	if n.NextSibling != nil {
		freq[string(n.Data)]++
		links = visit(links, n.NextSibling)
	}
	return links
}
