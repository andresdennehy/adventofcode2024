package main

import (
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
	var lab []string
	for scanner.Scan() {
		line := scanner.Text()
		lab = append(lab, line)
	}

	height := len(lab)
	width := len(lab[0])

	var result int = 1 // Current position is visited
	var currRow, currCol int
	// First, find the guard
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if lab[i][j] == '^' {
				currRow, currCol = i, j
				fmt.Printf("Start is %d,%d\n", i, j)
				break
			}
		}
	}

	// With more time, I could convert the map to a graph data structure
	// For now, do the old iterating way
	direction := utils.Direction{-1, 0} // Up, dRow=-1, dCol=0
	for currRow >= 0 && currRow < height && currCol >= 0 && currCol < width {
		if lab[currRow][currCol] == '#' {
			currRow -= direction[0]
			currCol -= direction[1]
			// Rotate 90°
			direction = utils.Direction{direction[1], -direction[0]}
			continue
		}
		if lab[currRow][currCol] == '.' {
			line := []rune(lab[currRow])
			line[currCol] = 'X' // Filler
			lab[currRow] = string(line)
			result++
		}
		currRow += direction[0]
		currCol += direction[1]
	}

	fmt.Printf("Result is %v\n", result)
}
