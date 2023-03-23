package algorithms

import (
	"sort"
)

func optimalMergePattern(slice [][]int) []int {
	if slice == nil {
		return nil
	}
	if len(slice) < 2 {
		return slice[0]
	}
	for true {
		sortToFindMinLength(slice)
		if len(slice) >= 2 {
			result := twoWaySort(slice[0], slice[1])
			slice = slice[1:]
			slice = slice[1:]
			slice = append(slice, result)
		} else {
			break
		}
	}

	return slice[0]
}

func sortToFindMinLength(slice [][]int) {
	sort.Slice(slice, func(i, j int) bool {
		return len(slice[i]) < len(slice[j])
	})
}
