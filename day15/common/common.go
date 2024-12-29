package common

import (
	"adventofcode2024/utils"
	"fmt"
)

var Directions = map[rune]utils.Direction{
	'^': {-1, 0},
	'>': {0, 1},
	'v': {1, 0},
	'<': {0, -1},
}

func FindRobot(chart *[][]rune) (utils.Position, bool) {
	for i := 0; i < len(*chart); i++ {
		for j := 0; j < len((*chart)[0]); j++ {
			if (*chart)[i][j] == '@' {
				(*chart)[i][j] = '.'
				return utils.Position{Row: i, Col: j}, true
			}
		}
	}
	return utils.Position{}, false
}

func CalculateScore(chart *[][]rune) int {
	var result int
	for i := 0; i < len(*chart); i++ {
		for j := 0; j < len((*chart)[0]); j++ {
			if (*chart)[i][j] == 'O' {
				result += 100*i + j
			}
		}
	}
	return result
}

func DisplayChart(chart *[][]rune) {
	for i := 0; i < len(*chart); i++ {
		var rowString string
		for _, char := range (*chart)[i] {
			rowString = rowString + string(char)
		}
		fmt.Println(rowString)
	}
}
