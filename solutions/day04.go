package solutions

import (
	"regexp"
	"strconv"
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

		_, cidPresent := values["cid"]
		keysAmount := len(values)

		if keysAmount == 8 || (keysAmount == 7 && !cidPresent) {
			correctAmount += 1
		}
	}

	return correctAmount
}

func Day04Part02(isTest bool) int {
	input := ReadInput("04", isTest)
	passports := strings.Split(input, "\n\n")

	regex, _ := regexp.Compile("[\\n\\s]")

	correctAmount := 0

	for _, passport := range passports {
		matches := regex.Split(passport, -1)
		// s := "hgt:59cm ecl:zzz\neyr:2038 hcl:74454a iyr:2023\npid:3556412378 byr:2007"
		// matches = regex.FindAllString(s, -1)

		values := make(map[string]string)

		for _, match := range matches {
			colonPosition := strings.Index(match, ":")
			key := match[:colonPosition]
			value := match[colonPosition+1:]
			values[key] = value
		}

		if !validateByr(values) || !validateIyr(values) || !validateEyr(values) ||
			!validateHgt(values) || !validateHcl(values) || !validateEcl(values) ||
			!validatePid(values) {
			continue
		}

		correctAmount += 1
	}

	return correctAmount
}

func validateYear(year string, min, max int) bool {
	yearNumber, err := strconv.Atoi(year)
	if err != nil {
		return false
	}

	return min <= yearNumber && yearNumber <= max
}

func validateByr(values map[string]string) bool {
	byr, _ := values["byr"]
	return validateYear(byr, 1920, 2002)
}

func validateIyr(values map[string]string) bool {
	byr, _ := values["iyr"]
	return validateYear(byr, 2010, 2020)
}

func validateEyr(values map[string]string) bool {
	byr, _ := values["eyr"]
	return validateYear(byr, 2020, 2030)
}

func validateHgt(values map[string]string) bool {
	height, ok := values["hgt"]
	if !ok {
		return false
	}

	cmIdx := strings.Index(height, "cm")
	if cmIdx != -1 {
		cm, err := strconv.Atoi(height[:cmIdx])
		if err != nil {
			return false
		}

		return 150 <= cm && cm <= 193
	}
	inIdx := strings.Index(height, "in")
	if inIdx != -1 {
		in, err := strconv.Atoi(height[:inIdx])
		if err != nil {
			return false
		}

		return 59 <= in && in <= 76
	}

	return false
}

func validateHcl(values map[string]string) bool {
	hairColor, ok := values["hcl"]
	if !ok {
		return false
	}

	matched, _ := regexp.MatchString("#[a-f0-9]{6}", hairColor)
	return matched
}

func validateEcl(values map[string]string) bool {
	eyeColor, ok := values["ecl"]
	if !ok {
		return false
	}

	return eyeColor == "amb" || eyeColor == "blu" || eyeColor == "brn" ||
		eyeColor == "gry" || eyeColor == "grn" || eyeColor == "hzl" || eyeColor == "oth"
}

func validatePid(values map[string]string) bool {
	passportId, ok := values["pid"]
	if !ok {
		return false
	}

	matched, _ := regexp.MatchString("\\d{9}", passportId)
	return matched
}
