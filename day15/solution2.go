package main

import (
	"adventofcode2024/day15/common"
	"adventofcode2024/utils"
	"bufio"
	"fmt"
	"os"
)

func canMove(pos utils.Position, direction utils.Direction, boxes *map[utils.Position]bool, walls *map[utils.Position]bool) bool {
	if direction[1] == 0 { // Moving up or down
		if (*walls)[utils.Position{Row: pos.Row + direction[0], Col: pos.Col}] || (*walls)[utils.Position{Row: pos.Row + direction[0], Col: pos.Col + 1}] {
			return false
		}
		for i := -1; i < 2; i++ {
			next := utils.Position{Row: pos.Row + direction[0], Col: pos.Col + i}
			if (*boxes)[next] && !canMove(next, direction, boxes, walls) {
				return false
			}
		}
	}
	if direction[1] == 1 { // Moving right
		if (*walls)[utils.Position{Row: pos.Row, Col: pos.Col + 2}] {
			return false
		}
		next := utils.Position{Row: pos.Row, Col: pos.Col + 2}
		if (*boxes)[next] && !canMove(next, direction, boxes, walls) {
			return false
		}
	}
	if direction[1] == -1 { // Moving left
		if (*walls)[utils.Position{Row: pos.Row, Col: pos.Col - 1}] {
			return false
		}
		next := utils.Position{Row: pos.Row, Col: pos.Col - 2}
		if (*boxes)[next] && !canMove(next, direction, boxes, walls) {
			return false
		}
	}
	return true
}

func moveBoxes(pos utils.Position, direction utils.Direction, boxes *map[utils.Position]bool, walls *map[utils.Position]bool) {
	if direction[1] == 0 { // Moving up or down
		for i := -1; i < 2; i++ {
			next := utils.Position{Row: pos.Row + direction[0], Col: pos.Col + i}
			if (*boxes)[next] {
				moveBoxes(next, direction, boxes, walls)
			}
		}
		(*boxes)[utils.Position{Row: pos.Row + direction[0], Col: pos.Col}] = true
	}
	if direction[0] == 0 { // Moving left or right
		next := utils.Position{Row: pos.Row, Col: pos.Col + direction[1]*2}
		if (*boxes)[next] {
			moveBoxes(next, direction, boxes, walls)
		}
		(*boxes)[utils.Position{Row: pos.Row, Col: pos.Col + direction[1]}] = true
	}
	delete(*boxes, pos)
}

func main() {

	file, err := os.Open("input.txt")
	utils.Check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	boxes := make(map[utils.Position]bool)
	walls := make(map[utils.Position]bool)
	var current utils.Position
	var row, col int
	// Scan initial map
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		col = 0
		for _, char := range line {
			switch char {
			case '#':
				walls[utils.Position{Row: row, Col: 2 * col}] = true
				walls[utils.Position{Row: row, Col: 2*col + 1}] = true
			case 'O':
				boxes[utils.Position{Row: row, Col: 2 * col}] = true
			case '@':
				current = utils.Position{Row: row, Col: 2 * col}
			}
			col++
		}
		row++
	}

	var movements []rune
	// Scan movements
	for scanner.Scan() {
		line := scanner.Text()
		for _, char := range line {
			movements = append(movements, char)
		}
	}

	// Main movement loop
	for _, movement := range movements {
		direction := common.Directions[movement]
		neighbor := utils.Position{Row: current.Row + direction[0], Col: current.Col + direction[1]}
		if direction[1] == 0 { // Moving up or down
			var contiguous = utils.Position{Row: neighbor.Row, Col: neighbor.Col - 1}
			if boxes[neighbor] && canMove(neighbor, direction, &boxes, &walls) {
				moveBoxes(neighbor, direction, &boxes, &walls)
			} else if boxes[contiguous] && canMove(contiguous, direction, &boxes, &walls) {
				moveBoxes(contiguous, direction, &boxes, &walls)
			}
		} else if direction[1] == 1 { // Move right
			if boxes[neighbor] && canMove(neighbor, direction, &boxes, &walls) {
				moveBoxes(neighbor, direction, &boxes, &walls)
			}
		} else if direction[1] == -1 { // Move left
			var contiguous = utils.Position{Row: neighbor.Row, Col: neighbor.Col - 1}
			if boxes[contiguous] && canMove(contiguous, direction, &boxes, &walls) {
				moveBoxes(contiguous, direction, &boxes, &walls)
			}
		}
		if !walls[neighbor] && !boxes[neighbor] && !boxes[utils.Position{Row: neighbor.Row, Col: neighbor.Col - 1}] {
			current = neighbor
		}
	}

	var result int
	for coordinate, _ := range boxes {
		result += 100*coordinate.Row + coordinate.Col
	}

	fmt.Printf("Result is %d\n", result)
}
