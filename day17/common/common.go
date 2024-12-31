package common

import (
	"adventofcode2024/utils"
)

func RunProgram(instructions []int, registerA, registerB, registerC int, includeZero bool) []int {
	var output []int
	var i int
	for i < len(instructions) {
		var operation = instructions[i]
		var operand = instructions[i+1]
		var comboOperand int
		switch operand {
		case 0, 1, 2, 3:
			comboOperand = operand
		case 4:
			comboOperand = registerA
		case 5:
			comboOperand = registerB
		case 6:
			comboOperand = registerC
		}

		// Perform operation
		switch operation {
		case 0:
			if includeZero {
				registerA = registerA / (utils.PowInt(2, comboOperand))
			}
		case 1:
			registerB = registerB ^ operand
		case 2:
			registerB = comboOperand % 8
		case 3:
			if registerA != 0 && includeZero {
				i = operand
				continue
			}
		case 4:
			registerB = registerB ^ registerC
		case 5:
			output = append(output, comboOperand%8)
		case 6:
			registerB = registerA / (utils.PowInt(2, comboOperand))
		case 7:
			registerC = registerA / (utils.PowInt(2, comboOperand))
		}

		i += 2
	}

	return output
}
