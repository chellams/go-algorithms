package algorithms

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	A := []int{3, 2, 6, 4, 8, 7, 5, 1, 9}
	result := mergeSort(A, 0, len(A)-1)
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf(" resultant array should be sorted, but %v", result)
	}
}

func TestMergeSort_SingleElement(t *testing.T) {
	A := []int{3}
	result := mergeSort(A, 0, len(A)-1)
	expected := []int{3}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf(" resultant array should be sorted, but got %v", result)
	}
}

func TestMergeSort_EmptySlice(t *testing.T) {
	var A []int
	result := mergeSort(A, 0, len(A)-1)
	if result != nil {
		t.Errorf(" resultant array should be nil, but got %v", result)
	}
}

func TestMergeSort_Duplicate(t *testing.T) {
	A := []int{3, 5, 2, 6, 4, 8, 7, 5, 1, 2, 9, 5, 3}
	result := mergeSort(A, 0, len(A)-1)
	expected := []int{1, 2, 2, 3, 3, 4, 5, 5, 5, 6, 7, 8, 9}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf(" resultant array should be sorted, but %v", result)
	}
}
