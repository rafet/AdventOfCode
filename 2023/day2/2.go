package main

import (
	"os"
	"strconv"
	"strings"
)

var contraints = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func main() {
	lines := getInputLines()
	p1Sum := 0
	p2Sum := 0

	for _, line := range lines {
		gameId, colorMap := getColorMapOfLine(line)
		if isPossible(colorMap) {
			p1Sum += gameId
		}

		p2Sum += multiplyMapValues(colorMap)
	}

	println("p1 answer:", p1Sum)
	println("p2 answer:", p2Sum)
}

func multiplyMapValues(colorMap map[string]int) int {
	multiple := 1
	for _, count := range colorMap {
		multiple *= count
	}

	return multiple
}

func getColorMapOfLine(line string) (int, map[string]int) {
	values := strings.FieldsFunc(line, func(r rune) bool {
		return r == ' ' || r == ':' || r == ';' || r == ','
	})

	gameId, _ := strconv.Atoi(values[1])

	colorMap := make(map[string]int)
	for i := 2; i < len(values); i += 2 {
		count, _ := strconv.Atoi(values[i])
		colorMap[values[i+1]] = max(colorMap[values[i+1]], count)
	}

	return gameId, colorMap
}

func isPossible(colorMap map[string]int) bool {
	for color, count := range contraints {
		if colorMap[color] > count {
			return false
		}
	}

	return true
}

func getInputLines() []string {
	file, err := os.ReadFile("2023/day2/input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(file), "\n")
	lines = lines[:len(lines)-1]

	return lines
}
