package main

import "testing"

type bigElemTestCase struct {
	input  []int
	output int
}

var bigElemTestCases = []bigElemTestCase{
	{
		input:  []int{1, 2, 3},
		output: 3,
	},
	{
		input: []int{5},
		output: 5,
	},
	{
		input:[]int{8,6,7,5,},
		output: 8,
	},
}

func TestBiggestElem(t *testing.T) {
	for caseNum, c := range bigElemTestCases {
		result := biggestElem(c.input)
		if result != c.output {
			t.Error("case: ", caseNum,
				"\r\nresult: ", result,
				"\r\nexpected: ", c.output)
		}
	}

}
