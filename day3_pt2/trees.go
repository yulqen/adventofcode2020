package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	// give us a lot of bytes...
	buf, err := ioutil.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(strings.Trim(string(buf), "\n"), "\n")

	var trees [5]int

	for i, line := range lines {
		if line[(1*i)%len(line)] == '#' {
			trees[0]++
		}
		if line[(3*i)%len(line)] == '#' {
			trees[1]++
		}
		if line[(5*i)%len(line)] == '#' {
			trees[2]++
		}
		if line[(7*i)%len(line)] == '#' {
			trees[3]++
		}
		if i%2 == 0 && line[(i/2)%len(line)] == '#' {
			fmt.Println(i)
			trees[4]++
		}
	}
	fmt.Println(trees)
	fmt.Println(trees[0] * trees[1] * trees[2] * trees[3] * trees[4])
}
