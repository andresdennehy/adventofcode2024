package main

import (
	"adventofcode2024/utils"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func isPossible(design string, towelPatterns []string) bool {
	if design == "" {
		return true
	}
	for _, pattern := range towelPatterns {
		if len(design) >= len(pattern) && design[:len(pattern)] == pattern {
			fmt.Printf("Comparing %s and %s\n", design[:len(pattern)], pattern)
			if isPossible(design[len(pattern):], towelPatterns) {
				return true
			}
		}
	}
	return false
}

func main() {

	file, err := os.Open("input.txt")
	utils.Check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	var towelPatterns = strings.Split(scanner.Text(), ", ")
	scanner.Scan()

	var possibleCount int
	for scanner.Scan() {
		if isPossible(scanner.Text(), towelPatterns) {
			possibleCount++
		}
	}

	fmt.Printf("Available patterns: %v\n", towelPatterns)
	fmt.Printf("Result is: %d\n", possibleCount)

}
