package main

import (
	"adventofcode2024/day18/common"
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
	var bytes = make(map[utils.Position]bool)
	var current int
	for scanner.Scan() {
		var row, col int
		_, err := fmt.Sscanf(scanner.Text(), "%d,%d", &col, &row)
		utils.Check(err)
		bytes[utils.Position{Row: row, Col: col}] = true

		shortestPath := common.BFS(utils.Position{}, utils.Position{Row: 70, Col: 70}, bytes)
		if shortestPath == 0 {
			fmt.Printf("First byte to block the end is %d,%d, index %d\n", col, row, current)
			break
		}

		current++
	}

}
