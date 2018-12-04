package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/swhite24/adventofcode2018/pkg/util"
)

func main() {
	var err error
	var lines []string

	if lines, err = util.GetInputLines(1); err != nil {
		log.Fatalf("failed to load input file: %s", err.Error())
	}

	fmt.Println("sum: ", findSum(lines))
	fmt.Println("first dup: ", findDup(lines))
}

func findSum(lines []string) int {
	sum := 0
	for _, line := range lines {
		sum = process(sum, line)
	}
	return sum
}

func findDup(lines []string) int {
	var dup *int
	sum := 0
	frequencies := map[int]bool{0: true}

main:
	for {
		for _, line := range lines {
			sum = process(sum, line)
			if frequencies[sum] {
				dup = &sum
				break main
			}
			frequencies[sum] = true
		}
	}
	return *dup
}

func process(sum int, line string) int {
	op := line[:1]
	val, _ := strconv.Atoi(line[1:])

	if op == "+" {
		return sum + val
	}
	return sum - val
}
