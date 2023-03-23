package algorithms

import (
	"testing"
)

func TestBinarySearch_AvailableNumber(t *testing.T) {
	var array = []int{2, 3, 4, 5, 6, 8, 9, 13, 15, 16, 19, 24, 26, 29, 36, 37, 39, 41, 43, 47, 48, 49, 50, 51, 53, 54}
	var given = 19
	index := BinarySearch(array, given, 0, len(array))
	if array[index] != given {
		t.Errorf(" error in finding element %v", given)
	}
}

func TestBinarySearch_UnavailableNumber(t *testing.T) {
	var array = []int{2, 3, 4, 5, 6, 8, 9, 13, 15, 16, 19, 24, 26, 29, 36, 37, 39, 41, 43, 47, 48, 49, 50, 51, 53, 54}
	var given = 20
	index := BinarySearch(array, given, 0, len(array))
	if index != NotFound {
		t.Errorf(" error in finding element %v", given)
	}
}

func TestBinarySearch_InvalidHigh(t *testing.T) {
	var array = []int{2, 3, 4, 5, 6, 8, 9, 13, 15, 16, 19, 24, 26, 29, 36, 37, 39, 41, 43, 47, 48, 49, 50, 51, 53, 54}
	var given = 20
	index := BinarySearch(array, given, 0, len(array)+2)
	if index != OutOfBound {
		t.Errorf(" error in finding element %v", given)
	}
}

func TestBinarySearch_InvalidLow(t *testing.T) {
	var array = []int{2, 3, 4, 5, 6, 8, 9, 13, 15, 16, 19, 24, 26, 29, 36, 37, 39, 41, 43, 47, 48, 49, 50, 51, 53, 54}
	var given = 20
	index := BinarySearch(array, given, -1, len(array))
	if index != OutOfBound {
		t.Errorf(" error in finding element %v", given)
	}
}
