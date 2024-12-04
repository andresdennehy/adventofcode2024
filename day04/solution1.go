package main

import (
	"adventofcode2024/utils"
	"bufio"
	"fmt"
	"os"
)

func checkForXmas(letters []string, row, col int, deltaY, deltaX int) bool {
	var spells string
	for i := 0; i < 4; i++ {
		if row < 0 || row > len(letters)-1 {
			return false
		}
		if col < 0 || col > len(letters[0])-1 {
			return false
		}
		spells += string(letters[row][col])
		row += deltaY
		col += deltaX
	}
	if spells == "XMAS" {
		return true
	}
	return false
}

func main() {

	file, err := os.Open("input.txt")
	utils.Check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	letters := make([]string, 0)
	var result int

	for scanner.Scan() {
		line := scanner.Text()
		letters = append(letters, line)
	}

	// Assuming correct input
	height := len(letters)
	width := len(letters[0])
	directions := [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}}
	// Run four operations per letter in the letters slice
	// to check for XMAS. This is O(4n) = O(n)
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			for _, direction := range directions {
				if checkForXmas(letters, i, j, direction[0], direction[1]) {
					result++
				}
			}
		}
	}

	fmt.Printf("Result is %v\n", result)
}
