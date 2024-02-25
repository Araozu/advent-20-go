package solutions

import (
	"strconv"
	"strings"
)

func checkPassword(input string) bool {
	dashPosition := strings.Index(input, "-")
	spacePosition := strings.Index(input, " ")
	colonPosition := strings.Index(input, ":")

	amountMin, err := strconv.Atoi(input[:dashPosition])
	if err != nil {
		panic("Error converting to int")
	}
	amountMax, err := strconv.Atoi(input[dashPosition+1 : spacePosition])
	if err != nil {
		panic("Error converting to int")
	}

	letter := input[spacePosition+1 : colonPosition]
	rest := input[colonPosition+2:]

	charMap := make(map[string]int)

	for _, char := range strings.Split(rest, "") {
		if _, ok := charMap[char]; ok {
			charMap[char] += 1
		} else {
			charMap[char] = 1
		}
	}

	charAmount := charMap[letter]

	if charAmount >= amountMin && charAmount <= amountMax {
		return true
	}

	return false
}

func Day02Part01(isTest bool) int {
	input := ReadInput("02", isTest)
	values := strings.Split(input, "\n")
	correctPasswords := 0

	for _, i := range values {
		if checkPassword(i) {
			correctPasswords += 1
		}
	}

	return correctPasswords
}

func Day02Part02(isTest bool) int {
	return -1
}
