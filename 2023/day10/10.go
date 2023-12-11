package main

import (
	"os"
	"strings"
)

type Direction string

var Right Direction = "r"
var Left Direction = "l"
var Top Direction = "t"
var Bottom Direction = "b"

var acceptableDirs = map[rune][]Direction{
	'|': {Top, Bottom},
	'-': {Left, Right},
	'L': {Bottom, Left},
	'J': {Bottom, Right},
	'7': {Right, Top},
	'F': {Left, Top},
	'.': {},
}

var (
	pipeToDirection = map[rune]map[Direction]Direction{
		'|': {
			Bottom: Bottom,
			Top:    Top,
		},
		'-': {
			Right: Right,
			Left:  Left,
		},
		'L': {
			Bottom: Right,
			Left:   Top,
		},
		'J': {
			Bottom: Left,
			Right:  Top,
		},
		'7': {
			Right: Bottom,
			Top:   Left,
		},
		'F': {
			Left: Bottom,
			Top:  Right,
		},
	}
)

type point struct {
	i int
	j int
}

func (p point) isValid() bool {
	return !(p.i < 0 || p.j < 0 || p.i >= len(matrix) || p.j >= len(matrix[0]))
}

func (p point) val(matrix [][]rune) rune {
	return matrix[p.i][p.j]
}

var matrix [][]rune

func main() {
	lines := getInputLines()

	for _, line := range lines {
		matrix = append(matrix, []rune(line))
	}

	S := getStartingPoint(matrix)

	move(matrix, Bottom, point{S.i + 1, S.j})

	println("answer 1:", len(allPoints)/2+1)

}

var allPoints []point

func move(matrix [][]rune, prevMove Direction, curr point) bool {
	v := curr.val(matrix)
	if v == 'S' {
		return true
	}

	nextMove := pipeToDirection[v][prevMove]
	if nextMove == "" {
		return false
	}

	if isValidPath(matrix, curr, nextMove) {
		nextPoint := getNextPoint(curr, nextMove)
		allPoints = append(allPoints, nextPoint)

		return move(matrix, nextMove, nextPoint)
	}

	return false
}

func getNextPoint(curr point, direction Direction) point {
	offset := getDirOffset(direction)
	return point{curr.i + offset.i, curr.j + offset.j}
}

func isValidPath(matrix [][]rune, from point, dir Direction) bool {
	v := from.val(matrix)
	if v == '.' {
		return false
	}
	offset := getDirOffset(dir)
	destinationPoint := point{offset.i + from.i, offset.j + from.j}
	if !destinationPoint.isValid() {
		return false
	}

	destinationV := destinationPoint.val(matrix)
	if destinationV == 'S' {
		return true
	}

	acceptableDirections, _ := acceptableDirs[destinationV]

	return contains(acceptableDirections, dir)
}

func contains[T comparable](list []T, elem T) bool {
	for _, e := range list {
		if e == elem {
			return true
		}
	}

	return false
}

func getDirOffset(dir Direction) point {
	switch dir {
	case Right:
		return point{0, 1}
	case Left:
		return point{0, -1}
	case Top:
		return point{-1, 0}
	case Bottom:
		return point{1, 0}

	}

	return point{}
}

func getStartingPoint(matrix [][]rune) point {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 'S' {
				return point{i, j}
			}
		}
	}

	return point{-1, -1}
}

func getInputLines() []string {
	file, err := os.ReadFile("2023/day10/input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(file), "\n")
	lines = lines[:len(lines)-1]

	return lines
}
