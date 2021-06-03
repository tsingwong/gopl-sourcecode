/*
 * @Description:
 * @Author: Tsingwong
 * @Date: 2021-06-03 07:12:00
 * @LastEditors: Tsingwong
 * @LastEditTime: 2021-06-03 15:10:45
 */
package links

import (
	"fmt"
	"net/http"

	"golang.org/x/net/html"
)

// func main() {
// 	for _, url := range os.Args[1:] {
// 		links, err := Extract(url)
// 		if err != nil {
// 			fmt.Fprintf(os.Stderr, "link: %v", err)
// 			continue
// 		}
// 		for _, link := range links {
// 			fmt.Println(link)
// 		}
// 	}
// }

func Extract(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	var links []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					// Bad URL
					continue
				}
				links = append(links, link.String())
			}
		}
	}
	forEachNode(doc, visitNode, nil)
	return links, nil
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}
