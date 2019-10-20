package main

func qsort(mySlice []int) ([]int) {

	if len(mySlice) < 2 {
		return mySlice
	}
	pivot := mySlice[0]
	greater := []int{}
	less := []int{}
	for i := 1; i <= len(mySlice)-1; i++ {
		if mySlice[i] >= pivot {
			greater = append(greater, mySlice[i])
		} else if mySlice[i] < pivot {
			less = append(less, mySlice[i])
		}
	}
	less=qsort(less)
	less = append(less, pivot)
	greater=qsort(greater)

	return append(less, greater...)
}
