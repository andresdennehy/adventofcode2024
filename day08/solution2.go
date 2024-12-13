package main

import (
	"adventofcode2024/utils"
	"bufio"
	"fmt"
	"os"
)

func markHarmonicAntinodes(positions []utils.Position, mapWidth, mapHeight int, antinodeMap *map[utils.Position]bool) {
	for i, position := range positions {
		for _, otherPosition := range positions[i+1:] {
			fmt.Printf("Comparing %v and %v\n", position, otherPosition)
			rowDist := position.Row - otherPosition.Row
			colDist := position.Col - otherPosition.Col

			// Mark the two antennas as antinodes
			(*antinodeMap)[position] = true
			(*antinodeMap)[otherPosition] = true

			currRow, currCol := position.Row+rowDist, position.Col+colDist
			// For the current position
			for currRow >= 0 && currRow < mapHeight && currCol >= 0 && currCol < mapWidth {
				fmt.Printf("Setting %d,%d\n", currRow, currCol)
				(*antinodeMap)[utils.Position{Row: currRow, Col: currCol}] = true
				currRow += rowDist
				currCol += colDist
			}

			currRow, currCol = otherPosition.Row-rowDist, otherPosition.Col-colDist
			// For the other position
			for currRow >= 0 && currRow < mapHeight && currCol >= 0 && currCol < mapWidth {
				fmt.Printf("Setting %d,%d\n", currRow, currCol)
				(*antinodeMap)[utils.Position{Row: currRow, Col: currCol}] = true
				currRow -= rowDist
				currCol -= colDist
			}
		}
	}
}

func main() {

	file, err := os.Open("input.txt")
	utils.Check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	antennaMap := make(map[rune][]utils.Position)
	var currentLine int
	var width, height int
	for scanner.Scan() {
		line := scanner.Text()
		if currentLine == 0 {
			width = len(line)
		}
		// Assuming all characters are ASCII!
		for col, char := range line {
			if char != '.' {
				antennaMap[char] = append(antennaMap[char], utils.Position{Row: currentLine, Col: col})
			}
		}
		currentLine++
	}
	height = currentLine

	antinodeMap := make(map[utils.Position]bool)
	for _, charPositions := range antennaMap {
		markHarmonicAntinodes(charPositions, width, height, &antinodeMap)
	}

	fmt.Printf("Result is %v\n", len(antinodeMap))
}
