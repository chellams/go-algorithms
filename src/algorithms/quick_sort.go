package algorithms

func partition(A []int, l, h int) int {

	pivot := A[l]
	i := l
	j := h
	// 8, 7, 4, 6, 3, 5, 2
	for i < j {

		for i < len(A) && A[i] <= pivot {
			i++
		}

		for j > 0 && A[j] > pivot {
			j--
		}

		if i < j {
			A[i], A[j] = A[j], A[i]
		}
	}
	A[l], A[j] = A[j], A[l]
	return j
}

func quickSort(A []int, l, h int) int {
	if l < 0 || h > len(A) || h < 0 {
		return OutOfBound
	}

	if l < h {
		mid := partition(A, l, h)
		quickSort(A, l, mid-1)
		quickSort(A, mid+1, h)
	}

	return 0
}
