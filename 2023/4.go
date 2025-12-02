//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Part 1:")
	part1()
	fmt.Println("Part 2:")
	part2()
}

func part1() {
	filePath := "../../AoC-input/2023/4/input.txt"
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

	pointsSum := 0

	for _, line := range fileLines {
		winningNumbers := make(map[int]bool)

		words := strings.Split(line, " ")
		words = words[2:]

		points := 0

		winningNumbersMode := true
		for _, word := range words {
			if winningNumbersMode {
				if word == "|" {
					winningNumbersMode = false
					continue
				}

				winningNumber, _ := strconv.Atoi(word)
				winningNumbers[winningNumber] = true
			} else {
				number, _ := strconv.Atoi(word)
				if winningNumbers[number] && number != 0 {
					if points == 0 {
						points = 1
					} else {
						points = points * 2
					}
				}
			}
		}

		pointsSum += points
	}

	fmt.Println(pointsSum)
}

func part2() {
	filePath := "../../AoC-input/2023/4/input.txt"
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

	cardPoints := make(map[int]int)
	numberOfCards := make(map[int]int)

	for i, line := range fileLines {
		winningNumbers := make(map[int]bool)

		words := strings.Split(line, " ")
		words = words[2:]

		points := 0

		winningNumbersMode := true
		for _, word := range words {
			if winningNumbersMode {
				if word == "|" {
					winningNumbersMode = false
					continue
				}

				winningNumber, _ := strconv.Atoi(word)
				winningNumbers[winningNumber] = true
			} else {
				number, _ := strconv.Atoi(word)
				if winningNumbers[number] && number != 0 {
					points++
				}
			}
		}

		cardPoints[i] = points
		numberOfCards[i] = 1
	}

	pointsSum := 0
	for i := 0; i < len(cardPoints); i++ {
		pointsSum += numberOfCards[i]
		for j := i + 1; j < i+1+cardPoints[i]; j++ {
			numberOfCards[j] += numberOfCards[i]
		}
	}

	fmt.Println(pointsSum)
}
