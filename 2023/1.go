//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

func main() {
	fmt.Println("Part 1:")
	part1()
	fmt.Println("Part 2:")
	part2()
}

func isNumber(c string) bool {
	re := regexp.MustCompile(`^\d$`)
	return re.MatchString(c)
}

func part1() {
	filePath := "../../AoC-input/2023/1/input.txt"
	file, err := os.Open(filePath)

	if err != nil {
		log.Fatal(err)
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var fileLines []string

	for scanner.Scan() {
		fileLines = append(fileLines, scanner.Text())
	}

	sum := 0
	for _, line := range fileLines {
		start := -1
		end := -1
		for _, char := range line {
			if isNumber(string(char)) {
				if start == -1 {
					start = int(char) - 48
				}
				end = int(char) - 48
			}
		}
		sum += start*10 + end
	}

	fmt.Println(sum)
}

func part2() {
	filePath := "../../AoC-input/2023/1/input.txt"
	file, err := os.Open(filePath)

	if err != nil {
		log.Fatal(err)
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var fileLines []string

	for scanner.Scan() {
		fileLines = append(fileLines, scanner.Text())
	}

	numbers := [...]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	sum := 0
	for _, line := range fileLines {
		start := -1
		end := -1
		for i, char := range line {
			if isNumber(string(char)) {
				if start == -1 {
					start = int(char) - 48
				}
				end = int(char) - 48
			} else {
				for j, number := range numbers {
					isInNumbers := true
					for k, letter := range number {
						if i+k < len(line) {
							if string(line[i+k]) != string(letter) {
								isInNumbers = false
								break
							}
						} else {
							isInNumbers = false
							break
						}
					}

					if isInNumbers {
						if start == -1 {
							start = j + 1
						}
						end = j + 1
					}
				}
			}
		}

		sum += start*10 + end
	}

	fmt.Println(sum)
}
