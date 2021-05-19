/*
 * @Description:
 * @Author: Tsingwong
 * @Date: 2021-05-19 23:18:12
 * @LastEditors: Tsingwong
 * @LastEditTime: 2021-05-19 23:24:06
 */
package main

var graph = make(map[string]map[string]bool)

func addEdge(from, to string) {
	edges := graph[from]
	if edges == nil {
		edges = make(map[string]bool)
		graph[from] = edges
	}
	edges[to] = true
}

func hasEdge(from, to string) bool {
	return graph[from][to]
}
