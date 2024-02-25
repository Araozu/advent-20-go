package solutions

import "strings"

func countTrees(input string, right, down int) int {
	verticalPos := down
	horizontalPos := right
	lines := strings.Split(input, "\n")
	width := len(lines[0])
	treesCount := 0

	for verticalPos < len(lines) {
		line := lines[verticalPos]
		charAt := string(line[horizontalPos])

		if charAt == "#" {
			treesCount += 1
		}

		horizontalPos += right
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
	input := ReadInput("03", isTest)

	amount1 := countTrees(input, 1, 1)
	amount2 := countTrees(input, 3, 1)
	amount3 := countTrees(input, 5, 1)
	amount4 := countTrees(input, 7, 1)
	amount5 := countTrees(input, 1, 2)

	return amount1 * amount2 * amount3 * amount4 * amount5
}
