package main

import "testing"

type testCase struct {
	input  graph
	output node
}

var (
	you    = createNode("you", false)
	bob    = createNode("bob", false)
	alice  = createNode("alice", false)
	claire = createNode("claire", false)
	aliceT = createNode("alice", true)
)

var testCases = []testCase{
	{
		input: graph{nodes: map[node][]node{
			you: {
				alice,
				claire,
				bob,
			},
			alice:  nil,
			claire: nil,
			bob:    nil,
		}},
		output: node{},
	},
	{
		input: graph{nodes: map[node][]node{
			you: {
				aliceT,
				claire,
				bob,
			},
			aliceT: nil,
			claire: nil,
			bob:    nil,
		}},
		output: aliceT,
	},
	{
		input: graph{nodes: map[node][]node{
			you: {
				alice,
			},
			alice: {
				you,
			},
		}},
		output: node{},
	},
}

func TestGraph_FindMangoSeller(t *testing.T) {
	for caseNum, c := range testCases {
		result := c.input.FindMangoSeller(node{"you", false})
		if result != c.output {
			t.Error("case: ", caseNum,
				"\r\nresult: ", result,
				"\r\nexpected: ", c.output)
		}
	}
}
