package two_pointer

import (
	"fmt"
	"sort"
)

/*
Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] such that
i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.
Notice that the solution set must not contain duplicate triplets.
*/
func threeSum(array []int) map[string]struct{} {

	if array == nil {
		return nil
	}

	if len(array) == 0 || len(array) == 1 || len(array) == 2 {
		return nil
	}

	sort.Slice(array, func(i, j int) bool {
		return array[i] < array[j]
	})

	resultMap := make(map[string]struct{})
	for i := 0; i < len(array)-2; i++ {
		a := i + 1
		b := len(array) - 1

		for a < b {
			if array[i]+array[a]+array[b] == 0 {
				result := fmt.Sprintf("%d,%d,%d", array[i], array[a], array[b])
				resultMap[result] = struct{}{}
				a++
				b--
			} else if array[i]+array[a]+array[b] < 0 {
				a++
			} else {
				b--
			}
		}
	}
	return resultMap
}
