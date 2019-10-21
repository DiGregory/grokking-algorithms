package main

import (
	"testing"
)

type testCase struct {
	input  []int
	output []int
}

var testCases = []testCase{
	{
		input:  []int{2, 3, 1, 5, 4},
		output: []int{1, 2, 3, 4, 5},
	},
	{
		input:  []int{6, 2, 6, 1, 3, 7, 4},
		output: []int{1, 2, 3, 4, 6, 6, 7},
	},
	{
		input:  []int{13, 1, 8, 5, 1, 2, 3},
		output: []int{1, 1, 2, 3, 5, 8, 13},
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

func TestSelectionSort(t *testing.T) {
	for caseNum, c := range testCases {
		result := selectionSort(c.input)
		if !slEqual(c.output, result) {
			t.Error("case: ", caseNum,
				"\r\nresult: ", result,
				"\r\nexpected: ", c.output)
		}
	}

}

