package main

import "testing"

type testCase struct {
	inputSlice []int
	item       int
	output     int
}

var testCases = []testCase{
	{
		inputSlice: []int{1, 2, 3, 4, 5},
		item:       5,
		output:     4,
	},
	{
		inputSlice: []int{1, 2, 3, 4, 5},
		item:       10,
		output:     -1,
	},
}

func Test_binarySearch(t *testing.T) {
	for caseNum, c := range testCases {
		result := binarySearch(c.inputSlice, c.item)
		if result != c.output {
			t.Error("case: ", caseNum,
				"\r\nresult: ", result,
				"\r\nexpected: ", c.output)
		}
	}
}
