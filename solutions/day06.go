package solutions

import (
	"strings"
)

func Day06Part01(isTest bool) int {
	input := ReadInput("06", isTest)
	groups := strings.Split(input, "\n\n")

	sum := 0

	for _, group := range groups {
		frequency := make(map[rune]int)

		for _, letter := range group {
			// remove non a-z character
			if letter < 97 || letter > 122 {
				continue
			}

			previousValue, _ := frequency[letter]
			frequency[letter] = previousValue + 1
		}

		sum += len(frequency)
	}

	return sum
}

func Day06Part02(isTest bool) int {
	input := ReadInput("06", isTest)
	groups := strings.Split(input, "\n\n")

	sum := 0

	for _, group := range groups {
		persons := strings.Split(group, "\n")
		personAmount := len(persons)

		frequency := make(map[rune]int)

		for _, person := range persons {
			for _, letter := range person {
				previousValue, _ := frequency[letter]
				frequency[letter] = previousValue + 1
			}
		}

		for _, value := range frequency {
			if value == personAmount {
				sum += 1
			}
		}
	}

	return sum
}
