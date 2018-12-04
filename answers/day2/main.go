package main

import (
	"fmt"
	"log"

	"github.com/swhite24/adventofcode2018/pkg/util"
)

func main() {
	var err error
	var lines []string
	var twocount int
	var threecount int

	if lines, err = util.GetInputLines(2); err != nil {
		log.Fatalf("failed to read input: %s", err.Error())
	}

	for _, line := range lines {
		var hastwo bool
		var hasthree bool
		legend := process(line)
		for _, count := range legend {
			if count == 2 {
				hastwo = true
			}
			if count == 3 {
				hasthree = true
			}
		}

		if hastwo {
			twocount++
		}
		if hasthree {
			threecount++
		}
	}

	fmt.Println("twocount: ", twocount)
	fmt.Println("threecount: ", threecount)
	fmt.Println("checksum: ", twocount*threecount)
}

func process(line string) map[string]int {
	legend := map[string]int{}
	for _, c := range line {
		legend[string(c)]++
	}
	return legend
}
