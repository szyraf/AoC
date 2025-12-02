//go:build ignore

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
	filePath := "../../AoC-input/2023/7/input.txt"
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

func part2() {
	filePath := "../../AoC-input/2023/7/input.txt"
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
