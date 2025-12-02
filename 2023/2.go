//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Part 1:")
	part1()
	fmt.Println("Part 2:")
	part2()
}

func isNumber(c string) bool {
	re := regexp.MustCompile(`^[0-9]+$`)
	return re.MatchString(c)
}

func part1() {
	filePath := "../../AoC-input/2023/2/input.txt"
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

	sum := 0
	for i, line := range fileLines {
		var isGood bool = true
		words := strings.Fields(line)
		words = words[2:]

		for j, word := range words {
			if isNumber(word) {
				wordInt, _ := strconv.Atoi(word)
				if strings.Contains(words[j+1], "red") {
					if wordInt > 12 {
						isGood = false
						break
					}
				} else if strings.Contains(words[j+1], "green") {
					if wordInt > 13 {
						isGood = false
						break
					}
				} else if strings.Contains(words[j+1], "blue") {
					if wordInt > 14 {
						isGood = false
						break
					}
				}
			}
		}

		if isGood {
			sum += i + 1
		}
	}

	fmt.Println(sum)
}

func part2() {
	filePath := "../../AoC-input/2023/2/input.txt"
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

	sum := 0
	for _, line := range fileLines {
		maxRed := 0
		maxGreen := 0
		maxBlue := 0
		words := strings.Fields(line)
		words = words[2:]

		for j, word := range words {
			if isNumber(word) {
				wordInt, _ := strconv.Atoi(word)
				if strings.Contains(words[j+1], "red") {
					if wordInt > maxRed {
						maxRed = wordInt
					}
				} else if strings.Contains(words[j+1], "green") {
					if wordInt > maxGreen {
						maxGreen = wordInt
					}
				} else if strings.Contains(words[j+1], "blue") {
					if wordInt > maxBlue {
						maxBlue = wordInt
					}
				}
			}
		}

		sum += maxRed * maxGreen * maxBlue
	}

	fmt.Println(sum)
}
