package solutions

import (
	"fmt"
	"math/big"
	"os"
	"slices"
)

func Day10Part01(isTest bool) int {
	numbers := ReadAndMapInt("10", isTest)
	slices.Sort(numbers)

	// The highest jolt value
	highestValue := numbers[len(numbers)-1]
	// Add the built-in adapter jolt value
	numbers = append(numbers, highestValue+3)

	jolt1Difference := 0
	jolt3Difference := 0
	currentJoltValue := 0

	for _, value := range numbers {
		difference := value - currentJoltValue
		switch difference {
		case 1:
			jolt1Difference += 1
		case 3:
			jolt3Difference += 1
		default:
			fmt.Printf("Found an invalid jolt difference: %d %d %d \n", currentJoltValue, value, difference)
			os.Exit(-1)
		}
		currentJoltValue = value
	}

	return jolt1Difference * jolt3Difference
}

// This section assumes that there are no 2-jolt difference connections.
// In my input and the test inputs there were no 2-jolt differences.
func Day10Part02(isTest bool) int {
	numbers := ReadAndMapInt("10", isTest)
	slices.Sort(numbers)

	// The highest jolt value
	highestValue := numbers[len(numbers)-1]
	// Add the built-in adapter jolt value
	numbers = append(numbers, highestValue+3)

	// 1. Group all groups of 3 or more consecutive numbers.
	counts := groupConsecutiveNumbers(numbers)

	// 2. Count the number of paths between these groups
	pathCounts := countPathsByConsecutiveNumbers(counts)

	// 3. Multiply these number of paths to get the final answer.
	answer := big.NewInt(1)
	for _, value := range pathCounts {
		answer.Mul(answer, big.NewInt(value))
	}

	fmt.Println("day 10 answer: " + answer.Text(10))

	return -1
}

// Returns the length of each group of 3 or more consecutive numbers on the input slice.
//
// E.g.: This input:
//
//	0-3-4-5-9-12-13-14-15-16-19-20
//
// Would split into:
//
//	0 3-4-5 9 12-13-14-15-16 19 20
//
// And then return a slice:
// [3, 5]
// Which contains the length of each group of 3 or more consecutive numbers
func groupConsecutiveNumbers(input []int) []int {
	result := make([]int, 0)

	previousNumber := 0
	consecutiveCount := 1

	for _, value := range input {
		if value == previousNumber+1 {
			consecutiveCount += 1
		} else {

			// If there was a consecutive group previously, append to the slice
			if consecutiveCount >= 3 {
				result = append(result, consecutiveCount)
			}

			// Reset the consecutive count
			consecutiveCount = 1
		}

		previousNumber = value
	}

	return result
}

// Transforms a slice of amount of consecutive numbers into a slice of path counts
//
// Let x be a single amount of consecutive numbers:
//
// The amount of paths in a group of x numbers equals
// the amount of paths in groups x-1 + x-2 + x-3.
//
// The initial state is as follows:
//
// - when x=1 then paths=1
//
// - when x=2 then paths=1
//
// - when x=3 then paths=2
//
// - when x=2 then paths = 2+1+1 = 4, and so on.
//
// This operation is mapped over every number in the input slice
func countPathsByConsecutiveNumbers(input []int) []int64 {
	// Initial state for memoization
	pathCount := []int64{1, 1, 2, 4, 7, 13, 24}
	result := make([]int64, len(input))

	for idx, value := range input {
		result[idx] = consecutiveCountToPathCount(value, &pathCount)
	}

	return result
}

func consecutiveCountToPathCount(consecutiveCount int, pathCountSlice *[]int64) int64 {
	currentMaxAmount := len(*pathCountSlice)

	// If the index is in the slice, return it
	if currentMaxAmount >= consecutiveCount {
		return (*pathCountSlice)[consecutiveCount-1]
	}

	// Otherwise, build the slice up to the index
	for i := currentMaxAmount; i < consecutiveCount; i += 1 {
		nextValue := (*pathCountSlice)[consecutiveCount-1] + (*pathCountSlice)[consecutiveCount-2] + (*pathCountSlice)[consecutiveCount-3]
		*pathCountSlice = append(*pathCountSlice, nextValue)
	}

	return (*pathCountSlice)[consecutiveCount-1]
}
