package main


func sum(nums []int)(s int){
	if len(nums)==0{
		return 0
	}
	s=s+nums[0]
	return s+sum(nums[1:])
}