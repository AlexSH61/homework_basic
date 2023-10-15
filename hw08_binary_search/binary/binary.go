package binary

func BinarySearching(needle int, array []int) int {
	low, high := 0, len(array)-1
	for low <= high {
		med := (low + high) / 2
		if array[med] == needle {
			return med
		} else if array[med] < needle {
			low = med + 1
		} else {
			high = med - 1
		}
	}
	return -1
}
