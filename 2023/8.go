//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Part 1:")
	part1()
	fmt.Println("Part 2:")
	part2()
}

type Choice struct {
	Left  string
	Right string
}

func part1() {
	filePath := "../../AoC-input/2023/8/input.txt"
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

	directions := fileLines[0]
	fileLines = fileLines[2:]

	nodes := make(map[string]Choice)

	for _, line := range fileLines {
		node := line[0:3]
		left := line[7:10]
		right := line[12:15]
		nodes[node] = Choice{left, right}
	}

	currentNode := "AAA"
	steps := 0

	for currentNode != "ZZZ" {
		if directions[steps%len(directions)] == 'L' {
			currentNode = nodes[currentNode].Left
		} else {
			currentNode = nodes[currentNode].Right
		}
		steps++
	}

	fmt.Println(steps)
}

func part2() {
	filePath := "../../AoC-input/2023/8/input.txt"
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

	directions := fileLines[0]
	fileLines = fileLines[2:]

	nodes := make(map[string]Choice)

	for _, line := range fileLines {
		node := line[0:3]
		left := line[7:10]
		right := line[12:15]
		nodes[node] = Choice{left, right}
	}

	currentNodes := []string{}
	activeNodes := []bool{}
	zHistory := []int64{}

	for node := range nodes {
		if node[2] == 'A' {
			currentNodes = append(currentNodes, node)
			activeNodes = append(activeNodes, true)
			zHistory = append(zHistory, 1)
		}
	}

	steps := 0

	for !allInactive(activeNodes) {
		newCurrentNodes := []string{}

		for i, currentNode := range currentNodes {
			if activeNodes[i] {
				if directions[steps%len(directions)] == 'L' {
					newCurrentNodes = append(newCurrentNodes, nodes[currentNode].Left)
				} else {
					newCurrentNodes = append(newCurrentNodes, nodes[currentNode].Right)
				}

				if newCurrentNodes[i][2] == 'Z' {
					zHistory[i] = int64(steps + 1)
					activeNodes[i] = false
				}
			} else {
				newCurrentNodes = append(newCurrentNodes, currentNode)
			}
		}

		currentNodes = newCurrentNodes

		steps++
	}

	result := lcmOfArray(zHistory)

	fmt.Println(result)
}

func allInactive(nodes []bool) bool {
	for _, node := range nodes {
		if node {
			return false
		}
	}
	return true
}

func lcmOfArray(numbers []int64) int64 {
	result := int64(1)
	for _, num := range numbers {
		result = (result * num) / gcd(result, num)
	}
	return result
}

func gcd(a, b int64) int64 {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
