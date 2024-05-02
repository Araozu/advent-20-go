package solutions

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const DIR = "./"

func ReadInput(day string, isTest bool) string {
	testStr := ""
	if isTest {
		testStr = "_test"
	}

	bytes, err := os.ReadFile(DIR + "inputs" + testStr + "/" + day + ".txt")
	if err != nil {
		fmt.Println(err)
		panic("Error reading file.")
	}

	return string(bytes)
}

func ReadAndMapInt(day string, isTest bool) []int {
	testStr := ""
	if isTest {
		testStr = "_test"
	}

	filePath := DIR + "inputs" + testStr + "/" + day + ".txt"

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	numbers := make([]int, 0)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		valueStr := scanner.Text()
		number, err := strconv.Atoi(valueStr)
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, number)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return numbers
}
