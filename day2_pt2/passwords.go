package main

import (
	"fmt"
	"io"
	"log"
	"strings"
)

func main() {

	var valid1, valid2 int

	for {
		var lo, hi int
		var ch byte
		var password string

		_, err := fmt.Scanf("%d-%d %c: %s", &lo, &hi, &ch, &password)
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}

		n := strings.Count(password, string(ch))
		if n >= lo && n <= hi {
			valid1++
		}

		if (password[lo-1] == ch) != (password[hi-1] == ch) {
			valid2++
		}

	}
	fmt.Println(valid1)
	fmt.Println(valid2)

	// var valid int

	// data := adventloader.ReadFile("input")

	// for _, line := range data {
	// 	var ch byte
	// 	var password string

	// 	_, err := strconv.Atoi(string(line[0]))
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	_, err = strconv.Atoi(string(line[2]))
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	ch = line[4]
	// 	password = string(line[4:])

	// 	n := strings.Count(password, string(ch))
	// 	fmt.Println(n)

	// }
	// fmt.Println(valid)
}
