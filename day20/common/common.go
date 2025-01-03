package common

import (
	"adventofcode2024/utils"
	"slices"
)

type state struct {
	pos    utils.Position
	length int
	path   []utils.Position
}

func BFS(walls map[utils.Position]bool, height, width int, start utils.Position, end utils.Position) []utils.Position {
	var queue []state
	var enqueue = func(p utils.Position, picoseconds int, prevpath []utils.Position) {
		var path = make([]utils.Position, len(prevpath)+1)
		copy(path, prevpath)
		path[len(path)-1] = p
		for i := 0; i < len(queue); i++ {
			if queue[i].length >= picoseconds {
				queue = slices.Insert(queue, i, state{p, picoseconds, path})
				return
			}
		}
		queue = append(queue, state{p, picoseconds, path})
	}
	enqueue(start, 0, []utils.Position{})
	var processed = make(map[utils.Position]bool)
	for len(queue) > 0 {
		var p, picoseconds, path = queue[0].pos, queue[0].length, queue[0].path
		queue = queue[1:]
		if p == end {
			return path
		}
		if processed[p] {
			continue
		}
		processed[p] = true
		for _, d := range utils.UpRightDownLeft {
			var next = utils.Position{Row: p.Row + d[0], Col: p.Col + d[1]}
			if next.Row >= 0 && next.Row < height &&
				next.Col >= 0 && next.Col < width &&
				!walls[next] {
				enqueue(next, picoseconds+1, path)
			}
		}
	}
	return []utils.Position{}
}
