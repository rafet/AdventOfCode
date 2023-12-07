package main

import (
	"os"
	"sort"
	"strconv"
	"strings"
)

type info struct {
	hand string
	bid  int
}

var valMap = map[rune]int{
	'A': 14,
	'K': 13,
	'Q': 12,
	'J': 11,
	'T': 10,
	'9': 9,
	'8': 8,
	'7': 7,
	'6': 6,
	'5': 5,
	'4': 4,
	'3': 3,
	'2': 2,
}

func main() {
	infos := getInfos()

	// Part 1
	sum := calculate(infos, false)
	println("answer 1:", sum)

	// Part 2
	valMap['J'] = 1
	sum2 := calculate(infos, true)
	println("answer 2:", sum2)

}
func calculate(infos []info, wildcard bool) int {
	sum := 0

	sort.Slice(infos, func(i, j int) bool {
		return compareHands(infos[i].hand, infos[j].hand, wildcard) == 0
	})

	for index, i := range infos {
		sum += i.bid * (len(infos) - index)
	}

	return sum
}

func isFiveOfAKind(hand string) bool {
	var chars = []rune(hand)
	for i := 0; i < len(chars); i++ {
		if chars[i] != chars[0] {
			return false
		}
	}

	return true
}

func isFourOfAKind(hand string) bool {
	var charCounts = getCardCounts(hand)

	for _, count := range charCounts {
		if count == 4 {
			return true
		}
	}

	return false
}

func isFullHouse(hand string) bool {
	var charCounts = getCardCounts(hand)

	var hasThreeOfAKind bool
	var hasTwoOfAKind bool
	for _, count := range charCounts {
		if count == 3 {
			hasThreeOfAKind = true
		}

		if count == 2 {
			hasTwoOfAKind = true
		}
	}

	return hasThreeOfAKind && hasTwoOfAKind
}

func isThreeOfAKind(hand string) bool {
	var charCounts = getCardCounts(hand)

	for _, count := range charCounts {
		if count == 3 {
			return true
		}
	}

	return false
}

func isTwoPairs(hand string) bool {
	var charCounts = getCardCounts(hand)

	var pairCount int
	for _, count := range charCounts {
		if count == 2 {
			pairCount++
		}
	}

	return pairCount == 2
}

func isOnePair(hand string) bool {
	var charCounts = getCardCounts(hand)

	for _, count := range charCounts {
		if count == 2 {
			return true
		}
	}

	return false
}

func isHighCard(hand string) bool {
	var charCounts = getCardCounts(hand)

	return len(charCounts) == 5
}

func compareHands(hand1, hand2 string, wildcard bool) int {
	var hand1Rank, hand2Rank int

	if strings.Contains(hand1, "J") && wildcard {
		hand1Rank = getRankWithWildCard(hand1)
	} else {
		hand1Rank = getHandRank(hand1)
	}

	if strings.Contains(hand2, "J") && wildcard {
		hand2Rank = getRankWithWildCard(hand2)
	} else {
		hand2Rank = getHandRank(hand2)
	}

	if hand1Rank > hand2Rank {
		return 0
	} else if hand1Rank < hand2Rank {
		return 1
	} else {
		return compareSameRankHands(hand1, hand2)
	}
}

func compareSameRankHands(hand1, hand2 string) int {
	for i := 0; i < len(hand1); i++ {
		if valMap[rune(hand1[i])] > valMap[rune(hand2[i])] {
			return 0
		}

		if valMap[rune(hand1[i])] < valMap[rune(hand2[i])] {
			return 1
		}
	}

	return -1
}

func getHandRank(hand string) int {
	if isFiveOfAKind(hand) {
		return 7
	}

	if isFourOfAKind(hand) {
		return 6
	}

	if isFullHouse(hand) {
		return 5
	}

	if isThreeOfAKind(hand) {
		return 4
	}

	if isTwoPairs(hand) {
		return 3
	}

	if isOnePair(hand) {
		return 2
	}

	if isHighCard(hand) {
		return 1
	}

	return 0
}

func getCardCounts(hand string) map[rune]int {
	var charCounts = make(map[rune]int)
	var chars = []rune(hand)
	for i := 0; i < len(chars); i++ {
		charCounts[chars[i]]++
	}

	return charCounts
}
func jokerCount(hand string) int {
	var count int
	var chars = []rune(hand)
	for i := 0; i < len(chars); i++ {
		if chars[i] == 'J' {
			count++
		}
	}

	return count
}
func getRankWithWildCard(hand string) int {
	if isFiveOfAKind(hand) { // Example: 2 2 2 2 2
		return 7
	}

	if isFourOfAKind(hand) { // Example: 2 2 2 2 3
		return 7
	}

	if isFullHouse(hand) { // Example: 2 2 2 3 3
		return 7
	}

	if isThreeOfAKind(hand) { // Example: 2 2 2 3 4
		return 6
	}

	if isTwoPairs(hand) { // Example: 2 2 3 3 4
		if jokerCount(hand) == 2 {
			return 6 // Example: J J 2 2 3 -> 2 2 2 2 3 (5 of a kind)
		} else {
			return 5 // Example: 2 2 3 3 J -> 2 2 3 3 3 (full house)
		}
	}

	if isOnePair(hand) { // Example: 2 2 3 4 5
		return 4
	}

	if isHighCard(hand) { // Example: 2 3 4 5 6
		return 2
	}

	return 0
}

func getInfos() []info {
	file, err := os.ReadFile("2023/day7/input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(file), "\n")
	lines = lines[:len(lines)-1]

	infos := make([]info, len(lines))
	for i, line := range lines {
		fields := strings.Fields(line)
		bid, _ := strconv.Atoi(fields[1])

		infos[i] = info{
			hand: fields[0],
			bid:  bid,
		}
	}

	return infos
}
