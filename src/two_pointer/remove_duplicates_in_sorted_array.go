package two_pointer

/*
*
Given an integer array nums sorted in non-decreasing order, remove the duplicates in-place such that
each unique element appears only once. The relative order of the elements should be kept the same.
*/
func removeDuplicates(arr []int) []int {

	if arr == nil {
		return nil
	}

	if len(arr) == 0 {
		return nil
	}

	j := 0
	for i := 0; i+1 < len(arr); i++ {
		if arr[i] != arr[i+1] {
			arr[j+1] = arr[i+1]
			j++
		}
	}
	j++
	for j < len(arr) {
		arr[j] = -1
		j++
	}

	return arr
}
