package main

import (
	"adventofcode2024/utils"
	"bufio"
	"fmt"
	"os"
	"slices"
)

type queueItem struct {
	position  utils.Position
	direction utils.Direction
	score     int
}

type visitedItem struct {
	position  utils.Position
	direction utils.Direction
}

func BFS(walls map[utils.Position]bool, start, end utils.Position) int {
	/*
		We need to modify breadth-first search, because not every node has the same weight.
	*/
	var queue = []queueItem{{position: start, direction: utils.Direction{0, 1}, score: 0}}
	var enqueue = func(item queueItem) {
		// If there is an item in the queue with higher score, visit this
		// node before that one
		for i := 0; i < len(queue); i++ {
			if queue[i].score >= item.score {
				queue = slices.Insert(queue, i, item)
				return
			}
		}
		queue = append(queue, item)
	}
	visited := make(map[visitedItem]bool)
	for len(queue) > 0 {
		popped := queue[0]
		queue = queue[1:]
		if popped.position == end {
			return popped.score
		}
		if visited[visitedItem{popped.position, popped.direction}] {
			continue
		}
		visited[visitedItem{popped.position, popped.direction}] = true
		var next = utils.Position{
			Row: popped.position.Row + popped.direction[0],
			Col: popped.position.Col + popped.direction[1],
		}
		if !walls[next] {
			enqueue(queueItem{next, popped.direction, popped.score + 1})
		}
		enqueue(queueItem{popped.position, utils.Direction{-popped.direction[1], popped.direction[0]}, popped.score + 1000})
		enqueue(queueItem{popped.position, utils.Direction{popped.direction[1], -popped.direction[0]}, popped.score + 1000})
	}
	return 0
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

	fmt.Printf("Result is %d\n", BFS(walls, start, end))

}
