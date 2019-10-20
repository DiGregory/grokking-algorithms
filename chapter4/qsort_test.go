package main

import (
	"testing"
)

type qsortTestCase struct {
	input  []int
	output []int
}

var qsortTestCases = []qsortTestCase{
	{
		input:  []int{3, 2, 5},
		output: []int{2, 3, 5},
	},
	{
		input:  []int{5},
		output: []int{5},
	},
	{
		input:  []int{8, 6, 7, 5,},
		output: []int{5, 6, 7, 8},
	},
	{
		input:  []int{8, 6, 7, 5, 1, 5, 2, 3},
		output: []int{1, 2, 3, 5, 5, 6, 7, 8},
	},
}

func slEqual(a, b []int) (bool) {

	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
func TestQsort(t *testing.T) {
	for caseNum, c := range qsortTestCases {
		result := qsort(c.input)
		if !slEqual(c.output, result) {
			t.Error("case: ", caseNum,
				"\r\nresult: ", result,
				"\r\nexpected: ", c.output)
		}
	}

}
