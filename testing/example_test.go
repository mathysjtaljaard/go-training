package main

import "testing"

func TestElse(t *testing.T) {

	type something struct {
		data   []int
		answer int
	}

	tests := []something{
		{[]int{1, 2, 3, 4}, 10},
		{[]int{1, 2, 3, 5}, 11},
		{[]int{1, 2, 3, 6}, 11},
	}

	for i, test := range tests {
		checkAnswer(i, test.answer, MySum(test.data...), t)
	}
}

func checkAnswer(index, expected, answer int, t *testing.T) {
	if expected != answer {
		t.Error("Test #", index, "Expected", expected, "Got", answer)
	}
}
