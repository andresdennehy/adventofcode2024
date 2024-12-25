package main

import (
	"adventofcode2024/day14/common"
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
	var robots []common.Robot
	for scanner.Scan() {
		var newRobot common.Robot
		_, err := fmt.Sscanf(scanner.Text(), "p=%d,%d v=%d,%d", &newRobot.Col, &newRobot.Row, &newRobot.DCol, &newRobot.DRow)
		utils.Check(err)
		robots = append(robots, newRobot)
	}

	for i := 0; i < 100; i++ {
		for i := range robots {
			robots[i].Move()
		}
	}

	middleRow := common.SpaceRows / 2
	middleCol := common.SpaceCols / 2
	counts := make(map[utils.Position]int)
	for _, robot := range robots {
		if robot.Row == middleRow || robot.Col == middleCol {
			continue
		}
		rowQuadrant := 0
		colQuadrant := 0

		// Determine row quadrant
		if robot.Row > middleRow {
			rowQuadrant = 1
		}

		// Determine column quadrant
		if robot.Col > middleCol {
			colQuadrant = 1
		}
		counts[utils.Position{Row: rowQuadrant, Col: colQuadrant}]++
	}

	var result = 1
	for _, factor := range counts {
		result *= factor
	}

	fmt.Printf("Result is %d\n", result)

}
