package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Hand struct {
	card     string
	bid      int
	handType int
}

func readFile(fname string) []string {
	var lines []string
	file, err := os.Open(fname)
	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	return lines
}

func cardValue(card string, jokers bool) int {
	CARDS := map[string]int{
		"2": 2,
		"3": 3,
		"4": 4,
		"5": 5,
		"6": 6,
		"7": 7,
		"8": 8,
		"9": 9,
		"T": 10,
		"J": 11,
		"Q": 12,
		"K": 13,
		"A": 14,
	}
	if jokers && card == "J" {
		return 1
	}
	return CARDS[card]
}

func compareHands(hand1 string, hand2 string, jokers bool) bool {

	for i := 0; i < len(hand1); i++ {
		if hand1[i] != hand2[i] {
			return cardValue(string(hand1[i]), jokers) < cardValue(string(hand2[i]), jokers)
		}
	}

	return false
}

func handValue(hand string, jokers bool) int {
	cards := map[byte]int{}
	highestCardCount := 0
	for _, card := range hand {
		cards[byte(card)]++
		highestCardCount = max(highestCardCount, cards[byte(card)])
	}
	jokersCount := cards['J']

	switch len(cards) {
	case 1:
		return 7
	case 2:
		if jokers && jokersCount > 0 {
			return 7
		}
		if highestCardCount == 4 {
			return 6
		} else {
			return 5

		}
	case 3:
		if highestCardCount == 3 {
			if jokers && jokersCount > 0 {
				return 6
			}
			return 4
		}
		if jokers && jokersCount > 1 {
			return 6
		} else if jokers && jokersCount > 0 {
			return 5
		}
		return 2

	case 4:
		if jokers && jokersCount > 0 {
			return 4
		}
		return 1
	case 5:
		if jokers && jokersCount > 0 {
			return 1
		}
		return 0
	}
	return 0
}

func main() {
	var hands []Hand
	INPUT := "input_7.txt"
	fileContent := readFile(INPUT)
	for _, line := range fileContent {
		var hand Hand
		lineContent := strings.Split(line, " ")
		hand.card = lineContent[0]
		hand.bid, _ = strconv.Atoi(lineContent[1])
		hand.handType = handValue(lineContent[0], false)
		hands = append(hands, hand)
	}
	sort.Slice(hands, func(i, j int) bool {
		if hands[i].handType == hands[j].handType {
			return compareHands(hands[i].card, hands[j].card, false)
		}
		return hands[i].handType < hands[j].handType
	})
	sum_part1 := 0
	for rank, hand := range hands {
		sum_part1 = sum_part1 + hand.bid*(rank+1)
	}
	fmt.Println("Sum part1:", sum_part1)

	for counter, _ := range hands {
		hands[counter].handType = handValue(hands[counter].card, true)

	}
	sort.Slice(hands, func(i, j int) bool {
		if hands[i].handType == hands[j].handType {
			return compareHands(hands[i].card, hands[j].card, true)
		}
		return hands[i].handType < hands[j].handType
	})
	sum_part2 := 0
	for rank, hand := range hands {
		sum_part2 = sum_part2 + hand.bid*(rank+1)
	}
	fmt.Println("Sum part2:", sum_part2)
}
