package solutions

import (
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func seatStrToSeatId(seat string) int {
	lowRegex := regexp.MustCompile("[FL]")
	highRegex := regexp.MustCompile("[BR]")

	seatStr := lowRegex.ReplaceAllString(seat, "0")
	seatStr = highRegex.ReplaceAllString(seatStr, "1")

	row, _ := strconv.ParseInt(seatStr[:7], 2, 64)
	column, _ := strconv.ParseInt(seatStr[7:], 2, 64)

	return int(row*8 + column)
}

func Day05Part01(isTest bool) int {
	input := ReadInput("05", isTest)
	seats := strings.Split(input, "\n")

	highestSeatId := 0

	for _, seat := range seats {
		seatId := seatStrToSeatId(seat)
		if seatId > highestSeatId {
			highestSeatId = seatId
		}
	}

	return highestSeatId
}

func Day05Part02(isTest bool) int {
	input := ReadInput("05", isTest)
	seats := strings.Split(input, "\n")
	seatsAmount := len(seats)

	seatIds := make([]int, seatsAmount)

	for i, seat := range seats {
		seatId := seatStrToSeatId(seat)
		seatIds[i] = seatId
	}

	slices.Sort(seatIds)

	for i := 0; i < seatsAmount-1; i++ {
		el1 := seatIds[i]
		el2 := seatIds[i+1]

		if el2-el1 == 2 {
			return el1 + 1
		}
	}

	return -1
}
