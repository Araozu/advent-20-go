package solutions

import (
	"regexp"
	"strings"
)

func Day04Part01(isTest bool) int {
	input := ReadInput("04", isTest)
	passports := strings.Split(input, "\n\n")

	regex, _ := regexp.Compile("(\\w+:[\\w#]+)")

	correctAmount := 0

	for _, passport := range passports {
		matches := regex.FindAllString(passport, -1)

		values := make(map[string]string)

		for _, match := range matches {
			colonPosition := strings.Index(match, ":")
			key := match[:colonPosition]
			value := match[colonPosition+1:]
			values[key] = value
		}

		_, cidMissing := values["cid"]
		keysAmount := len(values)

		if keysAmount == 8 || (keysAmount == 7 && !cidMissing) {
			correctAmount += 1
		}
	}

	return correctAmount
}

func Day04Part02(isTest bool) int {
	return -1
}
