package algorithms

import (
	"fmt"
	"sort"
)

const (
	InvalidInput = -1
)

type node struct {
	weight int
	src    int
	dst    int
}

type edge struct {
	parent int
	rank   int
}

var parentEdges []edge
var resultMst []node

func findRoot(v int) int {
	if parentEdges[v].parent == -1 {
		return v
	}
	parentEdges[v].parent = findRoot(parentEdges[v].parent)
	return parentEdges[v].parent
}

func union(src int, dst int) {
	if parentEdges[src].parent > parentEdges[dst].parent {
		parentEdges[dst].parent = src
	} else if parentEdges[src].parent < parentEdges[dst].parent {
		parentEdges[src].parent = dst
	} else {
		parentEdges[src].parent = dst
		parentEdges[dst].rank++
	}
}

func kruskal(graph [][]int) int {

	if graph == nil {
		return InvalidInput
	}

	edges := getEdgeArray(graph)

	for i := 0; i < len(graph); i++ {
		e := edge{rank: 0, parent: -1}
		parentEdges = append(parentEdges, e)
	}

	j := 0
	i := 0
	V := len(graph)
	E := V - 1

	totalWeight := 0
	for i < V-1 && j < E {
		n := edges[j]
		src := findRoot(n.src)
		dst := findRoot(n.dst)

		if src == dst {
			j++
			continue
		}
		union(src, dst)
		resultMst = append(resultMst, n)
		totalWeight += n.weight
		i++
	}

	return totalWeight
}

/*
Optional method : It's used to form the edges from given adjacency metrix.
*/
func getEdgeArray(graph [][]int) []node {

	var edges []node

	edgeMap := make(map[string]struct{})
	for i := 0; i < len(graph); i++ {
		row := graph[i]
		for j := 0; j < len(row); j++ {
			if i != j && graph[i][j] != 0 {
				n := node{src: i, dst: j, weight: graph[i][j]}
				if !isEdgeAlreadyExists(edgeMap, n) {
					edges = append(edges, n)
					e := fmt.Sprintf("%d-%d", i, j)
					edgeMap[e] = struct{}{}
				}
			}
		}
	}

	sort.Slice(edges, func(i, j int) bool {
		return edges[i].weight < edges[j].weight
	})

	return edges
}

func isEdgeAlreadyExists(edgeMap map[string]struct{}, input node) bool {
	_, b1 := edgeMap[fmt.Sprintf("%d-%d", input.src, input.dst)]
	_, b2 := edgeMap[fmt.Sprintf("%d-%d", input.dst, input.src)]
	return b1 || b2
}
