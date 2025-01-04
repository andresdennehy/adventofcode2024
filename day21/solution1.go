package main

import (
	"adventofcode2024/day21/common"
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
	var result int

	for scanner.Scan() {
		var code = scanner.Text()
		codeNumber, err := strconv.Atoi(strings.Replace(code, "A", "", -1))
		utils.Check(err)
		var cache = make(map[common.CacheKey]int)
		var movesMap = make(map[common.MoveKey][][]rune)
		result += common.SequenceLength([]rune(code), 0, 2, cache, movesMap) * codeNumber
	}
	fmt.Printf("Result is %d\n", result)
}
