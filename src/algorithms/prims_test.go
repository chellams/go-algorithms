package algorithms

import (
	"testing"
)

func TestPrims(t *testing.T) {
	var graph = [][]int{
		{0, 2, 0, 6, 0},
		{2, 0, 3, 8, 5},
		{0, 3, 0, 0, 7},
		{6, 8, 0, 0, 9},
		{0, 5, 7, 9, 0}}
	totalValue := prims(graph)
	expectedValue := 16

	if totalValue != expectedValue {
		t.Errorf(" total path value is not correct, expected %d, but got %d", expectedValue, totalValue)
	}
}

func TestPrims_Null(t *testing.T) {
	var graph [][]int
	totalValue := prims(graph)
	expectedValue := -1

	if totalValue != expectedValue {
		t.Errorf(" total path value is not correct, expected %d, but got %d", expectedValue, totalValue)
	}
}
