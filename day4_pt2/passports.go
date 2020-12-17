package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

// You can continue to ignore the cid field, but each other field
// has strict rules about what values are valid for automatic validation:

//     byr (Birth Year) - four digits; at least 1920 and at most 2002.
//     iyr (Issue Year) - four digits; at least 2010 and at most 2020.
//     eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
//     hgt (Height) - a number followed by either cm or in:
//         If cm, the number must be at least 150 and at most 193.
//         If in, the number must be at least 59 and at most 76.
//     hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
//     ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
//     pid (Passport ID) - a nine-digit number, including leading zeroes.
//     cid (Country ID) - ignored, missing or not.

// Your job is to count the passports where all required fields are both
// present and valid according to the above rules. Here are some example values:

var intRules = map[string][]int{
	"byr": []int{1920, 2002},
	"iyr": []int{2010, 2020},
	"eyr": []int{2020, 2030},
	"hgt": []int{150, 193},
}

var regexRules = map[string]*regexp.Regexp{
	"hgt": regexp.MustCompile(`(\d{4}cm|\d{4}in)`),
	"hcl": regexp.MustCompile(`#[0-9a-f]{6}`),
	"ecl": regexp.MustCompile(`amb|blu|bru|gry|grn|hzl|oth`),
	"pid": regexp.MustCompile(`\d{9}`),
}

func checkBirthYear(year string) bool {
	yearInt, err := strconv.Atoi(year)
	if err != nil {
		log.Fatal(err)
	}
	if yearInt < intRules["byr"][0] || yearInt > intRules["byr"][1] {
		return false
	}
	return true
}

func run(buf []byte) int {

	var valid int

	cleaned := strings.Split(strings.Trim(string(buf), "\n"), "\n\n")

	var passports []string
	for _, p := range cleaned {
		passports = append(passports, strings.ReplaceAll(p, "\n", " "))
	}

POOP:
	for _, ps := range passports {
		want := map[string]bool{
			"byr": true, // byr (Birth Year) - four digits; at least 1920 and at most 2002.
			"iyr": true, // iyr (Issue Year) - four digits; at least 2010 and at most 2020.
			"eyr": true, // eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
			"hgt": true, // hgt (Height) - a number followed by either cm or in: If cm, the number must be at least 150 and at most 193. If in, the number must be at least 59 and at most 76.
			"hcl": true, // hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
			"ecl": true, // ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
			"pid": true, // pid (Passport ID) - a nine-digit number, including leading zeroes.
			"cid": true, // cid (Country ID) - ignored, missing or not.
		}

		fields := strings.Split(ps, " ")

		for _, field := range fields {
			name := strings.SplitN(field, ":", 2)[0]
			if _, ok := want[name]; !ok {
				fmt.Println("missing:", name)
				continue POOP
			}
			delete(want, name)
		}
		delete(want, "cid")
		if len(want) == 0 {

			valid++
		}
	}
	return valid
}

func main() {
	buf, err := ioutil.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(run(buf))
}
