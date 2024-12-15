package common

import (
	"adventofcode2024/utils"
	"strconv"
)

func ConvertToDisk(input string) []int {
	var disk []int
	currentId := 0

	for i, char := range input {
		num, err := strconv.Atoi(string(char))
		utils.Check(err)

		if i%2 == 0 {
			for range num {
				disk = append(disk, currentId)
			}
			currentId++

		} else {
			for range num {
				disk = append(disk, -1)
			}
		}
	}

	return disk
}
