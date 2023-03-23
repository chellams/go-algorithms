package algorithms

import (
	"reflect"
	"testing"
)

func TestHeapSort(t *testing.T) {
	input := []int{5, 4, 6, 3, 7, 8, 2}
	heapSort(input)
	expected := []int{2, 3, 4, 5, 6, 7, 8}

	if !reflect.DeepEqual(input, expected) {
		t.Errorf("array %v is not sorted, result array is %v", input, expected)
	}
}

func TestHeapify(t *testing.T) {
	input := []int{5, 4, 6, 3, 7, 8, 2}
	size := len(input)
	for i := size / 2; i >= 0; i-- {
		heapify(input, i, size)
	}
	expected := []int{8, 7, 6, 3, 4, 5, 2}

	if !reflect.DeepEqual(input, expected) {
		t.Errorf("unable to make max heap from the array %v, resulut array is %v", input, expected)
	}
}
