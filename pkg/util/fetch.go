package util

import (
	"bufio"
	"fmt"
	"os"
)

// GetInputLines delivers a string slice filled with each line of the input file
func GetInputLines(day int) ([]string, error) {
	var err error
	var file *os.File
	lines := []string{}

	if file, err = os.Open(fmt.Sprintf("../../inputs/day%d.txt", day)); err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err = scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}
