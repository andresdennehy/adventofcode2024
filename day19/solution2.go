package main

import (
	"adventofcode2024/utils"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func countWays(design string, towelPatterns []string, memo map[string]int) int {
	if design == "" {
		return 1
	}
	if result, found := memo[design]; found { // Check memoization
		return result
	}

	var ways int
	for _, pattern := range towelPatterns {
		if len(design) >= len(pattern) && design[:len(pattern)] == pattern {
			ways += countWays(design[len(pattern):], towelPatterns, memo)
		}
	}

	memo[design] = ways
	return ways
}

func main() {

	file, err := os.Open("input.txt")
	utils.Check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	var towelPatterns = strings.Split(scanner.Text(), ", ")
	scanner.Scan()

	var possibleWays int
	var memo = make(map[string]int) // Memoization map
	for scanner.Scan() {
		possibleWays += countWays(scanner.Text(), towelPatterns, memo)
	}

	fmt.Printf("Result is: %d\n", possibleWays)

}
