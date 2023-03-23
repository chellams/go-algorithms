package algorithms

import (
	"math"
)

func prims(graph [][]int) int {
	if graph == nil {
		return -1
	}

	var v = len(graph)
	var parent = make([]int, v)
	var set = make([]bool, v)
	var values = make([]int, v)

	for i := 0; i < v; i++ {
		values[i] = math.MaxInt
	}

	values[0] = 0
	parent[0] = -1
	for i := 0; i < v; i++ {
		u := findMinVertex(set, values)
		set[u] = true
		for j := 0; j < v; j++ {
			if graph[u][j] != 0 && !set[j] && graph[u][j] < values[j] {
				values[j] = graph[u][j]
				parent[j] = u
			}
		}
	}

	totalValue := 0
	for _, value := range values {
		totalValue += value
	}

	return totalValue
}

func findMinVertex(set []bool, values []int) int {
	minVertex := math.MaxInt
	idx := 0

	for i := 0; i < len(values); i++ {
		if !set[i] && values[i] < minVertex {
			minVertex = values[i]
			idx = i
		}
	}

	return idx
}
