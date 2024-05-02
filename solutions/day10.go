package solutions

import (
	"fmt"
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
		case 2:
			continue
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

func Day10Part02(isTest bool) int {
	return -1
}
