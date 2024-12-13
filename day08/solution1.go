package main

import (
	"adventofcode2024/utils"
	"bufio"
	"fmt"
	"os"
)

func markAntinodes(positions []utils.Position, mapWidth, mapHeight int, antinodeMap *map[utils.Position]bool) {
	for i, position := range positions {
		for _, otherPosition := range positions[i+1:] {
			rowDist := position.Row - otherPosition.Row
			colDist := position.Col - otherPosition.Col

			// For the current position
			if position.Row+rowDist >= 0 && position.Row+rowDist < mapHeight && position.Col+colDist >= 0 && position.Col+colDist < mapWidth {
				(*antinodeMap)[utils.Position{Row: position.Row + rowDist, Col: position.Col + colDist}] = true
			}
			// For the other position
			if otherPosition.Row-rowDist >= 0 && otherPosition.Row-rowDist < mapHeight && otherPosition.Col-colDist >= 0 && otherPosition.Col-colDist < mapWidth {
				(*antinodeMap)[utils.Position{Row: otherPosition.Row - rowDist, Col: otherPosition.Col - colDist}] = true
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
		markAntinodes(charPositions, width, height, &antinodeMap)
	}

	fmt.Printf("Result is %v\n", len(antinodeMap))
}
