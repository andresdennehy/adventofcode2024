package main

import (
	"adventofcode2024/utils"
	"bufio"
	"fmt"
	"os"
)

func checkForMas(letters []string, row, col int) bool {
	if row < 1 || row > len(letters)-2 {
		return false
	}
	if col < 1 || col > len(letters[0])-2 {
		return false
	}
	firstDiagonal := string([]byte{letters[row-1][col-1], letters[row][col], letters[row+1][col+1]})
	secondDiagonal := string([]byte{letters[row+1][col-1], letters[row][col], letters[row-1][col+1]})
	if (firstDiagonal == "MAS" || firstDiagonal == "SAM") && (secondDiagonal == "MAS" || secondDiagonal == "SAM") {
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
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if checkForMas(letters, i, j) {
				result++
			}
		}
	}

	fmt.Printf("Result is %v\n", result)
}
