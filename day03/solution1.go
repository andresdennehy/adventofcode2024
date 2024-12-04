package main

import (
	"adventofcode2024/utils"
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {

	file, err := os.Open("input.txt")
	utils.Check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var result int

	for scanner.Scan() {
		line := scanner.Text()
		re := regexp.MustCompile("mul\\((\\d{1,3}),(\\d{1,3})\\)")
		matches := re.FindAllStringSubmatch(line, -1)

		var mult int
		for _, match := range matches {
			first, _ := strconv.Atoi(match[1])
			second, _ := strconv.Atoi(match[2])
			mult += first * second
		}
		result += mult
	}

	fmt.Printf("Result is %v\n", result)
}
