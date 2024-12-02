package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	file, err := os.Open("input.txt")
	check(err)
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

	frequencies := make(map[int]int)
	for _, value := range rightList {
		frequencies[value] += 1
	}

	for _, value := range leftList {
		result += value * frequencies[value]
	}
	fmt.Printf("Result is %v\n", result)
}
