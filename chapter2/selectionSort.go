package main

func findMin(mySlice []int) (int) {
	smallestIndex := 0
	smallest := mySlice[smallestIndex]
	for i, v := range mySlice {
		if v < smallest {
			smallest = v
			smallestIndex = i
		}
	}
	return smallestIndex
}

func selectionSort(mySlice []int) (sortedSlice []int) {

	for range mySlice {
		smallestIndex := findMin(mySlice)
		sortedSlice = append(sortedSlice, mySlice[smallestIndex])
		mySlice = append(mySlice[0:smallestIndex], mySlice[smallestIndex+1:]...)
	}
	return sortedSlice
}
