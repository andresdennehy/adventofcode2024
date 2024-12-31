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

func findRegisterA(program []int, target []int, start int) int {
	fmt.Printf("Comparing %v to %v\n", program, target)
	if len(target) == 0 {
		return start
	}
	for digit := 0; digit < 8; digit++ {
		var a = (start << 3) | digit
		output := common.RunProgram(program, a, 0, 0, false) // They are 0 in the input
		if output[len(output)-1] == target[len(target)-1] {
			var prev = findRegisterA(program, target[:len(target)-1], a)
			if prev != -1 {
				return prev
			}
		}
	}
	return -1
}

func main() {

	file, err := os.Open("input.txt")
	utils.Check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var registerB, registerC int
	scanner.Scan()
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

	fmt.Printf("Result is %v\n", findRegisterA(program, program, 0))

}
