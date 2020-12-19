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
	"byr":   []int{1920, 2002},
	"iyr":   []int{2010, 2020},
	"eyr":   []int{2020, 2030},
	"hgtcm": []int{150, 193},
	"hgtin": []int{59, 76},
}

var regexRules = map[string]*regexp.Regexp{
	"hcl": regexp.MustCompile(`#[0-9a-f]{6}`),
	"ecl": regexp.MustCompile(`amb|blu|brn|gry|grn|hzl|oth`),
	"pid": regexp.MustCompile(`^\d{9}$`),
}

func checkHairColour(hex string) bool {
	if regexRules["hcl"].MatchString(hex) == true {
		return true
	}
	return false
}

func checkEyeColour(key string) bool {
	if regexRules["ecl"].MatchString(key) == true {
		return true
	}
	return false
}

func checkPassportID(pid string) bool {
	if regexRules["pid"].MatchString(pid) == true {
		return true
	}
	return false
}

// checkHeight passes height strings (with cm or in)
func checkHeight(height string) bool {
	cmRegex := regexp.MustCompile(`(\d+)cm$`)
	inRegex := regexp.MustCompile(`(\d+)in$`)
	cmMatch := cmRegex.FindStringSubmatch(height)
	inMatch := inRegex.FindStringSubmatch(height)

	if len(cmMatch) > 0 {
		// do cm stuff
		cm, err := strconv.Atoi(cmMatch[1])
		if err != nil {
			log.Fatal(err)
		}
		if cm < intRules["hgtcm"][0] || cm > intRules["hgtcm"][1] {
			return false
		}
	}

	if len(inMatch) > 0 {
		// do in stuff
		in, err := strconv.Atoi(inMatch[1])
		if err != nil {
			log.Fatal(err)
		}
		if in < intRules["hgtin"][0] || in > intRules["hgtin"][1] {
			return false
		}
	}
	return true
}

func checkYearRange(key string, year string) bool {
	yearInt, err := strconv.Atoi(year)
	if err != nil {
		log.Fatal(err)
	}
	if yearInt < intRules[key][0] || yearInt > intRules[key][1] {
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
			"hcl": true, // hcl (Hair Color) - a # followed by exactly six characters 1-9 or a-f.
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
			} else if name == "byr" || name == "iyr" || name == "eyr" {
				if checkYearRange(name, strings.SplitN(field, ":", 2)[1]) != true {
					log.Printf("Rejecting name %v for %v", field, ps)
					continue POOP
				}
			} else if name == "hgt" {
				if checkHeight(strings.SplitN(field, ":", 2)[1]) != true {
					log.Printf("Rejecting name %v for %v", field, ps)
					continue POOP
				}
			} else if name == "hcl" {
				if checkHairColour(strings.SplitN(field, ":", 2)[1]) != true {
					log.Printf("Rejecting name %v for %v", field, ps)
					continue POOP
				}
			} else if name == "ecl" {
				if checkEyeColour(strings.SplitN(field, ":", 2)[1]) != true {
					log.Printf("Rejecting name %v for %v", field, ps)
					continue POOP
				}
			} else if name == "pid" {
				if checkPassportID(strings.SplitN(field, ":", 2)[1]) != true {
					log.Printf("Rejecting name %v for %v", field, ps)
					continue POOP
				}
			}
			delete(want, name)
		}
		delete(want, "cid")
		if len(want) == 0 {
			log.Printf("Passport correct: %v - number: %d", ps, valid)
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
