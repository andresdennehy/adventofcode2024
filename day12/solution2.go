package main

import (
	"adventofcode2024/utils"
	"bufio"
	"fmt"
	"os"
)

type Plot struct {
	tiles []utils.Position
}

var directions = []utils.Direction{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}

func isInSlice(positions *[]utils.Position, value utils.Position) bool {
	for _, position := range *positions {
		if value == position {
			return true
		}
	}
	return false
}

func bfsPlot(chart *[]string, start utils.Position, visited *map[utils.Position]bool) []utils.Position {
	// Returns a map with all tiles of the plot with the character given by start.
	character := (*chart)[start.Row][start.Col]
	queue := []utils.Position{{Row: start.Row, Col: start.Col}}
	result := []utils.Position{{Row: start.Row, Col: start.Col}}
	(*visited)[start] = true
	for len(queue) > 0 {
		popped := queue[0]
		queue = queue[1:]
		for _, direction := range directions {
			neighbor := utils.Position{Row: popped.Row + direction[0], Col: popped.Col + direction[1]}
			if neighbor.Row < 0 || neighbor.Col < 0 ||
				neighbor.Row >= len(*chart) || neighbor.Col >= len((*chart)[0]) {
				continue
			}
			if (*visited)[neighbor] {
				continue
			}
			if (*chart)[neighbor.Row][neighbor.Col] != character {
				continue
			}
			queue = append(queue, neighbor)
			(*visited)[neighbor] = true
			result = append(result, neighbor)
		}
	}
	return result
}

func findPlots(chart *[]string) map[utils.Position]*Plot {
	// Takes the input and returns a map of positions to its corresponding plot.
	plots := make(map[utils.Position]*Plot)
	visited := make(map[utils.Position]bool)
	for i := 0; i < len(*chart); i++ {
		for j := 0; j < len((*chart)[0]); j++ {
			tile := utils.Position{Row: i, Col: j}
			if visited[tile] {
				continue
			}
			plotTiles := bfsPlot(chart, tile, &visited)
			plots[tile] = &Plot{tiles: plotTiles}
		}
	}
	return plots
}

func main() {

	file, err := os.Open("input.txt")
	utils.Check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var garden []string
	for scanner.Scan() {
		line := scanner.Text()
		garden = append(garden, line)
	}

	var result int
	plots := findPlots(&garden)
	for _, plot := range plots {
		fmt.Println(plot.tiles)
		area := len(plot.tiles)
		var sides int
		for _, tile := range plot.tiles {
			for i, direction := range directions {
				newTile := utils.Position{Row: tile.Row + direction[0], Col: tile.Col + direction[1]}
				if isInSlice(&plot.tiles, newTile) {
					continue
				}
				sideDirection := directions[(i+1)%4]
				sideTile := utils.Position{
					Row: tile.Row + sideDirection[0], Col: tile.Col + sideDirection[1],
				}
				if isInSlice(&plot.tiles, sideTile) {
					sideNP := utils.Position{
						Row: sideTile.Row + direction[0], Col: sideTile.Col + direction[1],
					}
					if !isInSlice(&plot.tiles, sideNP) {
						continue
					}
				}
				sides++
			}
		}
		fmt.Printf("Plot has area %d and %d sides\n", area, sides)
		result += area * sides
	}

	fmt.Printf("Result is %d\n", result)

}
