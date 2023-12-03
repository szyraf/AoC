/*
Description:

--- Day 3: Gear Ratios ---
You and the Elf eventually reach a gondola lift station; he says the gondola lift will take you up to the water source, but this is as far as he can bring you. You go inside.

It doesn't take long to find the gondolas, but there seems to be a problem: they're not moving.

"Aaah!"

You turn around to see a slightly-greasy Elf with a wrench and a look of surprise. "Sorry, I wasn't expecting anyone! The gondola lift isn't working right now; it'll still be a while before I can fix it." You offer to help.

The engineer explains that an engine part seems to be missing from the engine, but nobody can figure out which one. If you can add up all the part numbers in the engine schematic, it should be easy to work out which part is missing.

The engine schematic (your puzzle input) consists of a visual representation of the engine. There are lots of numbers and symbols you don't really understand, but apparently any number adjacent to a symbol, even diagonally, is a "part number" and should be included in your sum. (Periods (.) do not count as a symbol.)

Here is an example engine schematic:

467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..
In this schematic, two numbers are not part numbers because they are not adjacent to a symbol: 114 (top right) and 58 (middle right). Every other number is adjacent to a symbol and so is a part number; their sum is 4361.

Of course, the actual engine schematic is much larger. What is the sum of all of the part numbers in the engine schematic?
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
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

func isSymbol(c string) bool {
	re := regexp.MustCompile(`^[^0-9.]$`)
	return re.MatchString(c)
}

func part1() {
	filePath := "./3/input.txt"
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
		for j := 0; j < len(line); j++ {
			if isNumber(string(line[j])) {
				number := ""
				for j < len(line) && isNumber(string(line[j])) {
					number += string(line[j])
					j++
				}

				isPartNumber := false
				if j < len(line) {
					if isSymbol(string(line[j])) {
						isPartNumber = true
					}
				}
				if !isPartNumber {
					if j-len(number)-1 >= 0 {
						if isSymbol(string(line[j-len(number)-1])) {
							isPartNumber = true
						}
					}
				}
				if !isPartNumber {
					for k := j - len(number) - 1; k <= j; k++ {
						if k >= 0 && k < len(line) {
							if i > 0 && isSymbol(string(fileLines[i-1][k])) {
								isPartNumber = true
								break
							}
							if i < len(fileLines)-1 && isSymbol(string(fileLines[i+1][k])) {
								isPartNumber = true
								break
							}
						}
					}
				}

				if isPartNumber {
					numberInt, _ := strconv.Atoi(number)
					sum += numberInt
				}
			}
		}
	}

	fmt.Println(sum)
}

/*
--- Part Two ---
The engineer finds the missing part and installs it in the engine! As the engine springs to life, you jump in the closest gondola, finally ready to ascend to the water source.

You don't seem to be going very fast, though. Maybe something is still wrong? Fortunately, the gondola has a phone labeled "help", so you pick it up and the engineer answers.

Before you can explain the situation, she suggests that you look out the window. There stands the engineer, holding a phone in one hand and waving with the other. You're going so slowly that you haven't even left the station. You exit the gondola.

The missing part wasn't the only issue - one of the gears in the engine is wrong. A gear is any * symbol that is adjacent to exactly two part numbers. Its gear ratio is the result of multiplying those two numbers together.

This time, you need to find the gear ratio of every gear and add them all up so that the engineer can figure out which gear needs to be replaced.

Consider the same engine schematic again:

467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..
In this schematic, there are two gears. The first is in the top left; it has part numbers 467 and 35, so its gear ratio is 16345. The second gear is in the lower right; its gear ratio is 451490. (The * adjacent to 617 is not a gear because it is only adjacent to one part number.) Adding up all of the gear ratios produces 467835.

What is the sum of all of the gear ratios in your engine schematic?
*/

type Coord struct {
	x int
	y int
}

type Output struct {
	number   int
	quantity int
}

func part2() {
	filePath := "./3/input.txt"
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

	coords := map[Coord]Output{}

	for i, line := range fileLines {
		for j := 0; j < len(line); j++ {
			if isNumber(string(line[j])) {
				number := ""
				for j < len(line) && isNumber(string(line[j])) {
					number += string(line[j])
					j++
				}

				symbols := 0
				lastSymbolCoords := []int{-1, -1}
				if j < len(line) {
					if isSymbol(string(line[j])) {
						symbols++
						lastSymbolCoords = []int{i, j}
					}
				}
				if j-len(number)-1 >= 0 {
					if isSymbol(string(line[j-len(number)-1])) {
						symbols++
						lastSymbolCoords = []int{i, j - len(number) - 1}
					}
				}
				for k := j - len(number) - 1; k <= j; k++ {
					if k >= 0 && k < len(line) {
						if i > 0 && isSymbol(string(fileLines[i-1][k])) {
							symbols++
							lastSymbolCoords = []int{i - 1, k}
						}
						if i < len(fileLines)-1 && isSymbol(string(fileLines[i+1][k])) {
							symbols++
							lastSymbolCoords = []int{i + 1, k}
						}
					}
				}

				if symbols == 1 {
					numberInt, _ := strconv.Atoi(number)
					c := Coord{lastSymbolCoords[0], lastSymbolCoords[1]}
					if coords[c].number == 0 {
						coords[c] = Output{numberInt, coords[c].quantity + 1}
					} else {
						coords[c] = Output{coords[c].number * numberInt, coords[c].quantity + 1}
					}
				}
			}
		}
	}

	sum := 0

	for _, output := range coords {
		if output.quantity == 2 {
			sum += output.number
		}
	}

	fmt.Println(sum)
}
