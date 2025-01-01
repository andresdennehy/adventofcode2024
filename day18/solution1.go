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
		if current == 1024 {
			break
		}
		var row, col int
		_, err := fmt.Sscanf(scanner.Text(), "%d,%d", &col, &row)
		utils.Check(err)
		bytes[utils.Position{Row: row, Col: col}] = true
		current++
	}

	fmt.Printf("Result is %v\n", common.BFS(utils.Position{}, utils.Position{Row: 70, Col: 70}, bytes))

}
