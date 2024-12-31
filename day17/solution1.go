package main

import (
	"adventofcode2024/day17/common"
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
	var registerA, registerB, registerC int
	scanner.Scan()
	_, err = fmt.Sscanf(scanner.Text(), "Register A: %d", &registerA)
	utils.Check(err)
	scanner.Scan()
	_, err = fmt.Sscanf(scanner.Text(), "Register B: %d", &registerB)
	utils.Check(err)
	scanner.Scan()
	_, err = fmt.Sscanf(scanner.Text(), "Register C: %d", &registerC)
	utils.Check(err)
	scanner.Scan()

	var program []int
	scanner.Scan()
	var programText string
	_, err = fmt.Sscanf(scanner.Text(), "Program: %s", &programText)
	utils.Check(err)
	for _, number := range strings.Split(programText, ",") {
		instruction, err := strconv.Atoi(number)
		utils.Check(err)
		program = append(program, instruction)
	}

	var output = common.RunProgram(program, registerA, registerB, registerC, true)
	var outputStr string
	for _, result := range output {
		if len(outputStr) > 0 {
			outputStr += ","
		}
		outputStr += strconv.Itoa(result)
	}
	fmt.Printf("Result is %v\n", outputStr)

}
