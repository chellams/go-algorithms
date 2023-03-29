package two_pointer

import (
	"reflect"
	"testing"
)

func TestRemoveDuplicatesFromSortedArray(t *testing.T) {
	ints := []int{1, 2, 2, 2, 3, 4, 5, 5}
	expected := []int{1, 2, 3, 4, 5, -1, -1, -1}

	actual := removeDuplicates(ints)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf(" expected array is %v, but got %v", expected, actual)
	}
}

func TestRemoveDuplicatesFromSortedArray_Nil(t *testing.T) {

	actual := removeDuplicates(nil)

	if actual != nil {
		t.Errorf(" expected outpout is nil, but got %v", actual)
	}
}

func TestRemoveDuplicatesFromSortedArray_Empty(t *testing.T) {

	ints := []int{}
	actual := removeDuplicates(ints)

	if actual != nil {
		t.Errorf(" expected outpout is nil, but got %v", actual)
	}
}
