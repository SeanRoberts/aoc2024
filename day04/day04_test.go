package main

import (
	"testing"

	"github.com/SeanRoberts/aoc2024/common"
)

func TestPart1(t *testing.T) {
	input, err := common.FileToGrid("./input_test.txt")
	if err != nil {
		t.Error("Failed to read input file.")
	}
	result := Part1(input)
	if result != 18 {
		t.Errorf("Expected 18 instances of XMAS, got %v", result)
	}
}

func TestPart2(t *testing.T) {
	input, err := common.FileToGrid("./input_test.txt")
	if err != nil {
		t.Error("Failed to read input file.")
	}
	result := Part2(input)
	if result != 9 {
		t.Errorf("Expected 9 MAS X-es, got %v", result)
	}
}
