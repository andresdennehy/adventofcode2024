package main

import (
	"adventofcode2024/day20/common"
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
	var walls = make(map[utils.Position]bool)
	var height int
	var start, end utils.Position
	var width int
	for scanner.Scan() {
		if height == 0 {
			width = len(scanner.Text())
		}
		for currentCol, char := range scanner.Text() {
			switch char {
			case '#':
				walls[utils.Position{Row: height, Col: currentCol}] = true
			case 'S':
				start = utils.Position{Row: height, Col: currentCol}
			case 'E':
				end = utils.Position{Row: height, Col: currentCol}
			}
		}
		height++
	}

	var normalPath = common.BFS(walls, height, width, start, end)
	var counter int
	for i := 0; i < len(normalPath)-100; i++ {
		for j := i + 100; j < len(normalPath); j++ {
			var cheatLen = utils.Abs(normalPath[i].Row-normalPath[j].Row) + utils.Abs(normalPath[i].Col-normalPath[j].Col)
			if cheatLen <= 20 && ((j-i)-cheatLen >= 100) {
				counter++
			}
		}
	}

	fmt.Printf("There are %d cheats that would save at least 100 picoseconds.", counter)

}
