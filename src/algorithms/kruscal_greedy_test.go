package algorithms

import "testing"

func TestKruskal(t *testing.T) {
	graph := [][]int{
		{0, 2, 0, 6, 0},
		{2, 0, 3, 8, 5},
		{0, 3, 0, 0, 7},
		{6, 8, 0, 0, 9},
		{0, 5, 7, 9, 0},
	}
	expected := 16
	result := kruskal(graph)

	if result != expected {
		t.Errorf("shortest path expected is %d, but got %d", expected, result)
	}
}

func TestKruskal_Null(t *testing.T) {
	var graph [][]int
	result := kruskal(graph)

	if result != InvalidInput {
		t.Errorf("given input is invalid %v", graph)
	}
}
