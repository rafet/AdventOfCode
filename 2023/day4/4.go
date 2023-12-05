package main

import (
	"math"
	"os"
	"strings"
)

func main() {
	lines := getInputLines()

	var sum float64 = 0

	for _, line := range lines {
		fields := strings.FieldsFunc(line, func(c rune) bool {
			return c == ':' || c == '|'
		})

		winnings := strings.Fields(strings.TrimSpace(fields[1]))
		numbers := strings.Fields(strings.TrimSpace(fields[2]))

		cnt := 0

		for _, num := range numbers {
			if contains(winnings, num) {
				cnt += 1
			}
		}

		if cnt > 0 {
			sum += math.Pow(2, float64(cnt-1))
		}
	}

	println(int(sum))

}

func contains(list []string, elem string) bool {
	for _, e := range list {
		if e == elem {
			return true
		}
	}

	return false
}

func getInputLines() []string {
	file, err := os.ReadFile("2023/day4/input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(file), "\n")
	lines = lines[:len(lines)-1]

	return lines
}
