package main

import (
	"fmt"
	"log"

	"github.com/swhite24/adventofcode2018/pkg/util"
)

func main() {
	var err error
	var lines []string

	if lines, err = util.GetInputLines(2); err != nil {
		log.Fatalf("failed to read input: %s", err.Error())
	}

	fmt.Println("checksum: ", getchecksum(lines))
}

func getchecksum(lines []string) int {
	var twocount int
	var threecount int
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
	return twocount * threecount
}

func process(line string) map[string]int {
	legend := map[string]int{}
	for _, c := range line {
		legend[string(c)]++
	}
	return legend
}
