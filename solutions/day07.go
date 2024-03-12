package solutions

import (
	"strconv"
	"strings"
)

type Bag struct {
	color    string
	contents []BagContent
	// -1 if not set, 0 if it doesn't contain, 1 if it does
	containsShinyGold int
	// -1 if not counted yet
	contentsCount int
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
		return Bag{color: bagColor, contents: make([]BagContent, 0), contentsCount: -1, containsShinyGold: -1}
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

	return Bag{color: bagColor, contents: bagContents, contentsCount: -1, containsShinyGold: -1}
}

func bagContainsShinyGold(bag *Bag, bagMap map[string]*Bag) bool {
	if bag.containsShinyGold == 1 {
		return true
	}
	if bag.containsShinyGold == 0 {
		return false
	}

	// recursively search
	for _, nextBag := range bag.contents {
		if nextBag.color == "shiny gold" {
			bag.containsShinyGold = 1
			return true
		}

		nextBagF := bagMap[nextBag.color]
		if bagContainsShinyGold(nextBagF, bagMap) {
			bag.containsShinyGold = 1
			return true
		}
	}

	bag.containsShinyGold = 0

	return false
}

func Day07Part01(isTest bool) int {
	input := ReadInput("07", isTest)
	groups := strings.Split(input, "\n")

	bags := make(map[string]*Bag)

	// parse and collect the bags
	for _, statement := range groups {
		parsedBag := parseBagStatement(statement)
		bags[parsedBag.color] = &parsedBag
	}

	shinyGoldContainers := 0

	// process the bags
	for _, bag := range bags {
		if bagContainsShinyGold(bag, bags) {
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

	bag.contentsCount = innerBagsCount

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
