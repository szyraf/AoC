/*
Description:

*/

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Part 1:")
	part1()
	// fmt.Println("Part 2:")
	// part2()
}

func part1() {
	filePath := "./blank/testinput.txt"
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

}

/*
--- Part Two ---

*/

func part2() {
	filePath := "./blank/testinput.txt"
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

}
