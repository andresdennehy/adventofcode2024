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
	block := -1
	var blockLength, checksum int

	for i := len(parsedLine) - 1; i >= 0; i-- {
		if parsedLine[i] == block {
			blockLength++
		} else {
			if block != -1 {
				freeSpace := 0
				for j := 0; j <= i; j++ {
					if parsedLine[j] == -1 {
						freeSpace++
					} else {
						freeSpace = 0
					}
					if freeSpace == blockLength {
						for k := range freeSpace {
							parsedLine[j-k], parsedLine[i+k+1] = parsedLine[i+k+1], parsedLine[j-k]
						}
						break
					}
				}
			}

			block = parsedLine[i]
			blockLength = 1
		}
	}

	for i, file := range parsedLine {
		if file != -1 {
			checksum += i * file
		}
	}

	fmt.Printf("Result is %v\n", checksum)
}
