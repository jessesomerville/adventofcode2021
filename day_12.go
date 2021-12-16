package main

import (
	_ "embed"
	"strings"
	"unicode"
)

var (
	//go:embed inputs/day_12.txt
	passageFile string
)

func passagePathing() int {
	graph := parseGraph(passageFile)
	graph.allWalksToEnd()
	return walkCount
}

func passagePathingRevisit() int {
	graph := parseGraph(passageFile)
	graph.allWalksToEndRevisit()
	return walkCount2
}

type node struct {
	Value    string
	Children []*node
	IsLarge  bool
}

var (
	visited   = map[string]bool{}
	walkCount int
)

func (n *node) allWalksToEnd() {
	if visited[n.Value] {
		return
	}
	if !n.IsLarge {
		visited[n.Value] = true
	}
	if n.Value == "end" {
		walkCount++
		visited[n.Value] = false
		return
	}
	for _, adj := range n.Children {
		adj.allWalksToEnd()
	}
	visited[n.Value] = false
}

var (
	visited2     = map[string]int{}
	walkCount2   int
	visitedTwice bool
)

func (n *node) allWalksToEndRevisit() {
	if visited2[n.Value] > 0 && visitedTwice {
		return
	}
	if !n.IsLarge {
		visited2[n.Value]++
		if visited2[n.Value] == 2 {
			visitedTwice = true
		}
	}
	if n.Value == "end" {
		walkCount2++
		visited2[n.Value]--
		return
	}
	for _, adj := range n.Children {
		adj.allWalksToEndRevisit()
	}
	visited2[n.Value]--
	if visited2[n.Value] == 1 {
		visitedTwice = false
	}
}

func parseGraph(f string) *node {
	graph := map[string]*node{}
	graph["start"] = &node{Value: "start"}
	graph["end"] = &node{Value: "end"}
	for _, edge := range getEdges(f) {
		if _, ok := graph[edge[0]]; !ok {
			graph[edge[0]] = &node{Value: edge[0], IsLarge: isUpper(edge[0])}
		}
		if _, ok := graph[edge[1]]; !ok {
			graph[edge[1]] = &node{Value: edge[1], IsLarge: isUpper(edge[1])}
		}
		graph[edge[0]].Children = append(graph[edge[0]].Children, graph[edge[1]])
	}
	return graph["start"]
}

// See if you can append an edge without making a new slice
func getEdges(f string) [][]string {
	lines := strings.Split(f, "\n")

	edges := [][]string{}
	for _, line := range lines {
		nodeVals := strings.Split(line, "-")
		if nodeVals[0] == "start" {
			edges = append(edges, nodeVals)
		} else if nodeVals[1] == "start" {
			edges = append(edges, []string{"start", nodeVals[0]})
		} else if nodeVals[0] == "end" {
			edges = append(edges, []string{nodeVals[1], "end"})
		} else if nodeVals[1] == "end" {
			edges = append(edges, nodeVals)
		} else {
			edges = append(edges, nodeVals)
			edges = append(edges, []string{nodeVals[1], nodeVals[0]})
		}
	}
	return edges
}

func isUpper(s string) bool {
	for _, r := range s {
		if !unicode.IsUpper(r) {
			return false
		}
	}
	return true
}
