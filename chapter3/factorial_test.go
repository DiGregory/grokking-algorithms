package main

import "testing"

type testCase struct {
	input  int
	output int
}

var testCases = []testCase{
	{
		input:  5,
		output: 120,
	},
	{
		input:  0,
		output: 1,
	},
	{
		input:  20,
		output: 2432902008176640000,
	},
}

func TestFactorial(t *testing.T) {
	for caseNum, c := range testCases {
		result := factorial(c.input)
		if c.output != result {
			t.Error("case: ", caseNum,
				"\r\nresult: ", result,
				"\r\nexpected: ", c.output)
		}

	}
}
