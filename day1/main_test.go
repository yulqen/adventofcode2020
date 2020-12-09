package main

import (
	"testing"
)

// ALGORITHM
//
// Mark and fix num[0] <- GOT FIRST OPERAND
// loop from num[1] to len(num) {
//	in this loop mark and fix num1[0] <- GOT SECOND OPERAND
//	  loop from num1[1] to len(num1)
//	   in this loop mark and fix num2[0] <- GOT THIRD OPERAND
//		addOperands(num[1], num1[0], num2[0]) : GET RESULT and TEST
// }

func TestEndToEnd(t *testing.T) {
	cases := []struct {
		id     int
		input  []int
		result int
	}{
		{1, []int{10, 2000, 10}, 200000}, // 10 * 1 * 2000 * 9 = 180000
		{2, []int{20, 1990, 10, 10, 200}, 398000},
		{3, []int{10000, 10, 2000, 10, 285, 45}, 200000},
	}
	for _, c := range cases {
		if addUp(c.input) != c.result {
			t.Errorf("The test %d expected %d, got %d", c.id, c.result, addUp(c.input))
		}
	}
}
