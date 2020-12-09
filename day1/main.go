package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Find two entries that sum to 2020 from
// file 'input'

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readFile(path string) []int {
	file, err := os.Open(path)
	check(err)

	var nums []int

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		nums = append(nums, n)
	}
	return nums
}

func addUp(x []int) int {
	var operand1 int
	var operand2 int
	var operand3 int
	var res int

	for idx, n := range x {
		operand1 = n
		newSlice := x[idx+1:]
		for idx2, m := range newSlice {
			operand2 = m
			newSlice2 := newSlice[idx2+1:]
			for _, o := range newSlice2 {
				operand3 = o
				if operand1+operand2+operand3 == 2020 {
					res = (operand1 * operand2 * operand3)
					return res
				}
			}
		}
	}
	return 0
}

func run() int {
	fmt.Println("Advent of Code 2020 - Day 1")
	nums := readFile("input")
	return addUp(nums)
}

func main() {
	fmt.Println(run())
}
