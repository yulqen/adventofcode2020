package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Password policy
//
// 1-3 a: abcde
// 1-3 b: cdefg
// 2-9 c: ccccccccc
//
// The second password does not comply because it requires a minimun of 1 b.

func validPasswordSingle(p string) bool {
	re := regexp.MustCompile(`^(\d+)\-(\d+) ([A-Za-z]): ([A-Za-z]+$)`)
	matches := re.FindStringSubmatch(p)
	min, _ := strconv.Atoi(matches[1])
	max, _ := strconv.Atoi(matches[2])
	target := matches[3]
	password := matches[4]
	pSlice := strings.Split(password, "")
	var count int
	for _, char := range pSlice {
		if char == target {
			count++
		}
	}
	if (count >= min) && (count <= max) {
		return true
	} else {
		return false
	}
}

func validPasswords(passwords []string) int {
	var valid int
	for _, p := range passwords {
		if validPasswordSingle(p) {
			valid++
		}
	}
	return valid
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readFile(path string) []string {
	file, err := os.Open(path)
	check(err)

	var ps []string

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		n := scanner.Text()
		ps = append(ps, n)
	}
	return ps
}

func run() int {
	fmt.Println("Advent of Code 2020 - Day 2")
	pwds := readFile("passwords")
	return validPasswords(pwds)
}

func main() {
	fmt.Println(run())
}
