/*
 * @Description:
 * @Author: Tsingwong
 * @Date: 2021-05-26 22:40:15
 * @LastEditors: Tsingwong
 * @LastEditTime: 2021-05-26 22:40:17
 */
/*
 * @Description:
 * @Author: Tsingwong
 * @Date: 2021-05-26 22:16:00
 * @LastEditors: Tsingwong
 * @LastEditTime: 2021-05-26 22:36:02
 */
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
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
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}
