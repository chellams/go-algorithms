package algorithms

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {

	input := []int{8, 7, 4, 6, 3, 5, 2}
	expected := []int{2, 3, 4, 5, 6, 7, 8}

	result := quickSort(input, 0, len(input)-1)

	if result != 0 {
		t.Errorf("error in quicksort")
	}

	if !reflect.DeepEqual(input, expected) {
		t.Errorf(" expected array is shoule be sorted, but %v", expected)
	}
}

func TestQuickSort_EmptySlice(t *testing.T) {

	var input []int

	result := quickSort(input, 0, len(input)-1)

	if result != OutOfBound {
		t.Errorf("error %d is expected", OutOfBound)
	}
}

func TestQuickSort_OutOfBound(t *testing.T) {

	var input []int

	result := quickSort(input, 0, 5)

	if result != OutOfBound {
		t.Errorf("error %d is expected", OutOfBound)
	}
}

func TestQuickSort_SingleElement(t *testing.T) {

	input := []int{5}
	expected := []int{5}

	result := quickSort(input, 0, len(input)-1)

	if result != 0 {
		t.Errorf("error in quicksort")
	}

	if !reflect.DeepEqual(input, expected) {
		t.Errorf(" expected array is shoule be sorted, but %v", expected)
	}
}

func TestQuickSort_Duplicate(t *testing.T) {

	input := []int{3, 5, 2, 6, 4, 8, 7, 5, 1, 2, 9, 5, 3}
	expected := []int{1, 2, 2, 3, 3, 4, 5, 5, 5, 6, 7, 8, 9}

	result := quickSort(input, 0, len(input)-1)

	if result != 0 {
		t.Errorf("error in quicksort")
	}

	if !reflect.DeepEqual(input, expected) {
		t.Errorf(" expected array is shoule be sorted, but %v", expected)
	}
}
