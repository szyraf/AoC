//go:build ignore

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
	filePath := "../../AoC-input/2023/3/input.txt"
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

type Coord struct {
	x int
	y int
}

type Output struct {
	number   int
	quantity int
}

func part2() {
	filePath := "../../AoC-input/2023/3/input.txt"
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
