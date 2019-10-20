package main

import "testing"

type sumTestCase struct {
	input  []int
	output int
}

var sumTestCases = []sumTestCase{
	{
		input:  []int{1, 2, 3},
		output: 6,
	},
	{
		input: []int{5},
		output: 5,
	},
	{
		input:[]int{5,6,7,8,},
		output: 26,
	},
}

func TestSum(t *testing.T) {
	for caseNum, c := range sumTestCases {
		result := sum(c.input)
		if result != c.output {
			t.Error("case: ", caseNum,
				"\r\nresult: ", result,
				"\r\nexpected: ", c.output)
		}
	}

}
