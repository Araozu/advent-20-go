package main

import (
	"advent-20/solutions"
	"fmt"
	"time"
)

func main() {
	runAndBenchmark("01", "1", false, solutions.Day01Part01)
	runAndBenchmark("01", "2", false, solutions.Day01Part02)

	runAndBenchmark("02", "1", false, solutions.Day01Part01)
	runAndBenchmark("02", "2", false, solutions.Day01Part02)
}

type execute func(bool) int

func runAndBenchmark(day, part string, isTest bool, fn execute) {
	startMs := time.Now().UnixMicro()
	computation := fn(isTest)
	endMs := time.Now().UnixMicro()
	duration := endMs - startMs
	fmt.Printf("Day %s part %s:\t%10d\t%5d micros.\n", day, part, computation, duration)
}
