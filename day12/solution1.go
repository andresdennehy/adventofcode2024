package main

import (
	"adventofcode2024/utils"
	"bufio"
	"fmt"
	"os"
)

func bfs(chart *[]string, start utils.Position, visited *map[utils.Position]bool) int {
	// Starts at a new character, expands as far as it can (i.e. traverses the graph)
	// and returns its perimeter times its area.
	character := (*chart)[start.Row][start.Col]
	queue := []utils.Position{{Row: start.Row, Col: start.Col}}
	directions := []utils.Direction{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	(*visited)[start] = true
	var perimeter int
	var area int
	for len(queue) > 0 {
		popped := queue[0]
		queue = queue[1:]
		area++
		for _, direction := range directions {
			neighbor := utils.Position{Row: popped.Row + direction[0], Col: popped.Col + direction[1]}
			if neighbor.Row < 0 || neighbor.Col < 0 ||
				neighbor.Row >= len(*chart) || neighbor.Col >= len((*chart)[0]) {
				perimeter++
			} else if (*chart)[neighbor.Row][neighbor.Col] != character {
				perimeter++
			} else if !(*visited)[neighbor] {
				queue = append(queue, neighbor)
				(*visited)[neighbor] = true
			}
		}
	}
	fmt.Printf("Area %d, Perimeter %d\n", area, perimeter)
	return perimeter * area
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

	visited := make(map[utils.Position]bool)
	var result int
	for i := 0; i < len(garden); i++ {
		for j := 0; j < len(garden[0]); j++ {
			if !visited[utils.Position{Row: i, Col: j}] {
				fmt.Printf("Found character %v in %v,%v\n", garden[i][j], i, j)
				result += bfs(&garden, utils.Position{Row: i, Col: j}, &visited)
			}
		}
	}

	fmt.Printf("Result is %d\n", result)

}
