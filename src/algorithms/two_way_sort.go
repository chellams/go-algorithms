package algorithms

/*
Takes two sorted array and merge it into single array.
*/
func twoWaySort(A []int, B []int) []int {
	i := 0
	j := 0
	var C []int
	for i < len(A) && j < len(B) {
		if A[i] < B[j] {
			C = append(C, A[i])
			i++
		} else {
			C = append(C, B[j])
			j++
		}
	}

	for i < len(A) {
		C = append(C, A[i])
		i++
	}
	for j < len(B) {
		C = append(C, B[j])
		j++
	}
	return C
}
