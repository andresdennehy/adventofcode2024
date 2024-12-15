package common

import (
	"adventofcode2024/utils"
	"strconv"
)

func Blink(stones map[int]int) map[int]int {
	newStones := make(map[int]int)
	for stone, count := range stones {
		if stone == 0 {
			newStones[1] = newStones[1] + count
		} else {
			stoneString := strconv.Itoa(stone)
			numberOfDigits := len(stoneString)
			if numberOfDigits%2 == 0 {
				first, err := strconv.Atoi(stoneString[:numberOfDigits/2])
				utils.Check(err)
				second, err := strconv.Atoi(stoneString[numberOfDigits/2:])
				utils.Check(err)
				newStones[first] = newStones[first] + count
				newStones[second] = newStones[second] + count
			} else {
				newStones[stone*2024] = newStones[stone*2024] + count
			}
		}
	}
	return newStones
}

func CountStones(stones map[int]int) int {
	var result int
	for _, count := range stones {
		result += count
	}
	return result
}
