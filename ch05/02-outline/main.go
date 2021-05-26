/*
 * @Description:
 * @Author: Tsingwong
 * @Date: 2021-05-26 22:43:52
 * @LastEditors: Tsingwong
 * @LastEditTime: 2021-05-26 23:16:44
 */
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

var arr [6]string

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}
	var stack = arr[0:0]
	fmt.Printf("stack:%v, len:%v, cap:%v;\narr:%v, len:%v, cap:%v;\n", stack, len(stack), cap(stack), arr, len(arr), cap(arr))

	outline(stack, doc)
	fmt.Printf("stack:%v, len:%v, cap:%v;\narr:%v, len:%v, cap:%v;\n", stack, len(stack), cap(stack), arr, len(arr), cap(arr))
}

func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data) // push tagName
		fmt.Printf("stack:%v, len:%v, cap:%v;\narr:%v, len:%v, cap:%v;\n", stack, len(stack), cap(stack), arr, len(arr), cap(arr))
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		// 虽然没有做出栈处理
		// outline 调用自身，接收到的是 stack 的拷贝
		// 这时候对 stack 的追加操作，修改的是 stack 的拷贝
		// 可能会导致其底层数组的变化或申请新内存空间扩容
		// 需要理解当函数返回时
		// 实际上底层数组已经改了，只是由于这个 stack 仍保持之前的len 和cap
		outline(stack, c)
	}
}
