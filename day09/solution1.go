package main

import (
	"adventofcode2024/day09/common"
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
	scanner.Scan()
	line := scanner.Text()
	parsedLine := common.ConvertToDisk(line)
	checksum := 0

	swapIndex := len(parsedLine) - 1
	for i := 0; i < len(parsedLine); i++ {
		if parsedLine[i] == -1 {
			for j := swapIndex; j > i; j-- {
				if parsedLine[j] != -1 {
					swapIndex = j
					parsedLine[i], parsedLine[j] = parsedLine[j], parsedLine[i]
					checksum += i * parsedLine[i]
					break
				}
			}
		} else {
			checksum += i * parsedLine[i]
		}
	}

	fmt.Printf("Result is %v\n", checksum)
}
