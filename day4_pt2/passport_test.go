package main

import "testing"

func TestBirthYearFilter(t *testing.T) {
	cases := []struct {
		id     int
		key    string
		year   string
		result bool
	}{
		{1, "byr", "1920", true},
		{2, "byr", "2002", true},
		{3, "byr", "2003", false},
	}
	for _, c := range cases {
		if checkYearRange(c.key, c.year) != c.result {
			t.Errorf("The test %d failed; expected %t but got %t",
				c.id, c.result, checkYearRange(c.key, c.year))
		}
	}
}

func TestIssueYearFilter(t *testing.T) {
	cases := []struct {
		id     int
		key    string
		year   string
		result bool
	}{
		{1, "iyr", "2010", true},
		{2, "iyr", "2020", true},
		{3, "iyr", "2021", false},
		{4, "iyr", "1999", false},
		{5, "iyr", "2009", false},
	}
	for _, c := range cases {
		if checkYearRange(c.key, c.year) != c.result {
			t.Errorf("The test %d failed; expected %t but got %t",
				c.id, c.result, checkYearRange(c.key, c.year))
		}
	}
}

func TestExprYearFilter(t *testing.T) {
	cases := []struct {
		id     int
		key    string
		year   string
		result bool
	}{
		{1, "eyr", "2010", false},
		{2, "eyr", "2020", true},
		{3, "eyr", "2021", true},
		{4, "eyr", "1999", false},
		{5, "eyr", "2030", true},
	}
	for _, c := range cases {
		if checkYearRange(c.key, c.year) != c.result {
			t.Errorf("The test %d failed; expected %t but got %t",
				c.id, c.result, checkYearRange(c.key, c.year))
		}
	}
}

func TestHeightFilter(t *testing.T) {
	cases := []struct {
		id     int
		height string
		result bool
	}{
		{1, "150cm", true},
		{2, "193cm", true},
		{3, "59in", true},
		{4, "76in", true},
		{5, "149cm", false},
		{6, "194cm", false},
		{7, "58in", false},
		{8, "77in", false},
		{9, "10202in", false},
		{10, "1cm", false},
	}
	for _, c := range cases {
		if checkHeight(c.height) != c.result {
			t.Errorf("The test %d failed; expected %t but got %t",
				c.id, c.result, checkHeight(c.height))
		}
	}
}

func TestHairColorFilter(t *testing.T) {
	cases := []struct {
		id     int
		hex    string
		result bool
	}{
		{1, "#123abc", true},
		{2, "#123abz", false},
		{3, "123abz", false},
		{4, "#1abz", false},
	}

	for _, c := range cases {
		if checkHairColour(c.hex) != c.result {
			t.Errorf("The test %d failed; expected %t but got %t",
				c.id, c.result, checkHairColour(c.hex))
		}
	}
}

func TestEyeColourFilter(t *testing.T) {
	cases := []struct {
		id     int
		key    string
		result bool
	}{
		{1, "amb", true},
		{2, "blu", true},
		{3, "brn", true},
		{4, "gry", true},
		{5, "grn", true},
		{6, "hzl", true},
		{7, "oth", true},
		{8, "tit", false},
	}
	for _, c := range cases {
		if checkEyeColour(c.key) != c.result {
			t.Errorf("The test %d failed; expected %t but got %t",
				c.id, c.result, checkEyeColour(c.key))
		}
	}

}
