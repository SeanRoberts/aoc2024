package common

import (
	"os"
	"strings"
)

func FileToString(path string) (string, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func FileToLines(path string) ([]string, error) {
	str, err := FileToString(path)
	if err != nil {
		return []string{}, err
	}

	lines := strings.Split(str, "\n")
	if lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	return lines, nil
}

func FileToGrid(path string) (Grid, error) {
	grid := make(Grid)
	str, err := FileToString(path)
	if err != nil {
		return grid, err
	}
	grid.FillFromString(str)
	return grid, nil
}
