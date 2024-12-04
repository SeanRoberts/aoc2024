package main

import (
	"fmt"

	"github.com/SeanRoberts/aoc2024/common"
)

func main() {
	input, err := common.FileToLines("./input.txt")
	if err != nil {
		panic("Couldn't read input file")
	}

	fmt.Println("Part 1:", Part1(input))
	fmt.Println("Part 1:", Part2(input))
}
