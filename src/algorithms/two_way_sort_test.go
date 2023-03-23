package algorithms

import (
	"reflect"
	"testing"
)

func TestTwoWaySort(t *testing.T) {
	A := []int{2, 3, 9, 10}
	B := []int{4, 5, 6, 8}
	expected := []int{2, 3, 4, 5, 6, 8, 9, 10}
	C := twoWaySort(A, B)

	if !reflect.DeepEqual(expected, C) {
		t.Errorf("given array should be sorted, but result is %v", expected)
	}
}

func TestTwoWaySort_WithUnequalList(t *testing.T) {
	// first slice is bigger than second slice
	A := []int{2, 3, 9, 10, 14, 19, 23, 57, 89}
	B := []int{4, 5, 6, 8}
	expected := []int{2, 3, 4, 5, 6, 8, 9, 10, 14, 19, 23, 57, 89}
	C := twoWaySort(A, B)

	if !reflect.DeepEqual(expected, C) {
		t.Errorf("given array should be sorted, but result is %v", expected)
	}

	// first slice is smaller than second slice
	X := []int{4, 5, 6, 8}
	Y := []int{2, 3, 9, 10, 14, 19, 23, 57, 89}
	expected = []int{2, 3, 4, 5, 6, 8, 9, 10, 14, 19, 23, 57, 89}
	Z := twoWaySort(Y, X)

	if !reflect.DeepEqual(expected, Z) {
		t.Errorf("given array should be sorted, but result is %v", expected)
	}
}
