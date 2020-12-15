package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

// Example

// VALID - all eight fields present
// ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
// byr:1937 iyr:2017 cid:147 hgt:183cm

// INVALID - missing hgt field
// iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884
// hcl:#cfa07d byr:1929

// VALID - although missing cid field, still valid
// hcl:#ae17e1 iyr:2013
// eyr:2024
// ecl:brn pid:760753108 byr:1931
// hgt:179cm

// INVALID - missing cid (fine) and byr (not fine)
// hcl:#cfa07d eyr:2025 pid:166559648
// iyr:2011 ecl:brn hgt:59in

// Algorithm

// 1. Get batch as bytes
// 2. Split into password batches (starts at top and goes until finds two \n)
// 3. Trim all \n from that to create slice of all k/v strings
// 4. Split that on spaces to slice
// 5. Loop through each and test for presence of each required key, and count

func main() {
	buf, err := ioutil.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	var invalid int

	cleaned := strings.Split(string(buf), "\n\n")

	var passports []string
	for _, p := range cleaned {
		passports = append(passports, strings.ReplaceAll(p, "\n", " "))
	}

	total := len(passports)
	fmt.Println(total)

	for _, ps := range passports {
		sp := strings.Split(ps, " ")
		if len(sp) != 8 {
			invalid++
		}

	KVLOOP:
		for _, kv := range sp {
			k := strings.Split(kv, ":")
			switch k[0] {
			case "byr":
				continue KVLOOP
			}
		}
	}
}
