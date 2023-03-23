package algorithms

import (
	"reflect"
	"testing"
)

func TestOptimalMergePattern(t *testing.T) {
	array := [][]int{
		{2, 4, 5, 7, 8, 9},                       // 6
		{5, 7, 8, 9, 10, 13, 15, 23, 26, 47, 48}, // 11
		{3, 4, 6, 7, 9, 11, 65},                  // 7
		{1, 3, 4, 7, 9},                          // 5
		{4, 6, 7, 8, 12, 15, 16, 19, 23, 26, 27}, // 11
	}
	expected := []int{1, 2, 3, 3, 4, 4, 4, 4, 5, 5, 6, 6,
		7, 7, 7, 7, 7, 8, 8, 8, 9, 9, 9, 9, 10, 11, 12, 13,
		15, 15, 16, 19, 23, 23, 26, 26, 27, 47, 48, 65}

	result := optimalMergePattern(array)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("expected array should be sorted , but %v", expected)
	}
}

func TestOptimalMergePattern_OneSlice(t *testing.T) {
	array := [][]int{
		{5, 7, 8, 9, 10, 13, 15, 23, 26, 47, 48}, // 11
	}
	expected := []int{5, 7, 8, 9, 10, 13, 15, 23, 26, 47, 48}

	result := optimalMergePattern(array)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("expected array should be sorted , but %v", expected)
	}
}

func TestOptimalMergePattern_TwoSlice(t *testing.T) {
	array := [][]int{
		{2, 4, 5, 7, 8, 9},                       // 6
		{5, 7, 8, 9, 10, 13, 15, 23, 26, 47, 48}, // 11
	}
	expected := []int{2, 4, 5, 5, 7, 7, 8, 8, 9, 9, 10, 13, 15, 23, 26, 47, 48}

	result := optimalMergePattern(array)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("expected array should be sorted , but %v", expected)
	}
}

func TestOptimalMergePattern_EmptySlice(t *testing.T) {
	var array [][]int
	var expected []int

	result := optimalMergePattern(array)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("expected array should be sorted , but %v", expected)
	}
}

func TestOptimalMergePattern_OneOfEmpty(t *testing.T) {
	array := [][]int{
		{2, 4, 5, 7, 8, 9},                       // 6
		{5, 7, 8, 9, 10, 13, 15, 23, 26, 47, 48}, // 11
		{},                                       // 0
		{1, 3, 4, 7, 9},                          // 5
		{4, 6, 7, 8, 12, 15, 16, 19, 23, 26, 27}, // 11
	}
	expected := []int{1, 2, 3, 4, 4, 4, 5, 5, 6, 7, 7, 7, 7,
		8, 8, 8, 9, 9, 9, 10, 12, 13, 15, 15, 16, 19, 23, 23,
		26, 26, 27, 47, 48}

	result := optimalMergePattern(array)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("expected array should be sorted , but %v", expected)
	}
}

func TestOptimalMergePattern_AllEmpty(t *testing.T) {
	array := [][]int{
		{}, // 0
		{}, // 0
		{}, // 0
		{}, // 0
		{}, // 0
	}
	var expected []int

	result := optimalMergePattern(array)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("expected array should be sorted , but %v", expected)
	}
}
