package main

import (
	"slices"
	"strings"

	"github.com/SeanRoberts/aoc2024/common"
)

func Part1(grid common.Grid) int {
	totalInstances := 0
	for coords, letter := range grid {
		if letter != "X" && letter != "S" {
			continue
		}

		result := []bool{
			forwardXmas(grid, coords),
			downwardXmas(grid, coords),
			diagonalXmas(grid, coords, 1),
			diagonalXmas(grid, coords, -1),
		}

		for _, val := range result {
			if val {
				totalInstances++
			}
		}
	}
	return totalInstances
}

func Part2(grid common.Grid) int {
	totalInstances := 0
	for coords, letter := range grid {
		if letter != "A" {
			continue
		}

		tl := grid.RelativeValueAt(coords, -1, -1)
		tr := grid.RelativeValueAt(coords, 1, -1)
		bl := grid.RelativeValueAt(coords, -1, 1)
		br := grid.RelativeValueAt(coords, 1, 1)

		word1 := strings.Join([]string{tl, letter, br}, "")
		word2 := strings.Join([]string{tr, letter, bl}, "")
		if (word1 == "MAS" || word1 == "SAM") && (word2 == "MAS" || word2 == "SAM") {
			totalInstances++
		}
	}

	return totalInstances
}

func forwardXmas(grid common.Grid, coords common.Coords) bool {
	wanted := getWantedLetters(grid[coords])

	for i := range wanted {
		if grid.RelativeValueAt(coords, i, 0) != wanted[i] {
			return false
		}
	}

	return true
}

func downwardXmas(grid common.Grid, coords common.Coords) bool {
	wanted := getWantedLetters(grid[coords])

	for i := range wanted {
		if grid.RelativeValueAt(coords, 0, i) != wanted[i] {
			return false
		}
	}

	return true
}

func diagonalXmas(grid common.Grid, coords common.Coords, xMultiplier int) bool {
	wanted := getWantedLetters(grid[coords])
	for i := range wanted {
		if grid.RelativeValueAt(coords, i*xMultiplier, i) != wanted[i] {
			return false
		}
	}

	return true
}

func getWantedLetters(letter string) []string {
	wanted := []string{"X", "M", "A", "S"}
	if letter == "S" {
		slices.Reverse(wanted)
	}
	return wanted
}
