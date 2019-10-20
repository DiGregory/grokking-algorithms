package main

func biggestElem(mySlice []int) (b int) {
	b = mySlice[0]
	if len(mySlice) == 1 {
		return b
	}
	if biggestElem(mySlice[1:]) > b {
		return biggestElem(mySlice[1:])
	} else {
		return b
	}

}
