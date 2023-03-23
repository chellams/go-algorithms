package algorithms

func mergeSort(A []int, l, h int) []int {
	if len(A) == 0 {
		return nil
	}
	if l < h {
		mid := (l + h) / 2
		X := mergeSort(A, l, mid)
		Y := mergeSort(A, mid+1, h)
		return twoWaySort(X, Y)
	}
	return []int{A[l]}
}
