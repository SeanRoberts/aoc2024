package main

import (
	"fmt"

	"github.com/SeanRoberts/aoc2024/common"
)

func Part1(leftcol []int, rightcol []int) int {
	totalDistance := 0

	for i := range leftcol {
		diff := common.Abs(rightcol[i] - leftcol[i])
		fmt.Println(leftcol[i], rightcol[i], diff)
		totalDistance += diff
	}

	return totalDistance
}

func Part2(leftcol []int, rightcol []int) int {
	similarity := 0
	counts := make(map[int]int)

	for i := range rightcol {
		val := rightcol[i]
		counts[val] = counts[val] + 1
	}

	for i := range leftcol {
		v := leftcol[i]
		count := counts[v]
		similarity += v * count
	}

	return similarity
}
