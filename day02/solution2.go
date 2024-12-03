package main

import (
	"bufio"
	"fmt"
	"os"
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

func safe(reports []int) bool {
	var direction = reports[len(reports)-1] - reports[0]
	for i := 1; i < len(reports); i++ {
		difference := reports[i] - reports[i-1]
		// Use XOR to avoid multiplying
		if abs(difference) == 0 || abs(difference) > 3 || difference^direction < 0 {
			return false
		}
	}
	return true
}

func main() {

	file, err := os.Open("input.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var result int

	for scanner.Scan() {
		line := scanner.Text()

		var reports []int
		numbers := strings.Split(line, " ")
		for i := 0; i < len(numbers); i++ {
			report, err := strconv.Atoi(numbers[i])
			check(err)
			reports = append(reports, report)
		}

		for i, _ := range reports {
			var r = make([]int, len(reports))
			copy(r, reports)
			if safe(append(r[0:i], r[i+1:]...)) {
				result++
				break
			}
		}

	}

	fmt.Printf("Result is %v\n", result)
}
