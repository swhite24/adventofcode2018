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

	fmt.Println("checksum: ", getCheckSum(lines))
	fmt.Println("fabric box: ", getFabricBox(lines))
}

func getCheckSum(lines []string) int {
	var twocount int
	var threecount int
	for _, line := range lines {
		var hastwo bool
		var hasthree bool
		legend := getOccurrences(line)
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

func getFabricBox(lines []string) string {
	var common string
main:
	for _, line := range lines {
		for _, comp := range lines {
			common = ""
			if line == comp {
				continue
			}
			diff := 0
			for i := 0; i < len(line); i++ {
				c := line[i]
				if c != comp[i] {
					diff++
				} else {
					common += string(c)
				}
			}
			if diff == 1 {
				break main
			}
		}
	}

	return common
}

func getOccurrences(line string) map[string]int {
	legend := map[string]int{}
	for _, c := range line {
		legend[string(c)]++
	}
	return legend
}
