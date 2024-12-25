package main

import (
	"adventofcode2024/day14/common"
	"adventofcode2024/utils"
	"bufio"
	"bytes"
	"compress/gzip"
	"fmt"
	"os"
	"strings"
)

func renderGrid(robots []common.Robot) string {
	grid := make([][]rune, common.SpaceRows)
	for i := range grid {
		grid[i] = make([]rune, common.SpaceCols)
		for j := range grid[i] {
			grid[i][j] = ' ' // Empty tiles
		}
	}
	for _, robot := range robots {
		grid[robot.Row][robot.Col] = '#' // Robot positions
	}

	var sb strings.Builder
	for _, row := range grid {
		sb.WriteString(string(row))
		sb.WriteRune('\n')
	}
	return sb.String()
}

func compressedSize(data string) int {
	var buf bytes.Buffer
	writer := gzip.NewWriter(&buf)
	_, _ = writer.Write([]byte(data))
	_ = writer.Close()
	return buf.Len()
}

func main() {

	file, err := os.Open("input.txt")
	utils.Check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var robots []common.Robot
	for scanner.Scan() {
		var newRobot common.Robot
		_, err := fmt.Sscanf(scanner.Text(), "p=%d,%d v=%d,%d", &newRobot.Col, &newRobot.Row, &newRobot.DCol, &newRobot.DRow)
		utils.Check(err)
		robots = append(robots, newRobot)
	}

	minSize := -1
	var minGrid string
	var minTime int

	for seconds := 0; ; seconds++ {
		// Render and compress the grid
		grid := renderGrid(robots)
		size := compressedSize(grid)

		// Check if this is the smallest compressed size
		if minSize == -1 || size < minSize {
			minSize = size
			minGrid = grid
			minTime = seconds

			fmt.Printf("New minimum size: %d at %d seconds\n", size, seconds)
			fmt.Println(strings.Repeat("-", 75))
			fmt.Println(grid)
			fmt.Println(strings.Repeat("-", 75))
		}

		for i := range robots {
			robots[i].Move()
		}

		if seconds > 20000 {
			break
		}
	}

	fmt.Printf("Minimum entropy grid occurred at %d seconds.\n", minTime)
	fmt.Println(minGrid)

}
