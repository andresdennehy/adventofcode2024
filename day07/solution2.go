package main

import (
	"adventofcode2024/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func evaluateWithConcat(number int, operands []int) bool {
	// This is really slow - O(3^n)!!
	// There is a DP solution. I didn't have enough time to work it out
	if len(operands) == 1 {
		return operands[0] == number
	}
	var withSum, withProduct, withConcatenation bool
	withSum = evaluateWithConcat(number, append([]int{operands[0] + operands[1]}, operands[2:]...))
	withProduct = evaluateWithConcat(number, append([]int{operands[0] * operands[1]}, operands[2:]...))

	concatenated, _ := strconv.Atoi(fmt.Sprintf("%d%d", operands[0], operands[1]))
	withConcatenation = evaluateWithConcat(number, append([]int{concatenated}, operands[2:]...))

	return withSum || withProduct || withConcatenation
}

func main() {

	file, err := os.Open("input.txt")
	utils.Check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var result int
	for scanner.Scan() {
		line := scanner.Text()
		splitLine := strings.Split(line, ": ")
		number, err := strconv.Atoi(splitLine[0])
		utils.Check(err)
		var operands []int
		for _, operand := range strings.Split(splitLine[1], " ") {
			parsed, err := strconv.Atoi(operand)
			utils.Check(err)
			operands = append(operands, parsed)
		}
		if evaluateWithConcat(number, operands) {
			result += number
		}
	}

	fmt.Printf("Result is %v\n", result)
}
