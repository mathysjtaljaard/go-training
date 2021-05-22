package main

import (
	"fmt"
	"testing"
)

func TestMySum(t *testing.T) {

	fmt.Println(MySum(4, 4))

	type test struct {
		data   []int
		answer int
	}

	tests := []test{
		{[]int{21, 21}, 42},
		{[]int{4, 4}, 8},
	}

	for _, test := range tests {
		sum := MySum(test.data...)

		if sum != test.answer {
			t.Error("Expected", test.answer, "got", sum)
		}
	}
}
