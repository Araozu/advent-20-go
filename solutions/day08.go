package solutions

import (
	"fmt"
	"strconv"
	"strings"
)

type InstructionStore struct {
	instruction    Instruction
	value          int
	executionCount int
}

type Instruction int

const (
	nop Instruction = iota
	acc
	jmp
)

func parseInstruction(line string) (Instruction, int) {
	instr := strings.Split(line, " ")[0]
	valueStr := line[4:]
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		fmt.Println(err)
		panic("expected a number, got `" + valueStr + "`")
	}

	var instruction Instruction

	if instr == "nop" {
		instruction = nop
	} else if instr == "acc" {
		instruction = acc
	} else if instr == "jmp" {
		instruction = jmp
	} else {
		panic("Invalid instruction found: " + instr)
	}

	return instruction, value
}

func Day08Part01(isTest bool) int {
	input := ReadInput("08", isTest)
	groups := strings.Split(input, "\n")

	instructionArray := make([]*InstructionStore, len(groups))

	for idx, line := range groups {
		instr, value := parseInstruction(line)
		instructionArray[idx] = &InstructionStore{instruction: instr, value: value, executionCount: 0}
	}

	accumulator := 0
	instructionPointer := 0

	for {
		currentInstruction := instructionArray[instructionPointer]
		inst := currentInstruction.instruction
		value := currentInstruction.value
		executionCount := currentInstruction.executionCount

		if executionCount != 0 {
			return accumulator
		}

		switch inst {
		case nop:
			currentInstruction.executionCount += 1
			instructionPointer += 1
		case acc:
			currentInstruction.executionCount += 1
			accumulator += value
			instructionPointer += 1
		case jmp:
			currentInstruction.executionCount += 1
			instructionPointer += value
		default:
			panic("Found an invalid instruction while trying to execute it. " + strconv.Itoa(int(inst)))
		}
	}
}

func Day08Part02(isTest bool) int {
	input := ReadInput("08", isTest)
	groups := strings.Split(input, "\n")

	instructionArray := make([]*InstructionStore, len(groups))

	for idx, line := range groups {
		instr, value := parseInstruction(line)
		instructionArray[idx] = &InstructionStore{instruction: instr, value: value, executionCount: 0}
	}

	for currentExecutionCount := 0; currentExecutionCount < len(instructionArray); currentExecutionCount += 1 {
		accumulator := 0
		instructionPointer := 0
		jmpNopCount := 0

		for {
			// success condition
			if instructionPointer == len(instructionArray) {
				return accumulator
			}

			currentInstruction := instructionArray[instructionPointer]
			inst := currentInstruction.instruction
			value := currentInstruction.value
			executionCount := currentInstruction.executionCount

			// If a loop is detected, break this loop
			if executionCount == currentExecutionCount+1 {
				break
			}

			// Check if the intruction is jmp or not, if it's the nth such instr, and change them if so)
			if inst == nop || inst == jmp {
				if jmpNopCount == currentExecutionCount {
					// switch jmp & nop
					if inst == nop {
						inst = jmp
					} else if inst == jmp {
						inst = nop
					}
				}

				jmpNopCount += 1
			}

			switch inst {
			case nop:
				currentInstruction.executionCount = currentExecutionCount + 1
				instructionPointer += 1
			case acc:
				currentInstruction.executionCount = currentExecutionCount + 1
				accumulator += value
				instructionPointer += 1
			case jmp:
				currentInstruction.executionCount = currentExecutionCount + 1
				instructionPointer += value
			default:
				panic("Found an invalid instruction while trying to execute it. " + strconv.Itoa(int(inst)))
			}
		}
	}

	return -1
}
