package solutions

import (
	"fmt"
	"os"
)

const DIR = "/home/fernando/GolandProjects/advent-20/"

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
