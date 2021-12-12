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
	graph.AllWalksToEnd()
	return walkCount
}

func passagePathingRevisit() int {
	graph := parseGraph(passageFile)
	graph.AllWalksToEndRevisit()
	return walkCount2
}

type Node struct {
	Value    string
	Children []*Node
	IsLarge  bool
}

var (
	visited   = map[string]bool{}
	walkCount int
)

func (n *Node) AllWalksToEnd() {
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
		adj.AllWalksToEnd()
	}
	visited[n.Value] = false
}

var (
	visited2     = map[string]int{}
	walkCount2   int
	visitedTwice bool
)

func (n *Node) AllWalksToEndRevisit() {
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
		adj.AllWalksToEndRevisit()
	}
	visited2[n.Value]--
	if visited2[n.Value] == 1 {
		visitedTwice = false
	}
}

func parseGraph(f string) *Node {
	graph := map[string]*Node{}
	graph["start"] = &Node{Value: "start"}
	graph["end"] = &Node{Value: "end"}
	for _, edge := range getEdges(f) {
		if _, ok := graph[edge[0]]; !ok {
			graph[edge[0]] = &Node{Value: edge[0], IsLarge: isUpper(edge[0])}
		}
		if _, ok := graph[edge[1]]; !ok {
			graph[edge[1]] = &Node{Value: edge[1], IsLarge: isUpper(edge[1])}
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
