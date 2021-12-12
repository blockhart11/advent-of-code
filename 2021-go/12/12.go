package _12

import (
	"strings"
)

type node struct {
	name string
	edges []string
}

func countPaths(graph []node, start string, path string) int {
	var result int
	var startNode *node
	if start == "" {
		start = "start"
		path = "start"
	}
	startNode = getNode(start, graph)
	for _, edge := range startNode.edges {
		if isEnd(edge) {
			result++
			//fmt.Println("Path:", path + "-end")
			continue
		}
		if isStart(edge) {
			continue
		}
		if canVisit(edge, path) {
			result += countPaths(graph, edge, path + "-" + edge)
		}
	}
	return result
}

func countPathsDoubleBackOnce(graph []node, start string, path string, doubledBack bool) int {
	var result int
	var startNode *node
	if start == "" {
		start = "start"
		path = "start"
	}
	startNode = getNode(start, graph)
	for _, edge := range startNode.edges {
		if isEnd(edge) {
			result++
			//fmt.Println("Path:", path + "-end")
			continue
		}
		if isStart(edge) {
			continue
		}
		if canVisit(edge, path) {
			result += countPathsDoubleBackOnce(graph, edge, path + "-" + edge, doubledBack)
		} else if isSmall(edge) && !doubledBack {
			// small room, but second visit
			result += countPathsDoubleBackOnce(graph, edge, path + "-" + edge, true)
		}
	}
	return result
}

func getNode(name string, graph []node) *node {
	for i, n := range graph {
		if n.name == name {
			return &graph[i]
		}
	}
	return nil
}

func getStart(graph []node) *node {
	return getNode("start", graph)
}

func isStart(n string) bool {
	return n == "start"
}

func isEnd(n string) bool {
	return n == "end"
}

func isSmall(n string) bool {
	return n == strings.ToLower(n)
}

func canVisit(n string, path string) bool {
	return !isSmall(n) || !strings.Contains(path, "-"+n+"-")
}