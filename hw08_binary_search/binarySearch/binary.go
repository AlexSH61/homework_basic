package binary

func Searching(needle int, array []int) int {
	low, high := 0, len(array)-1
	if needle >= low && needle <= high {
		for low <= high {
			med := (low + high) / 2
			switch array[med] == needle {
			case array[med] < needle:
				low = med + 1
			case array[med] > needle:
				high = med - 1
			default:
				return med
			}
		}
	}
	return -1
}
