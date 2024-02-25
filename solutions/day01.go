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
