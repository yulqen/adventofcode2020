package main

import "testing"

func TestBirthYearFilter(t *testing.T) {
	cases := []struct {
		id     int
		year   string
		result bool
	}{
		{1, "1920", true},
		{2, "2002", true},
		{3, "2003", false},
	}
	for _, c := range cases {
		if checkBirthYear(c.year) != c.result {
			t.Errorf("The test %d failed; expected %t but got %t",
				c.id, c.result, checkBirthYear(c.year))
		}
	}
}
