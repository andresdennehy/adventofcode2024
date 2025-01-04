package main

import (
	"adventofcode2024/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func iterateNumber(initial int, iterations int) int {
	var number = initial
	for i := 0; i < iterations; i++ {
		number = ((number * 64) ^ number) % 16777216
		number = ((number / 32) ^ number) % 16777216
		number = ((number * 2048) ^ number) % 16777216
	}
	fmt.Printf("%d: %d\n", initial, number)
	return number
}

func main() {

	file, err := os.Open("input.txt")
	utils.Check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var result int

	for scanner.Scan() {
		number, err := strconv.Atoi(scanner.Text())
		utils.Check(err)
		result += iterateNumber(number, 2000)
	}
	fmt.Printf("Result is %d\n", result)
}
