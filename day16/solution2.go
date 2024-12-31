package main

import (
	"adventofcode2024/utils"
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
)

type state struct {
	position  utils.Position
	direction utils.Direction
	score     int
	path      map[utils.Position]bool
}

type posDir struct {
	position  utils.Position
	direction utils.Direction
}

func findBestPathTiles(walls map[utils.Position]bool, start, end utils.Position) int {
	/*
		We need to modify breadth-first search, because not every node has the same weight.
	*/
	var queue []state
	var processed = make(map[posDir]int)
	var enqueue = func(item state) {
		var extendedPath = make(map[utils.Position]bool)
		for p, _ := range item.path {
			extendedPath[p] = true
		}
		extendedPath[item.position] = true
		for i := 0; i < len(queue); i++ {
			if queue[i].score >= item.score {
				queue = slices.Insert(queue, i, state{
					item.position,
					item.direction,
					item.score,
					extendedPath,
				})
				return
			}
		}
		queue = append(queue, state{
			item.position,
			item.direction,
			item.score,
			extendedPath,
		})
	}

	var minScore = math.MaxInt
	startPath := make(map[utils.Position]bool)
	startPath[start] = true
	enqueue(state{start, utils.Direction{0, 1}, 0, startPath})
	seats := make(map[utils.Position]bool) // Number of keys in this map is the result
	for len(queue) > 0 {
		popped := queue[0]
		queue = queue[1:]
		if popped.score > minScore {
			continue
		}
		poppedPosDir := posDir{popped.position, popped.direction}
		if v, ok := processed[poppedPosDir]; ok && popped.score > v {
			continue
		}
		processed[poppedPosDir] = popped.score
		if popped.position == end {
			if popped.score < minScore {
				minScore = popped.score
			}
			for p, _ := range popped.path {
				seats[p] = true
			}
		}
		var next = utils.Position{
			Row: popped.position.Row + popped.direction[0],
			Col: popped.position.Col + popped.direction[1],
		}
		if !walls[next] {
			enqueue(state{next, popped.direction, popped.score + 1, popped.path})
		}
		enqueue(state{
			popped.position,
			utils.Direction{-popped.direction[1], popped.direction[0]},
			popped.score + 1000,
			popped.path,
		})
		enqueue(state{
			popped.position,
			utils.Direction{popped.direction[1], -popped.direction[0]},
			popped.score + 1000,
			popped.path,
		})
	}
	return len(seats)
}

func main() {

	file, err := os.Open("input.txt")
	utils.Check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	walls := make(map[utils.Position]bool)
	var start, end utils.Position
	var row int
	for scanner.Scan() {
		line := scanner.Text()
		for col, char := range line {
			currentPosition := utils.Position{Row: row, Col: col}
			switch char {
			case '#':
				walls[currentPosition] = true
			case 'S':
				start = currentPosition
			case 'E':
				end = currentPosition
			}
		}
		row++
	}

	fmt.Printf("Result is %d\n", findBestPathTiles(walls, start, end))

}
