package common

import (
	"adventofcode2024/utils"
)

type state struct {
	pos    utils.Position
	length int
}

func BFS(start, end utils.Position, bytes map[utils.Position]bool) int {
	var queue = []state{{pos: start, length: 0}}
	var visited = make(map[utils.Position]bool)
	visited[start] = true
	for len(queue) > 0 {
		var popped = queue[0]
		if popped.pos == end {
			return popped.length
		}
		queue = queue[1:]
		for _, direction := range utils.UpRightDownLeft {
			var next = state{utils.Position{Row: popped.pos.Row + direction[0], Col: popped.pos.Col + direction[1]}, popped.length + 1}
			var _, isByte = bytes[next.pos]
			if _, ok := visited[next.pos]; !ok &&
				!isByte &&
				next.pos.Row >= 0 &&
				next.pos.Row <= end.Row &&
				next.pos.Col >= 0 &&
				next.pos.Col <= end.Col {
				queue = append(queue, next)
				visited[next.pos] = true
			}
		}
	}
	return 0
}
