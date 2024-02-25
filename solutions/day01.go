package solutions

import (
	"strconv"
	"strings"
)

func Day01Part01(isTest bool) int {
	input := ReadInput("01", isTest)
	values := strings.Split(input, "\n")
	remainders := make(map[int]bool)

	for _, valueStr := range values {
		value, err := strconv.Atoi(valueStr)
		if err != nil {
			panic("Error converting to int")
		}

		currentRemainder := 2020 - value

		if _, err := remainders[currentRemainder]; err {
			return value * currentRemainder
		} else {
			remainders[value] = true
		}
	}

	return -1
}

func Day01Part02(isTest bool) int {
	input := ReadInput("01", isTest)
	strValues := strings.Split(input, "\n")
	values := make([]int, len(strValues))
	for i, valueStr := range strValues {
		value, err := strconv.Atoi(valueStr)
		if err != nil {
			panic("Error converting to int")
		}
		values[i] = value
	}

	arrLen := len(values)
	for i := 0; i < arrLen; i += 1 {
		for j := i + 1; j < arrLen; j += 1 {
			for k := j + 1; k < arrLen; k += 1 {
				if values[i]+values[j]+values[k] == 2020 {
					return values[i] * values[j] * values[k]
				}
			}
		}
	}

	return -1
}
