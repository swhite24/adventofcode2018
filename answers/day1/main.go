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

	sum := 0

	if lines, err = util.GetInputLines(1); err != nil {
		log.Fatalf("failed to load input file: %s", err.Error())
	}

	for _, line := range lines {
		sum = process(sum, line)
	}

	fmt.Println("sum: ", sum)
}

func process(sum int, line string) int {
	op := line[:1]
	val, _ := strconv.Atoi(line[1:])

	if op == "+" {
		return sum + val
	}
	return sum - val
}
