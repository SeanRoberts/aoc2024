package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	bytes, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	str := string(bytes)

	lines := strings.Split(str, "\n")

	leftcol := []int{}
	rightcol := []int{}

	for _, s := range lines {
		vals := strings.Split(s, "   ")
		if len(vals) < 2 {
			break
		}
		leftint, err1 := strconv.Atoi(vals[0])
		rightint, err2 := strconv.Atoi(vals[1])
		if err1 != nil || err2 != nil {
			fmt.Println("One or more values are not convertable to integers")
			return
		}

		leftcol = append(leftcol, leftint)
		rightcol = append(rightcol, rightint)
	}

	slices.Sort(leftcol)
	slices.Sort(rightcol)

	fmt.Println("totalDistance", Part1(leftcol, rightcol))
	fmt.Println("similarity", Part2(leftcol, rightcol))
}
