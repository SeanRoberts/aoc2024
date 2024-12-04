package main

import (
	"slices"
	"strings"
)

func Part1(input []string) int {
	totalInstances := 0
	graph := make([][]string, len(input))
	for i, line := range input {
		graph[i] = strings.Split(line, "")
	}

	for y, line := range graph {
		for x, letter := range line {
			if letter != "X" && letter != "S" {
				continue
			}

			totalInstances += forwardXmas(graph, x, y)
			totalInstances += verticalForwardXmas(graph, x, y)
			totalInstances += seDiagonalXmas(graph, x, y)
			totalInstances += swDiagonalXmas(graph, x, y)
		}
	}
	return totalInstances
}

func Part2(input []string) int {
	totalInstances := 0
	graph := make([][]string, len(input))
	for i, line := range input {
		graph[i] = strings.Split(line, "")
	}

	for y, line := range graph {
		for x, letter := range line {
			if letter != "A" || y < 1 || x < 1 || y == len(graph)-1 || x == len(line)-1 {
				continue
			}

			letters1 := make([]string, 3)
			letters2 := make([]string, 3)
			for i := -1; i < 2; i++ {
				letters1[i+1] = graph[y+i][x+i]
				letters2[i+1] = graph[y+i*-1][x+i]
			}
			slices.Sort(letters1)
			slices.Sort(letters2)
			if strings.Join(letters1, "") == "AMS" && strings.Join(letters2, "") == "AMS" {
				totalInstances++
			}
		}
	}

	return totalInstances
}

func forwardXmas(graph [][]string, x int, y int) int {
	line := graph[y]
	if len(line) < x+4 {
		return 0
	}

	word := strings.Join(line[x:x+4], "")
	if word == "XMAS" || word == "SAMX" {
		return 1
	}

	return 0
}

func verticalForwardXmas(graph [][]string, x int, y int) int {
	if len(graph) < y+4 {
		return 0
	}

	letters := make([]string, 4)
	for i := 0; i < 4; i++ {
		letters[i] = graph[y+i][x]
	}
	word := strings.Join(letters, "")

	if word == "XMAS" || word == "SAMX" {
		return 1
	}

	return 0
}

func seDiagonalXmas(graph [][]string, x int, y int) int {
	if len(graph) < y+4 || len(graph[y]) < x+4 {
		return 0
	}

	letters := make([]string, 4)
	for i := 0; i < 4; i++ {
		letters[i] = graph[y+i][x+i]
	}

	word := strings.Join(letters, "")

	if word == "XMAS" || word == "SAMX" {
		return 1
	}

	return 0
}

func swDiagonalXmas(graph [][]string, x int, y int) int {
	if len(graph) < y+4 || x < 3 {
		return 0
	}

	letters := make([]string, 4)
	for i := 0; i < 4; i++ {
		letters[i] = graph[y+i][x-i]
	}

	word := strings.Join(letters, "")

	if word == "XMAS" || word == "SAMX" {
		return 1
	}

	return 0
}
