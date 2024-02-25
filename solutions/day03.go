package solutions

import "strings"

func countTrees(input string, left, down int) int {
	verticalPos := down
	horizontalPos := left
	lines := strings.Split(input, "\n")
	width := len(lines[0])
	treesCount := 0

	for verticalPos < len(lines) {
		line := lines[verticalPos]
		charAt := string(line[horizontalPos])

		if charAt == "#" {
			treesCount += 1
		}

		horizontalPos += left
		horizontalPos = horizontalPos % width
		verticalPos += down
	}

	return treesCount
}

func Day03Part01(isTest bool) int {
	input := ReadInput("03", isTest)

	return countTrees(input, 3, 1)
}

func Day03Part02(isTest bool) int {
	return -1
}
