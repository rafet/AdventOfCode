package main

import (
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	lines := getInputLines()

	extrapolatedValues := 0
	part2Values := 0
	for _, line := range lines {
		fields := sliceToInts(strings.Fields(line))
		var tree = [][]int{
			fields,
		}

		curr := fields

		for {
			var t []int
			for i := 0; i < len(curr)-1; i++ {
				t = append(t, curr[i+1]-curr[i])
			}

			tree = append(tree, t)

			if isAllZeroes(t) {
				break
			} else {
				curr = t

			}
		}

		extVal := 0
		for i, t := range tree {
			extVal += t[len(t)-1]
			part2Values += int(math.Pow(-1, float64(i))) * t[0]
		}

		extrapolatedValues += extVal
	}

	println("answer 1:", extrapolatedValues)
	println("answer 2:", part2Values)
}

func isAllZeroes(slice []int) bool {
	for _, i := range slice {
		if i != 0 {
			return false
		}
	}

	return true
}

func sliceToInts(slice []string) []int {
	var ints []int
	for _, s := range slice {
		i, _ := strconv.Atoi(s)
		ints = append(ints, i)
	}

	return ints
}

func getInputLines() []string {
	file, err := os.ReadFile("2023/day9/input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(file), "\n")
	lines = lines[:len(lines)-1]

	return lines
}
