/*
Description:

--- Day 7: Camel Cards ---
Your all-expenses-paid trip turns out to be a one-way, five-minute ride in an airship. (At least it's a cool airship!) It drops you off at the edge of a vast desert and descends back to Island Island.

"Did you bring the parts?"

You turn around to see an Elf completely covered in white clothing, wearing goggles, and riding a large camel.

"Did you bring the parts?" she asks again, louder this time. You aren't sure what parts she's looking for; you're here to figure out why the sand stopped.

"The parts! For the sand, yes! Come with me; I will show you." She beckons you onto the camel.

After riding a bit across the sands of Desert Island, you can see what look like very large rocks covering half of the horizon. The Elf explains that the rocks are all along the part of Desert Island that is directly above Island Island, making it hard to even get there. Normally, they use big machines to move the rocks and filter the sand, but the machines have broken down because Desert Island recently stopped receiving the parts they need to fix the machines.

You've already assumed it'll be your job to figure out why the parts stopped when she asks if you can help. You agree automatically.

Because the journey will take a few days, she offers to teach you the game of Camel Cards. Camel Cards is sort of similar to poker except it's designed to be easier to play while riding a camel.

In Camel Cards, you get a list of hands, and your goal is to order them based on the strength of each hand. A hand consists of five cards labeled one of A, K, Q, J, T, 9, 8, 7, 6, 5, 4, 3, or 2. The relative strength of each card follows this order, where A is the highest and 2 is the lowest.

Every hand is exactly one type. From strongest to weakest, they are:

Five of a kind, where all five cards have the same label: AAAAA
Four of a kind, where four cards have the same label and one card has a different label: AA8AA
Full house, where three cards have the same label, and the remaining two cards share a different label: 23332
Three of a kind, where three cards have the same label, and the remaining two cards are each different from any other card in the hand: TTT98
Two pair, where two cards share one label, two other cards share a second label, and the remaining card has a third label: 23432
One pair, where two cards share one label, and the other three cards have a different label from the pair and each other: A23A4
High card, where all cards' labels are distinct: 23456
Hands are primarily ordered based on type; for example, every full house is stronger than any three of a kind.

If two hands have the same type, a second ordering rule takes effect. Start by comparing the first card in each hand. If these cards are different, the hand with the stronger first card is considered stronger. If the first card in each hand have the same label, however, then move on to considering the second card in each hand. If they differ, the hand with the higher second card wins; otherwise, continue with the third card in each hand, then the fourth, then the fifth.

So, 33332 and 2AAAA are both four of a kind hands, but 33332 is stronger because its first card is stronger. Similarly, 77888 and 77788 are both a full house, but 77888 is stronger because its third card is stronger (and both hands have the same first and second card).

To play Camel Cards, you are given a list of hands and their corresponding bid (your puzzle input). For example:

32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483
This example shows five hands; each hand is followed by its bid amount. Each hand wins an amount equal to its bid multiplied by its rank, where the weakest hand gets rank 1, the second-weakest hand gets rank 2, and so on up to the strongest hand. Because there are five hands in this example, the strongest hand will have rank 5 and its bid will be multiplied by 5.

So, the first step is to put the hands in order of strength:

32T3K is the only one pair and the other hands are all a stronger type, so it gets rank 1.
KK677 and KTJJT are both two pair. Their first cards both have the same label, but the second card of KK677 is stronger (K vs T), so KTJJT gets rank 2 and KK677 gets rank 3.
T55J5 and QQQJA are both three of a kind. QQQJA has a stronger first card, so it gets rank 5 and T55J5 gets rank 4.
Now, you can determine the total winnings of this set of hands by adding up the result of multiplying each hand's bid with its rank (765 * 1 + 220 * 2 + 28 * 3 + 684 * 4 + 483 * 5). So the total winnings in this example are 6440.

Find the rank of every hand in your set. What are the total winnings?
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	fmt.Println("Part 1:")
	part1()
	fmt.Println("Part 2:")
	part2()
}

func part1() {
	filePath := "./7/input.txt"
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var fileLines []string

	for scanner.Scan() {
		fileLines = append(fileLines, scanner.Text())
	}

	array := make([][]string, 7)
	for i := range array {
		array[i] = make([]string, 0)
	}

	for _, line := range fileLines {
		hand := line[:5]
		if isXOfAKind(hand, 5) {
			array[0] = append(array[0], line)
		} else if isXOfAKind(hand, 4) {
			array[1] = append(array[1], line)
		} else if isFullHouse(hand) {
			array[2] = append(array[2], line)
		} else if isXOfAKind(hand, 3) {
			array[3] = append(array[3], line)
		} else if isTwoPair(hand) {
			array[4] = append(array[4], line)
		} else if isXOfAKind(hand, 2) {
			array[5] = append(array[5], line)
		} else {
			array[6] = append(array[6], line)
		}
	}

	for i := range array {
		array[i] = sortHands(array[i])
	}

	sum := 0
	multiplayer := 1
	for i := len(array) - 1; i >= 0; i-- {
		for _, hand := range array[i] {
			handInt, _ := strconv.Atoi(hand[6:])
			sum += multiplayer * handInt
			multiplayer++
		}
	}

	fmt.Println(sum)
}

func isXOfAKind(hand string, x int) bool {
	cards := map[string]int{}
	for _, card := range hand {
		cards[string(card)]++
	}
	for _, count := range cards {
		if count == x {
			return true
		}
	}
	return false
}

func isFullHouse(hand string) bool {
	cards := map[string]int{}
	for _, card := range hand {
		cards[string(card)]++
	}
	if len(cards) == 2 {
		for _, count := range cards {
			if count == 3 || count == 2 {
				return true
			}
		}
	}
	return false
}

func isTwoPair(hand string) bool {
	cards := map[string]int{}
	for _, card := range hand {
		cards[string(card)]++
	}
	if len(cards) == 3 {
		for _, count := range cards {
			if count == 2 {
				return true
			}
		}
	}
	return false
}

func sortHands(hands []string) []string {
	sort.Slice(hands, func(i, j int) bool {
		index := 0
		for hands[i][index] == hands[j][index] {
			index++
		}
		return points(string(hands[i][index])) < points(string(hands[j][index]))
	})
	return hands
}

func points(card string) int {
	points := 0
	switch card {
	case "A":
		points = 14
	case "K":
		points = 13
	case "Q":
		points = 12
	case "J":
		points = 11
	case "T":
		points = 10
	default:
		cardInt, _ := strconv.Atoi(card)
		points = cardInt
	}

	return points
}

/*
--- Part Two ---
To make things a little more interesting, the Elf introduces one additional rule. Now, J cards are jokers - wildcards that can act like whatever card would make the hand the strongest type possible.

To balance this, J cards are now the weakest individual cards, weaker even than 2. The other cards stay in the same order: A, K, Q, T, 9, 8, 7, 6, 5, 4, 3, 2, J.

J cards can pretend to be whatever card is best for the purpose of determining hand type; for example, QJJQ2 is now considered four of a kind. However, for the purpose of breaking ties between two hands of the same type, J is always treated as J, not the card it's pretending to be: JKKK2 is weaker than QQQQ2 because J is weaker than Q.

Now, the above example goes very differently:

32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483
32T3K is still the only one pair; it doesn't contain any jokers, so its strength doesn't increase.
KK677 is now the only two pair, making it the second-weakest hand.
T55J5, KTJJT, and QQQJA are now all four of a kind! T55J5 gets rank 3, QQQJA gets rank 4, and KTJJT gets rank 5.
With the new joker rule, the total winnings in this example are 5905.

Using the new joker rule, find the rank of every hand in your set. What are the new total winnings?
*/

func part2() {
	filePath := "./7/input.txt"
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var fileLines []string

	for scanner.Scan() {
		fileLines = append(fileLines, scanner.Text())
	}

	array := make([][]string, 7)
	for i := range array {
		array[i] = make([]string, 0)
	}

	for _, line := range fileLines {
		hand := line[:5]
		if p2_isXOfAKind(hand, 5) {
			array[0] = append(array[0], line)
		} else if p2_isXOfAKind(hand, 4) {
			array[1] = append(array[1], line)
		} else if p2_isFullHouse(hand) {
			array[2] = append(array[2], line)
		} else if p2_isXOfAKind(hand, 3) {
			array[3] = append(array[3], line)
		} else if p2_isTwoPair(hand) {
			array[4] = append(array[4], line)
		} else if p2_isXOfAKind(hand, 2) {
			array[5] = append(array[5], line)
		} else {
			array[6] = append(array[6], line)
		}
	}

	for i := range array {
		array[i] = p2_sortHands(array[i])
	}

	sum := 0
	multiplayer := 1
	for i := len(array) - 1; i >= 0; i-- {
		for _, hand := range array[i] {
			handInt, _ := strconv.Atoi(hand[6:])
			sum += multiplayer * handInt
			multiplayer++
		}
	}

	fmt.Println(sum)
}

func p2_isXOfAKind(hand string, x int) bool {
	cards := map[string]int{}
	jokers := 0
	for _, card := range hand {
		if string(card) == "J" {
			jokers++
		} else {
			cards[string(card)]++
		}
	}
	if jokers == x {
		return true
	}
	for _, count := range cards {
		if count+jokers == x {
			return true
		}
	}
	return false
}

func p2_isFullHouse(hand string) bool {
	cards := map[string]int{}
	jokers := 0
	for _, card := range hand {
		if string(card) == "J" {
			jokers++
		} else {
			cards[string(card)]++
		}
	}
	if len(cards) == 2 {
		for _, count := range cards {
			if count+jokers == 3 || count+jokers == 2 {
				return true
			}
		}
	}
	return false
}

func p2_isTwoPair(hand string) bool {
	cards := map[string]int{}
	jokers := 0
	for _, card := range hand {
		if string(card) == "J" {
			jokers++
		} else {
			cards[string(card)]++
		}
	}

	if jokers == 0 {
		if len(cards) == 3 {
			for _, count := range cards {
				if count == 2 {
					return true
				}
			}
		}
	}
	return false
}

func p2_sortHands(hands []string) []string {
	sort.Slice(hands, func(i, j int) bool {
		index := 0
		for hands[i][index] == hands[j][index] {
			index++
		}
		return p2_points(string(hands[i][index])) < p2_points(string(hands[j][index]))
	})
	return hands
}

func p2_points(card string) int {
	points := 0
	switch card {
	case "A":
		points = 14
	case "K":
		points = 13
	case "Q":
		points = 12
	case "J":
		points = 0
	case "T":
		points = 10
	default:
		cardInt, _ := strconv.Atoi(card)
		points = cardInt
	}

	return points
}
