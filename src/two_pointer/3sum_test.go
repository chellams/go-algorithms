package two_pointer

import (
	"testing"
)

func Test3Sum(t *testing.T) {
	nums := []int{-1, 0, 1, 2, -1, -4}
	expected := []string{"-1,-1,2", "-1,0,1"}
	actual := threeSum(nums)
	isNotPassed := true
	idx := 0
	for key, _ := range actual {
		if key != expected[idx] {
			isNotPassed = false
			break
		}
		idx++
	}

	if !isNotPassed {
		t.Errorf(" received wrong sets %v", actual)
	}
}

func Test3Sum_Nil(t *testing.T) {
	actual := threeSum(nil)

	if actual != nil {
		t.Errorf(" received wrong sets %v", actual)
	}
}

func Test3Sum_Empty(t *testing.T) {
	ints := []int{}
	actual := threeSum(ints)
	if actual != nil {
		t.Errorf(" received wrong sets %v", actual)
	}

	ints = []int{1}
	actual = threeSum(ints)
	if actual != nil {
		t.Errorf(" received wrong sets %v", actual)
	}

	ints = []int{1, 2}
	actual = threeSum(ints)
	if actual != nil {
		t.Errorf(" received wrong sets %v", actual)
	}
}
