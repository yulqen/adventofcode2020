package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {

	// for {
	// 	var line string
	// 	_, err := fmt.Scanf("%s", &line)
	// 	if err != nil {
	// 		if err != io.EOF {
	// 			log.Fatal(err)
	// 		}
	// 		break
	// 	}
	buf, err := ioutil.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(strings.Trim(string(buf), "\n"), "\n")

	trees := 0
	// Snaffled this algorithm from https://www.twitch.tv/videos/834273244
	// 59:00 in. Still not sure I understand it, but it works.
	for i, line := range lines {
		if line[(3*i)%len(line)] == '#' {
			fmt.Println("Index ", i)
			trees++
		}
	}

	// TODO - for Part 2, he mdempsky starts same video at 01:01.

	fmt.Println(trees)

	// My toil below:
	//
	// for j := 0; j <= len(line); j++ {
	// 	if (j%3 == 0) && (j != 0) && j >= min && string(line[j]) == "#" {
	// 		fmt.Printf("At index %d with %c\n", j, line[j])
	// 		min = j
	// 		trees++
	// 		break
	// 	}
	// }
}
