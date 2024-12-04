package main

import (
	"adventofcode2024/utils"
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {

	file, err := os.Open("input.txt")
	utils.Check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var leftList, rightList []int
	var result int

	for scanner.Scan() {
		line := scanner.Text()
		var left, right int

		numbers := strings.Split(line, "   ")
		left, _ = strconv.Atoi(numbers[0])
		right, _ = strconv.Atoi(numbers[1])

		leftList = append(leftList, left)
		rightList = append(rightList, right)
	}

	slices.Sort(leftList)
	slices.Sort(rightList)

	for i, _ := range leftList {
		result += utils.Abs(leftList[i] - rightList[i])
	}
	fmt.Printf("Result is %v\n", result)
}
