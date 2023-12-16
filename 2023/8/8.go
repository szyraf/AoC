/*
Description:

--- Day 8: Haunted Wasteland ---
You're still riding a camel across Desert Island when you spot a sandstorm quickly approaching. When you turn to warn the Elf, she disappears before your eyes! To be fair, she had just finished warning you about ghosts a few minutes ago.

One of the camel's pouches is labeled "maps" - sure enough, it's full of documents (your puzzle input) about how to navigate the desert. At least, you're pretty sure that's what they are; one of the documents contains a list of left/right instructions, and the rest of the documents seem to describe some kind of network of labeled nodes.

It seems like you're meant to use the left/right instructions to navigate the network. Perhaps if you have the camel follow the same instructions, you can escape the haunted wasteland!

After examining the maps for a bit, two nodes stick out: AAA and ZZZ. You feel like AAA is where you are now, and you have to follow the left/right instructions until you reach ZZZ.

This format defines each node of the network individually. For example:

RL

AAA = (BBB, CCC)
BBB = (DDD, EEE)
CCC = (ZZZ, GGG)
DDD = (DDD, DDD)
EEE = (EEE, EEE)
GGG = (GGG, GGG)
ZZZ = (ZZZ, ZZZ)
Starting with AAA, you need to look up the next element based on the next left/right instruction in your input. In this example, start with AAA and go right (R) by choosing the right element of AAA, CCC. Then, L means to choose the left element of CCC, ZZZ. By following the left/right instructions, you reach ZZZ in 2 steps.

Of course, you might not find ZZZ right away. If you run out of left/right instructions, repeat the whole sequence of instructions as necessary: RL really means RLRLRLRLRLRLRLRL... and so on. For example, here is a situation that takes 6 steps to reach ZZZ:

LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)
Starting at AAA, follow the left/right instructions. How many steps are required to reach ZZZ?
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
	fmt.Println("Part 2:")
	part2()
}

type Choice struct {
	Left  string
	Right string
}

func part1() {
	filePath := "./8/input.txt"
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

/*
--- Part Two ---
The sandstorm is upon you and you aren't any closer to escaping the wasteland. You had the camel follow the instructions, but you've barely left your starting position. It's going to take significantly more steps to escape!

What if the map isn't for people - what if the map is for ghosts? Are ghosts even bound by the laws of spacetime? Only one way to find out.

After examining the maps a bit longer, your attention is drawn to a curious fact: the number of nodes with names ending in A is equal to the number ending in Z! If you were a ghost, you'd probably just start at every node that ends with A and follow all of the paths at the same time until they all simultaneously end up at nodes that end with Z.

For example:

LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)
Here, there are two starting nodes, 11A and 22A (because they both end with A). As you follow each left/right instruction, use that instruction to simultaneously navigate away from both nodes you're currently on. Repeat this process until all of the nodes you're currently on end with Z. (If only some of the nodes you're on end with Z, they act like any other node and you continue as normal.) In this example, you would proceed as follows:

Step 0: You are at 11A and 22A.
Step 1: You choose all of the left paths, leading you to 11B and 22B.
Step 2: You choose all of the right paths, leading you to 11Z and 22C.
Step 3: You choose all of the left paths, leading you to 11B and 22Z.
Step 4: You choose all of the right paths, leading you to 11Z and 22B.
Step 5: You choose all of the left paths, leading you to 11B and 22C.
Step 6: You choose all of the right paths, leading you to 11Z and 22Z.
So, in this example, you end up entirely on nodes that end in Z after 6 steps.

Simultaneously start on every node that ends with A. How many steps does it take before you're only on nodes that end with Z?
*/

func part2() {
	filePath := "./8/input.txt"
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
