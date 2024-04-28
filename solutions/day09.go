package solutions

import (
	"math"
	"strconv"
	"strings"
)

func findInvalidNumber(preambleSize int, n *[]int) int {
	numbers := *n
	dataSize := len(numbers)

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

	preambleSize := 25
	if isTest {
		preambleSize = 5
	}

	return findInvalidNumber(preambleSize, &numbers)
}

func Day09Part02(isTest bool) int {
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

	preambleSize := 25
	if isTest {
		preambleSize = 5
	}

	invalidNumber := findInvalidNumber(preambleSize, &numbers)

	for i := 0; i < len(groups); i += 1 {
		firstNumber := numbers[i]
		sum := firstNumber

		resultNumbers := make([]int, 0)
		resultNumbers = append(resultNumbers, firstNumber)

		for j := i + 1; j < len(groups); j += 1 {
			currentNumber := numbers[j]
			sum += currentNumber

			if sum < invalidNumber {
				resultNumbers = append(resultNumbers, currentNumber)
			} else if sum == invalidNumber {
				resultNumbers = append(resultNumbers, currentNumber)

				// Search the smallest & largest numbers in the slice

				smallest := math.MaxInt
				largest := math.MinInt

				for _, value := range resultNumbers {
					if value < smallest {
						smallest = value
					}
					if value > largest {
						largest = value
					}
				}

				return smallest + largest
			} else {
				break
			}
		}
	}

	return -1
}
