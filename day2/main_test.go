package main

import (
	"fmt"
	"testing"
)

// 1-3 a: abcde
// 1-3 b: cdefg
// 2-9 c: ccccccccc

func TestvalidPasswordSingle(t *testing.T) {
	cases := []struct {
		id         int
		policyLine string
		valid      bool
	}{
		{1, "1-3 a: abcde", true},
		{2, "1-3 b: cdefg", false},
		{3, "2-9 c: ccccccccc", true},
	}
	for _, c := range cases {
		fmt.Println("In test - with ", c)
		if validPasswordSingle(c.policyLine) != c.valid {
			t.Errorf("The test %d expected %v but got %v", c.id, c.valid, validPasswordSingle(c.policyLine))
		}
	}
}
