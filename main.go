package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const DIR = "/home/fernando/GolandProjects/advent-20/"

func readInput(isTest bool) string {
	bytes, err := os.ReadFile(DIR + "inputs/01.txt")
	if err != nil {
		fmt.Println(err)
		panic("Error reading file.")
	}

	return string(bytes)
}

func main() {
	input := readInput(false)
	values := strings.Split(input, "\n")
	remainders := make(map[int]bool)

	for _, valueStr := range values {
		value, ok := strconv.Atoi(valueStr)
		if ok != nil {
			panic("Error converting to int")
		}

		currentRemainder := 2020 - value

		if _, ok := remainders[currentRemainder]; ok {
			fmt.Println(value * currentRemainder)
			break
		} else {
			remainders[value] = true
		}
	}
}
