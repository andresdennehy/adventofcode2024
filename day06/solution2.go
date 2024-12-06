package main

import (
	"adventofcode2024/utils"
	"bufio"
	"fmt"
	"os"
)

func isLoop(lab *[][]rune, pos utils.Position, direction utils.Direction) bool {
	var visited = make(map[utils.DirectedPosition]bool)
	state := utils.DirectedPosition{Pos: pos, Dir: direction}
	visited[state] = true
	height := len(*lab)
	width := len((*lab)[0])
	for {
		// Determine next position
		nextPos := utils.Position{Row: pos.Row + direction[0], Col: pos.Col + direction[1]}

		// Check if the guard is out of bounds
		if nextPos.Row < 0 || nextPos.Row >= height || nextPos.Col < 0 || nextPos.Col >= width {
			return false // No loop, guard left the lab
		}

		// Check if the state was visited
		state = utils.DirectedPosition{Pos: nextPos, Dir: direction}
		if visited[state] {
			return true // Loop detected
		}

		// Mark the state as visited
		visited[state] = true

		// Handle movement or obstacle
		if (*lab)[nextPos.Row][nextPos.Col] == '#' {
			// Rotate 90° if there's an obstacle
			direction = utils.Direction{direction[1], -direction[0]}
		} else {
			// Move forward
			pos = nextPos
		}
	}
}

func main() {

	file, err := os.Open("input.txt")
	utils.Check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lab [][]rune
	for scanner.Scan() {
		line := scanner.Text()
		lab = append(lab, []rune(line))
	}

	height := len(lab)
	width := len(lab[0])
	var guardRow, guardCol int
	// First, find the guard
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if lab[i][j] == '^' {
				guardRow, guardCol = i, j
				fmt.Printf("Start is %d,%d\n", i, j)
				break
			}
		}
	}

	var result int
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if lab[i][j] != '^' && lab[i][j] != '#' {
				lab[i][j] = '#'
				if isLoop(&lab, utils.Position{Row: guardRow, Col: guardCol}, utils.Direction{-1, 0}) {
					result++
				}
				lab[i][j] = '.'
			}
		}
	}

	fmt.Printf("Result is %v\n", result)
}
