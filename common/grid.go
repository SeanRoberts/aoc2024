package common

import "strings"

type Coords struct {
	X int
	Y int
}

type Grid map[Coords]string

func (g *Grid) FillFromString(s string) {
	lines := strings.Split(s, "\n")
	for y, line := range lines {
		chars := strings.Split(line, "")
		for x, char := range chars {
			(*g)[Coords{x, y}] = char
		}
	}
}

func (g Grid) ValueAt(x int, y int) string {
	return g[Coords{x, y}]
}

func (g Grid) RelativeValueAt(start Coords, distanceX int, distanceY int) string {
	return g[Coords{start.X + distanceX, start.Y + distanceY}]
}
