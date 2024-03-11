package solutions

import (
	"strconv"
	"strings"
)

type Bag struct {
	color             string
	contents          []BagContent
	containsShinyGold bool
}

type BagContent struct {
	count int
	color string
}

// Given a bag statement string, returns its contents
func parseBagStatement(bagStatement string) Bag {
	bagEndIndex := strings.Index(bagStatement, " bags contain ")
	bagColor := bagStatement[:bagEndIndex]
	bagContentsStr := bagStatement[bagEndIndex+14:]

	// if no aditional bags
	if bagContentsStr[:2] == "no" {
		return Bag{color: bagColor, contents: make([]BagContent, 0)}
	}

	// parse remainder bags
	bagsContained := strings.Split(bagContentsStr, ", ")
	bagContents := make([]BagContent, len(bagsContained))

	for idx, bag := range bagsContained {
		bagCount, err := strconv.Atoi(bag[:1])
		if err != nil {
			panic("expected a number, got `" + bag + "`")
		}

		bagStrIndex := strings.Index(bag, " bag")
		bagColor2 := bag[2:bagStrIndex]

		bagContents[idx] = BagContent{count: bagCount, color: bagColor2}
	}

	/*
		grammar for bag statements:

		statement = bag color, "bags contain", (bag list | empty bag), "."
		empty bag = "no other bags"
		bag list = bag declaration, bag declaration+
		bag declaration = number, bag color, "bag", "s"?
		bag color = word, word
	*/

	return Bag{color: bagColor, contents: bagContents}
}

func bagContainsShinyGold(bag *Bag, bagMap map[string]Bag) bool {
	if bag.containsShinyGold {
		return true
	}

	// recursively search
	for _, bagName := range bag.contents {
		if bagName.color == "shiny gold" {
			bag.containsShinyGold = true
			return true
		}

		nextBag := bagMap[bagName.color]

		if bagContainsShinyGold(&nextBag, bagMap) {
			bag.containsShinyGold = true
			return true
		}
	}

	return false
}

func Day07Part01(isTest bool) int {
	input := ReadInput("07", isTest)
	groups := strings.Split(input, "\n")

	bags := make(map[string]Bag)

	// parse and collect the bags
	for _, statement := range groups {
		parsedBag := parseBagStatement(statement)
		bags[parsedBag.color] = parsedBag
	}

	shinyGoldContainers := 0

	// process the bags
	for _, bag := range bags {
		if bagContainsShinyGold(&bag, bags) {
			shinyGoldContainers += 1
		}
	}

	return shinyGoldContainers
}

func Day07Part02(isTest bool) int {
	return -1
}
