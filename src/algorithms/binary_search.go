package algorithms

const (
	NotFound   = -1
	OutOfBound = -2
)

func BinarySearch(array []int, given int, low int, high int) int {

	if high > len(array) || low < 0 {
		return OutOfBound
	}

	for low <= high {
		mid := (low + high) / 2

		if array[mid] == given {
			return mid
		} else if array[mid] > given {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return NotFound
}
