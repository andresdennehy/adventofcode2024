package main

import (
	"adventofcode2024/utils"
	"bufio"
	"fmt"
	"os"
)

func DFS(chart [][]int, start utils.Position) int {
	// Returns number of 9's reached from the start position.

	var result int
	directions := []utils.Direction{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	stack := []utils.Position{start}
	height := len(chart)
	width := len(chart[0])
	visited := make(map[utils.Position]bool)
	for len(stack) > 0 {
		popped := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		visited[popped] = true

		if chart[popped.Row][popped.Col] == 9 {
			result++
			continue
		}

		for _, direction := range directions {
			next := utils.Position{Row: popped.Row + direction[0], Col: popped.Col + direction[1]}
			if next.Row >= 0 && next.Col >= 0 && next.Row < height && next.Col < width && chart[next.Row][next.Col]-chart[popped.Row][popped.Col] == 1 && !visited[next] {
				stack = append(stack, next)
			}
		}
	}

	fmt.Printf("Count is %d\n", result)
	return result
}

func main() {

	file, err := os.Open("input.txt")
	utils.Check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var result int
	var chart [][]int
	for scanner.Scan() {
		line := scanner.Text()
		var chartLine []int
		// Assuming nice input with heights <= 9!
		for _, char := range line {
			parsed := int(char - '0')
			chartLine = append(chartLine, parsed)
		}
		chart = append(chart, chartLine)
	}

	var height, width = len(chart), len(chart[0])
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if chart[i][j] == 0 {
				result += DFS(chart, utils.Position{Row: i, Col: j})
			}
		}
	}

	fmt.Printf("Result is %v\n", result)
}
