package main

import (
	"adventofcode2024/day13/common"
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
	var machines []common.Machine
	for scanner.Scan() {
		var newMachine common.Machine
		_, err := fmt.Sscanf(scanner.Text(), "Button A: X+%d, Y+%d\n", &newMachine.ButtonA.X, &newMachine.ButtonA.Y)
		utils.Check(err)
		scanner.Scan()
		_, err = fmt.Sscanf(scanner.Text(), "Button B: X+%d, Y+%d\n", &newMachine.ButtonB.X, &newMachine.ButtonB.Y)
		utils.Check(err)
		scanner.Scan()
		_, err = fmt.Sscanf(scanner.Text(), "Prize: X=%d, Y=%d\n", &newMachine.Prize.X, &newMachine.Prize.Y)
		utils.Check(err)
		scanner.Scan()

		machines = append(machines, newMachine)
	}

	var result int
	for _, machine := range machines {
		result += common.CalculateCost(&machine)
	}

	fmt.Printf("Result is %d\n", result)

}
