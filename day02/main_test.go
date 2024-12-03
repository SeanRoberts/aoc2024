package main

import (
	"testing"

	"github.com/SeanRoberts/aoc2024/day02/report"
)

func TestIsSafe(t *testing.T) {
	tests := []struct {
		values        []int
		safeWithNoPad bool
		safeWithPad   bool
	}{
		{[]int{7, 6, 4, 2, 1}, true, true},
		{[]int{1, 2, 7, 8, 9}, false, false},
		{[]int{9, 7, 6, 2, 1}, false, false},
		{[]int{1, 3, 2, 4, 5}, false, true},
		{[]int{8, 6, 4, 4, 1}, false, true},
		{[]int{1, 3, 6, 7, 9}, true, true},
		{[]int{48, 46, 47, 49, 51, 54, 56}, false, true},
		{[]int{1, 1, 2, 3, 4, 5}, false, true},
		{[]int{1, 2, 3, 4, 5, 5}, false, true},
		{[]int{5, 1, 2, 3, 4, 5}, false, true},
		{[]int{1, 4, 3, 2, 1}, false, true},
		{[]int{1, 6, 7, 8, 9}, false, true},
		{[]int{1, 2, 3, 4, 3}, false, true},
		{[]int{9, 8, 7, 6, 7}, false, true},
		{[]int{7, 10, 8, 10, 11}, false, true},
		{[]int{29, 28, 27, 25, 26, 25, 22, 20}, false, true},
	}

	for _, tt := range tests {
		report := report.NewReport(tt.values)
		noPadResult := report.IsSafe()
		padResult := report.IsKindaSafe()

		if noPadResult != tt.safeWithNoPad {
			t.Error("noPad result incorrect for", tt.values, "got", noPadResult, "expected", tt.safeWithNoPad)
		}

		if padResult != tt.safeWithPad {
			t.Error("pad result incorrect for", tt.values, "got", padResult, "expected", tt.safeWithPad)
		}
	}
}
