package main

import (
	"regexp"
	"strconv"
	"strings"
)

func Part1(input string) int {
	return totalFromMemory(input)
}

func Part2(input string) int {
	// Note: This will fail if the input beings with "don't()"
	slices := strings.Split(input, "do()")

	runningTotal := 0
	for _, slice := range slices {
		validSection := strings.Split(slice, "don't()")[0]
		runningTotal += totalFromMemory(validSection)
	}

	return runningTotal
}

func totalFromMemory(memory string) int {
	r := regexp.MustCompile(`mul\((\d{1,3}\,\d{1,3})\)`)
	matches := r.FindAllStringSubmatch(memory, -1)

	runningTotal := 0
	for _, match := range matches {
		numPair := strings.Split(match[1], ",")
		left, _ := strconv.Atoi(numPair[0])
		right, _ := strconv.Atoi(numPair[1])
		runningTotal += left * right
	}
	return runningTotal
}
