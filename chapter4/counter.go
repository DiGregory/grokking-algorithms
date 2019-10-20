package main


func countOfElem(mySlice []int)(count int){
	if len(mySlice)==1{
		return 1
	}
	count++
	return count+countOfElem(mySlice[1:])
}