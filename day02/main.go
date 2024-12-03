package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/SeanRoberts/aoc2024/day02/report"
)

func main() {
	bytes, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	str := string(bytes)

	lines := strings.Split(str, "\n")

	fmt.Println("total reports", len(lines))

	safeReportCount := 0
	paddedSafeReportCount := 0

	for _, line := range lines {
		report, err := report.NewReportFromString(line)

		if err != nil {
			continue
		}

		if report.IsSafe() {
			safeReportCount++
		}

		if report.IsKindaSafe() {
			paddedSafeReportCount++
		}
	}

	fmt.Println("safeReportCount", safeReportCount)
	fmt.Println("paddedSafeReportCount", paddedSafeReportCount)
}
