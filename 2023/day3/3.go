package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type point struct {
	i, j int
}

var size = 140

var gearMap = make(map[string][]int)

func (p point) isValid() bool {
	return (p.i >= 0 && p.i < size) && (p.j >= 0 && p.j < size)
}

var matrix = make([][]rune, size)

func main() {
	lines := getInputLines()
	result := 0

	fillMatrix(lines)

	// Part 1
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if isNumber(matrix[i][j]) {
				val, length := getValAndLength(i, j)
				adjacentPoints := getAdjacentPoints(i, j, length)
				if hasSymbol(adjacentPoints, val) {
					result += val
				}

				j += length - 1
			}
		}
	}

	// Part 2
	gearResult := 0

	for _, gearValues := range gearMap {
		if len(gearValues) > 1 {
			mult := 1
			for _, gVal := range gearValues {
				mult *= gVal
			}

			gearResult += mult
		}
	}

	println("part1 answer:", result)
	println("part2 answer", gearResult)
}

func fillMatrix(lines []string) {
	for i, line := range lines {
		matrix[i] = make([]rune, size)
		for j, c := range line {
			matrix[i][j] = c
		}
	}
}

func isNumber(v rune) bool {
	return v >= '0' && v <= '9'
}

func hasSymbol(points []point, val int) bool {
	for _, p := range points {
		if matrix[p.i][p.j] != '.' {
			pointStr := fmt.Sprintf("%d:%d", p.i, p.j)
			gearMap[pointStr] = append(gearMap[pointStr], val)
			return true
		}
	}
	return false
}

func getAdjacentPoints(i, j, length int) []point {
	var adjacentPoints []point

	for jj := j; jj < j+length; jj++ {
		top := point{i - 1, jj}
		bottom := point{i + 1, jj}

		if top.isValid() {
			adjacentPoints = append(adjacentPoints, top)
		}

		if bottom.isValid() {
			adjacentPoints = append(adjacentPoints, bottom)
		}
	}

	corners := []point{
		{i - 1, j - 1},      // top left
		{i, j - 1},          // left
		{i + 1, j - 1},      //bottom left
		{i - 1, j + length}, // top right
		{i, j + length},     // right
		{i + 1, j + length}, // bottom right
	}
	for pi := 0; pi < len(corners); pi++ {
		if corners[pi].isValid() {
			adjacentPoints = append(adjacentPoints, corners[pi])
		}
	}

	return adjacentPoints

}

func getValAndLength(i, j int) (int, int) {
	chars := ""
	for jj := j; jj < 140; jj++ {
		if isNumber(matrix[i][jj]) {
			chars += string(matrix[i][jj])
		} else {
			break
		}
	}

	val, _ := strconv.Atoi(chars)
	return val, len(chars)

}

func getInputLines() []string {
	file, err := os.ReadFile("2023/day3/input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(file), "\n")
	lines = lines[:len(lines)-1]

	return lines
}
