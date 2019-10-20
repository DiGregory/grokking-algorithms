package main

func binarySearch(mySlice []int, item int) (int) {
	low := 0
	high := len(mySlice) - 1
	mid := low + high
	for low <= high {
		mid = low + high
		guess := mySlice[mid]
		switch {
		case guess == item:
			return mid

		case guess > item:
			high = mid - 1
		case guess < item:
			low = mid + 1
		}
	}
	return -1

}

