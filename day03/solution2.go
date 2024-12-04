package main

import (
	"adventofcode2024/utils"
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {

	file, err := os.Open("input.txt")
	utils.Check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var result int
	enabled := true // Initially, mul instructions are enabled

	for scanner.Scan() {
		line := scanner.Text()

		// Regex for different patterns
		pattern := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)|do\(\)|don't\(\)`)
		matches := pattern.FindAllString(line, -1)

		for _, match := range matches {
			if match[0:3] == "mul" {
				var first, second int
				fmt.Sscanf(match, "mul(%d,%d)", &first, &second)
				if enabled {
					result += first * second
				}
			} else {
				enabled = match == "do()"
			}
		}
	}

	fmt.Printf("Result is %v\n", result)
}
