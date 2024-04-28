package solutions

import (
	"strconv"
	"strings"
)

func Day09Part01(isTest bool) int {
	input := ReadInput("09", isTest)
	groups := strings.Split(input, "\n")

	dataSize := len(groups)
	numbers := make([]int, dataSize)

	for i, str := range groups {
		value, err := strconv.Atoi(str)
		if err != nil {
			panic(err)
		}

		numbers[i] = value
	}

	// loop
	preambleSize := 25

out:
	for i := 0; i < dataSize-preambleSize; i += 1 {
		sum := numbers[i+preambleSize]
		// A slice that stores: sum - i
		indexes := make([]int, preambleSize)

		// Iterate over every `preambleSize` items
		for j := 0; j < preambleSize; j += 1 {
			number := numbers[j+i]

			// On every iteration check if the current value is present in indexes
			for _, value := range indexes {
				// If so, a pair is found
				if value == number {
					continue out
				}
			}

			indexes = append(indexes, sum-number)
		}

		// When this is reached no sum was found
		return sum
	}

	return -1
}

func Day09Part02(isTest bool) int {
	// input := ReadInput("09", isTest)
	// groups := strings.Split(input, "\n")

	return -1
}
