package algorithms

func heapify(arr []int, index, length int) {

	largest := index
	left := 2*index + 1
	right := 2*index + 2

	if left < length && arr[left] > arr[largest] {
		largest = left
	}
	if right < length && arr[right] > arr[largest] {
		largest = right
	}
	if largest != index {
		arr[largest], arr[index] = arr[index], arr[largest]
		heapify(arr, largest, length)
	}
}

func heapSort(arr []int) {
	size := len(arr)
	// make max heap
	for i := size / 2; i >= 0; i-- {
		heapify(arr, i, size)
	}

	for i := size - 1; i >= 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		heapify(arr, 0, i)
	}
}
