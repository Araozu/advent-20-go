package main

import (
	"advent-20/solutions"
	"fmt"
	"time"
)

func main() {
	runAndBenchmark("01", "1", false, solutions.Day01Part01)
	runAndBenchmark("01", "2", false, solutions.Day01Part02)

	runAndBenchmark("02", "1", false, solutions.Day02Part01)
	runAndBenchmark("02", "2", false, solutions.Day02Part02)

	runAndBenchmark("03", "1", false, solutions.Day03Part01)
	runAndBenchmark("03", "2", false, solutions.Day03Part02)

	runAndBenchmark("04", "1", false, solutions.Day04Part01)
	runAndBenchmark("04", "2", false, solutions.Day04Part02)

	runAndBenchmark("05", "1", false, solutions.Day05Part01)
	runAndBenchmark("05", "2", false, solutions.Day05Part02)

	runAndBenchmark("06", "1", false, solutions.Day06Part01)
	runAndBenchmark("06", "2", false, solutions.Day06Part02)
}

type execute func(bool) int

func runAndBenchmark(day, part string, isTest bool, fn execute) {
	startMs := time.Now().UnixMicro()
	computation := fn(isTest)
	endMs := time.Now().UnixMicro()
	duration := endMs - startMs
	fmt.Printf("Day %s part %s:\t%12d\t%5d micros.\n", day, part, computation, duration)
}
