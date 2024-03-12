package solutions

import (
	"strconv"
	"strings"
)

type Bag struct {
	color             string
	contents          []BagContent
	containsShinyGold bool
	contentsCount     int
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
		return Bag{color: bagColor, contents: make([]BagContent, 0), contentsCount: -1}
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

	return Bag{color: bagColor, contents: bagContents, contentsCount: -1}
}

func bagContainsShinyGold(bagColor string, bagMap map[string]*Bag) bool {
	bag := bagMap[bagColor]
	if bag.containsShinyGold {
		return true
	}

	// recursively search
	for _, nextBag := range bag.contents {
		if nextBag.color == "shiny gold" {
			bag.containsShinyGold = true
			return true
		}

		if bagContainsShinyGold(nextBag.color, bagMap) {
			bag.containsShinyGold = true
			return true
		}
	}

	return false
}

func Day07Part01(isTest bool) int {
	input := ReadInput("07", isTest)
	groups := strings.Split(input, "\n")

	bags := make(map[string]*Bag)
	bagColors := make([]string, len(groups))

	// parse and collect the bags
	for i, statement := range groups {
		parsedBag := parseBagStatement(statement)
		bags[parsedBag.color] = &parsedBag
		bagColors[i] = parsedBag.color
	}

	shinyGoldContainers := 0

	// process the bags
	for _, bagColor := range bagColors {
		if bagContainsShinyGold(bagColor, bags) {
			shinyGoldContainers += 1
		}
	}

	return shinyGoldContainers
}

func countInnerBags(bag *Bag, bagMap map[string]*Bag) int {
	if bag.contentsCount != -1 {
		return bag.contentsCount
	}

	innerBagsCount := 0

	// recursively count bags
	for _, nextBagStruct := range bag.contents {
		nextBagCount := nextBagStruct.count
		nextBag := bagMap[nextBagStruct.color]

		innerBagsCount += nextBagCount + nextBagCount*countInnerBags(nextBag, bagMap)
	}

	return innerBagsCount
}

func Day07Part02(isTest bool) int {
	input := ReadInput("07", isTest)
	groups := strings.Split(input, "\n")

	bags := make(map[string]*Bag)

	// parse and collect the bags
	for _, statement := range groups {
		parsedBag := parseBagStatement(statement)
		bags[parsedBag.color] = &parsedBag
	}

	shinyGoldBag := bags["shiny gold"]

	return countInnerBags(shinyGoldBag, bags)
}
