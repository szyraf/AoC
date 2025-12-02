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
	filePath := "../../AoC-input/2023/6/input.txt"
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

	times := []int{}
	distances := []int{}

	fields := strings.Fields(fileLines[0])
	fields = fields[1:]
	for _, field := range fields {
		time, _ := strconv.Atoi(field)
		times = append(times, time)
	}

	fields = strings.Fields(fileLines[1])
	fields = fields[1:]
	for _, field := range fields {
		distance, _ := strconv.Atoi(field)
		distances = append(distances, distance)
	}

	output := 1
	for i := 0; i < len(times); i++ {
		sum := 0
		for j := 0; j <= times[i]; j++ {
			if (times[i]-j)*j > distances[i] {
				sum++
			}
		}
		output *= sum
	}

	fmt.Println(output)
}

func part2() {
	filePath := "../../AoC-input/2023/6/input.txt"
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

	totalTime := ""
	totalDistance := ""

	fields := strings.Fields(fileLines[0])
	fields = fields[1:]
	for _, field := range fields {
		totalTime += field
	}

	fields = strings.Fields(fileLines[1])
	fields = fields[1:]
	for _, field := range fields {
		totalDistance += field
	}

	totalTimeInt, _ := strconv.Atoi(totalTime)
	totalDistanceInt, _ := strconv.Atoi(totalDistance)

	sum := 0
	for j := 0; j <= totalTimeInt; j++ {
		if (totalTimeInt-j)*j > totalDistanceInt {
			sum++
		}
	}

	fmt.Println(sum)
}
