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
	filePath := "../../AoC-input/2023/5/input.txt"
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

	seeds := []int{}
	seedsLock := []bool{}
	for _, line := range fileLines {
		numbers := strings.Split(line, " ")
		if numbers[0] == "seeds:" {
			for _, seed := range numbers[1:] {
				seedInt, _ := strconv.Atoi(seed)
				seeds = append(seeds, seedInt)
				seedsLock = append(seedsLock, false)
			}

		} else if len(numbers) == 3 {
			for i := 0; i < len(seeds); i++ {
				if !seedsLock[i] {
					seed := seeds[i]
					destinationStart, _ := strconv.Atoi(numbers[0])
					sourceStart, _ := strconv.Atoi(numbers[1])
					length, _ := strconv.Atoi(numbers[2])
					if seed >= sourceStart && seed < sourceStart+length {
						seeds[i] = destinationStart + (seed - sourceStart)
						seedsLock[i] = true
					}
				}
			}

		} else if line == "" {
			for i := 0; i < len(seedsLock); i++ {
				seedsLock[i] = false
			}
		}
	}

	minLocation := seeds[0]
	for _, seed := range seeds {
		if seed < minLocation {
			minLocation = seed
		}
	}
	fmt.Println(minLocation)
}

type Range struct {
	start int64
	end   int64
	lock  bool
}

func part2() {
	filePath := "../../AoC-input/2023/5/input.txt"
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

	seeds := []Range{}
	for _, line := range fileLines {
		numbers := strings.Split(line, " ")
		if numbers[0] == "seeds:" {
			for i, seed := range numbers[1:] {
				if i%2 == 1 {
					continue
				}
				seedInt, _ := strconv.ParseInt(seed, 10, 64)
				length, _ := strconv.ParseInt(numbers[1+i+1], 10, 64)
				seeds = append(seeds, Range{seedInt, seedInt + length - 1, false})
			}

		} else if len(numbers) == 3 {
			for i := 0; i < len(seeds); i++ {
				if !seeds[i].lock {
					seed := seeds[i]
					destinationStart, _ := strconv.ParseInt(numbers[0], 10, 64)
					sourceStart, _ := strconv.ParseInt(numbers[1], 10, 64)
					length, _ := strconv.ParseInt(numbers[2], 10, 64)

					if seed.start >= sourceStart && seed.start < sourceStart+length && seed.end >= sourceStart && seed.end < sourceStart+length {
						newStart := destinationStart + (seed.start - sourceStart)
						newEnd := destinationStart + (seed.end - sourceStart)

						seeds[i].start = newStart
						seeds[i].end = newEnd
						seeds[i].lock = true
					} else if seed.start >= sourceStart && seed.start < sourceStart+length {
						seeds = append(seeds, Range{sourceStart + length, seed.end, false})
						newStart := destinationStart + (seed.start - sourceStart)
						newEnd := newStart + ((sourceStart + length) - seed.start - 1)

						seeds[i].start = newStart
						seeds[i].end = newEnd
						seeds[i].lock = true
					} else if seed.end >= sourceStart && seed.end < sourceStart+length {
						seeds = append(seeds, Range{seed.start, sourceStart - 1, false})
						newStart := destinationStart
						newEnd := newStart + (seed.end - sourceStart)

						seeds[i].start = newStart
						seeds[i].end = newEnd
						seeds[i].lock = true
					} else if seed.start < sourceStart && seed.end > sourceStart+length {
						seeds = append(seeds, Range{seed.start, sourceStart - 1, false})
						seeds = append(seeds, Range{sourceStart + length, seed.end, false})
						newStart := destinationStart
						newEnd := newStart + (seed.end - seed.start)

						seeds[i].start = newStart
						seeds[i].end = newEnd
						seeds[i].lock = true
					}
				}
			}

		} else if line == "" {
			for i := 0; i < len(seeds); i++ {
				seeds[i].lock = false
			}
		}
	}

	minLocation := seeds[0].start

	for _, seed := range seeds {
		if seed.start < minLocation {
			minLocation = seed.start
		}
	}

	fmt.Println(minLocation)
}
