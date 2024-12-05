package main

import (
	"adventofcode2024/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func checkIncorrectUpdate(update string, rules *map[string][]string) int {
	// This is O(n^2). Most probably not efficient enough
	pages := strings.Split(update, ",")
	swapped := false
	for i := 0; i < len(pages); i++ {
		for j := i + 1; j < len(pages); j++ {
			// If there is a rule forbidding this order
			for _, goesBefore := range (*rules)[pages[j]] {
				if goesBefore == pages[i] {
					temp := pages[i]
					pages[i] = pages[j]
					pages[j] = temp
					swapped = true
				}
			}
		}
	}
	if swapped {
		middleValue, err := strconv.Atoi(pages[len(pages)/2])
		utils.Check(err)
		return middleValue
	}
	return 0
}

func main() {

	file, err := os.Open("input.txt")
	utils.Check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var result int
	rules := make(map[string][]string)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		rule := strings.Split(line, "|")
		var before, after = rule[0], rule[1]
		rules[before] = append(rules[before], after)
	}

	for scanner.Scan() {
		line := scanner.Text()
		result += checkIncorrectUpdate(line, &rules)
	}

	fmt.Printf("Result is %v\n", result)
}
