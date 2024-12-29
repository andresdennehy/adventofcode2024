package main

import (
	"adventofcode2024/day15/common"
	"adventofcode2024/utils"
	"bufio"
	"fmt"
	"os"
)

func main() {

	file, err := os.Open("input.txt")
	utils.Check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var chart [][]rune
	// Scan initial map
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		var newRow []rune
		for _, char := range line {
			newRow = append(newRow, char)
		}
		chart = append(chart, newRow)
	}

	var movements []rune
	// Scan movements
	for scanner.Scan() {
		line := scanner.Text()
		for _, char := range line {
			movements = append(movements, char)
		}
	}

	current, found := common.FindRobot(&chart)
	if !found {
		fmt.Println("Robot not found! Check the input")
	}

	// Main movement loop
	for _, movement := range movements {
		direction := common.Directions[movement]
		neighbor := utils.Position{Row: current.Row + direction[0], Col: current.Col + direction[1]}
		switch chart[neighbor.Row][neighbor.Col] {
		case '#':
			continue
		case 'O':
			// Check if the robot can push boxes
			canMove := false
			nextNeighbor := utils.Position{Row: neighbor.Row + direction[0], Col: neighbor.Col + direction[1]}
			for chart[nextNeighbor.Row][nextNeighbor.Col] != '#' {
				if chart[nextNeighbor.Row][nextNeighbor.Col] == '.' {
					canMove = true
					break
				}
				nextNeighbor.Row += direction[0]
				nextNeighbor.Col += direction[1]
			}
			if canMove {
				chart[nextNeighbor.Row][nextNeighbor.Col] = 'O'
				chart[neighbor.Row][neighbor.Col] = '.'
				current.Row = neighbor.Row
				current.Col = neighbor.Col
			}
		case '.':
			current.Row = neighbor.Row
			current.Col = neighbor.Col
		}
	}

	fmt.Printf("Result is %d\n", common.CalculateScore(&chart))
	fmt.Printf("End position is %v\n", current)
	common.DisplayChart(&chart)

}
