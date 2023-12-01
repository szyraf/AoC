/*
Description:

--- Day 1: Trebuchet?! ---
Something is wrong with global snow production, and you've been selected to take a look. The Elves have even given you a map; on it, they've used stars to mark the top fifty locations that are likely to be having problems.

You've been doing this long enough to know that to restore snow operations, you need to check all fifty stars by December 25th.

Collect stars by solving puzzles. Two puzzles will be made available on each day in the Advent calendar; the second puzzle is unlocked when you complete the first. Each puzzle grants one star. Good luck!

You try to ask why they can't just use a weather machine ("not powerful enough") and where they're even sending you ("the sky") and why your map looks mostly blank ("you sure ask a lot of questions") and hang on did you just say the sky ("of course, where do you think snow comes from") when you realize that the Elves are already loading you into a trebuchet ("please hold still, we need to strap you in").

As they're making the final adjustments, they discover that their calibration document (your puzzle input) has been amended by a very young Elf who was apparently just excited to show off her art skills. Consequently, the Elves are having trouble reading the values on the document.

The newly-improved calibration document consists of lines of text; each line originally contained a specific calibration value that the Elves now need to recover. On each line, the calibration value can be found by combining the first digit and the last digit (in that order) to form a single two-digit number.

For example:

1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet
In this example, the calibration values of these four lines are 12, 38, 15, and 77. Adding these together produces 142.

Consider your entire calibration document. What is the sum of all of the calibration values?
*/

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
	filePath := "1/input.txt"
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

/*
--- Part Two ---
Your calculation isn't quite right. It looks like some of the digits are actually spelled out with letters: one, two, three, four, five, six, seven, eight, and nine also count as valid "digits".

Equipped with this new information, you now need to find the real first and last digit on each line. For example:

two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen
In this example, the calibration values are 29, 83, 13, 24, 42, 14, and 76. Adding these together produces 281.

What is the sum of all of the calibration values?
*/

func part2() {
	filePath := "1/input.txt"
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
