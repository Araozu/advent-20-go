package solutions

import (
	"regexp"
	"strconv"
	"strings"
)

func seatToValue(seat string) (int, int, int) {
	lowRegex := regexp.MustCompile("[FL]")
	highRegex := regexp.MustCompile("[BR]")

	seatStr := lowRegex.ReplaceAllString(seat, "0")
	seatStr = highRegex.ReplaceAllString(seatStr, "1")

	row, _ := strconv.ParseInt(seatStr[:7], 2, 64)
	column, _ := strconv.ParseInt(seatStr[7:], 2, 64)

	return int(row), int(column), int(row*8 + column)
}

func Day05Part01(isTest bool) int {
	input := ReadInput("05", isTest)
	seats := strings.Split(input, "\n")

	highestSeatId := 0

	for _, seat := range seats {
		_, _, seatId := seatToValue(seat)
		if seatId > highestSeatId {
			highestSeatId = seatId
		}
	}

	return highestSeatId
}

func Day05Part02(isTest bool) int {
	return -1
}
