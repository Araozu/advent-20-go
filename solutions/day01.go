package solutions

import (
	"fmt"
	"strconv"
	"strings"
)

func Day01Part01() {
	input := ReadInput("01", true)
	values := strings.Split(input, "\n")
	remainders := make(map[int]bool)

	for _, valueStr := range values {
		value, err := strconv.Atoi(valueStr)
		if err != nil {
			panic("Error converting to int")
		}

		currentRemainder := 2020 - value

		if _, err := remainders[currentRemainder]; err {
			fmt.Println(value * currentRemainder)
			break
		} else {
			remainders[value] = true
		}
	}
}

func Day01Part02() {
	input := ReadInput("01", false)
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
					fmt.Println(values[i] * values[j] * values[k])
					return
				}
			}
		}
	}
}
