package solutions

import (
	"fmt"
	"os"
)

const DIR = "/home/fernando/GolandProjects/advent-20/"

func ReadInput(day string, isTest bool) string {
	testStr := ""
	if isTest {
		testStr = "test_"
	}

	bytes, err := os.ReadFile(DIR + testStr + "inputs/" + day + ".txt")
	if err != nil {
		fmt.Println(err)
		panic("Error reading file.")
	}

	return string(bytes)
}
