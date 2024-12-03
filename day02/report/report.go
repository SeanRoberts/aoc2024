package report

import (
	"fmt"
	"strconv"
	"strings"
)

type Report struct {
	Values      []int
	Direction   Direction
	MinDistance int
	MaxDistance int
}

type Direction int

const (
	DirectionUnknown Direction = iota
	DirectionAsc
	DirectionDesc
	DirectionFlat
	DirectionMixed
)

func NewReportFromString(line string) (Report, error) {
	strValues := strings.Split(line, " ")
	values := make([]int, len(strValues))
	for i := range strValues {
		value, err := strconv.Atoi(strValues[i])
		if err != nil {
			continue
		}
		values[i] = value
	}

	if len(values) < 2 {
		return Report{}, fmt.Errorf("Report must have at least 2 levels")
	}

	return NewReport(values), nil
}

func NewReport(values []int) Report {
	report := Report{Values: values, Direction: DirectionUnknown, MinDistance: -1}

	var prevValue int
	for i := range report.Values {
		value := report.Values[i]

		if i == 0 {
			prevValue = value
			continue
		}

		distance := abs(value - prevValue)
		if distance > report.MaxDistance {
			report.MaxDistance = distance
		} else if distance < report.MinDistance || report.MinDistance == -1 {
			report.MinDistance = distance
		}

		var direction Direction
		if value > prevValue {
			direction = DirectionAsc
		} else if value < prevValue {
			direction = DirectionDesc
		} else {
			direction = DirectionFlat
		}
		report.updateDirection(direction)
		prevValue = value
	}

	return report
}

func (r Report) IsSafe() bool {
	return r.Direction != DirectionMixed && r.MaxDistance <= 3 && r.MinDistance > 0
}

func (r Report) IsKindaSafe() bool {
	if r.IsSafe() {
		return true
	}

	for i := range r.Values {
		if NewReport(r.valuesWithoutIndex(i)).IsSafe() {
			return true
		}
	}

	return false
}

func (r Report) valuesWithoutIndex(i int) []int {
	newValues := make([]int, len(r.Values))
	copy(newValues, r.Values)
	return append(newValues[:i], newValues[i+1:]...)
}

func (r *Report) updateDirection(newDirection Direction) {
	if newDirection == r.Direction {
		return
	}

	if r.Direction == DirectionUnknown {
		r.Direction = newDirection
		return
	}

	r.Direction = DirectionMixed
}

func abs(i int) int {
	if i >= 0 {
		return i
	}

	return i * -1
}
