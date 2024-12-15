package main

import (
	"adventofcode2024/day11/common"
	"adventofcode2024/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	file, err := os.Open("input.txt")
	utils.Check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// Store sparsely in map for memory efficiency
	stones := make(map[int]int)
	for scanner.Scan() {
		line := scanner.Text()
		stringLine := strings.Split(line, " ")
		for _, stoneString := range stringLine {
			parsed, err := strconv.Atoi(stoneString)
			utils.Check(err)
			stones[parsed] = 1
		}
	}

	for i := 0; i < 25; i++ {
		stones = common.Blink(stones)
	}
	fmt.Printf("Result is %d\n", common.CountStones(stones))

}
