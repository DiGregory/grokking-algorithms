package main

import "testing"

type counterTestCase struct {
	input  []int
	output int
}

var counterTestCases = []counterTestCase{
	{
		input:  []int{1, 2, 3},
		output: 3,
	},
	{
		input: []int{5},
		output: 1,
	},
	{
		input:[]int{5,6,7,8,},
		output: 4,
	},
}

func TestCounter(t *testing.T) {
	for caseNum, c := range counterTestCases {
		result := countOfElem(c.input)
		if result != c.output {
			t.Error("case: ", caseNum,
				"\r\nresult: ", result,
				"\r\nexpected: ", c.output)
		}
	}

}
