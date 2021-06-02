/*
 * @Description:
 * @Author: Tsingwong
 * @Date: 2021-06-02 22:51:52
 * @LastEditors: Tsingwong
 * @LastEditTime: 2021-06-03 07:02:44
 */
package main

import (
	"fmt"
	"sort"
)

// 每个课程的前置课程
// 拓扑排序：前置条件可以构成有向图。
// 顶点是课程，边是课程之间的依赖关系
var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d: \t%s\n", i+1, course)
	}

}

func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items []string)
	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}
	// 遍历 按照 key 排序方便后面的顺序读取
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	visitAll(keys)
	return order
}
