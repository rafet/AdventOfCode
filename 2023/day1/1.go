package main

import (
	"os"
	"strings"
)

var numbersMap = map[string]int{
	"zero":  0,
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func main() {
	// part 1
	lines := getInputLines()
	sum := 0

	for _, line := range lines {
		sum += 10*getFirstDigitWithOnlyNumbers(line) + getLastDigitWithOnlyNumbers(line)
	}

	println("part1 answer:", sum)

	// part 2
	sum = 0
	for _, line := range lines {
		sum += 10*getFirstDigitWithCharacters(line) + getLastDigitWithCharacters(line)
	}

	println("part2 answer:", sum)

}

func getFirstDigitWithOnlyNumbers(s string) int {
	for i := 0; i < len(s); i++ {
		c := s[i]
		if isDigit(c) {
			return int(c - '0')
		}
	}

	return 0
}

func getLastDigitWithOnlyNumbers(s string) int {
	for i := len(s) - 1; i >= 0; i-- {
		c := s[i]
		if isDigit(c) {
			return int(c - '0')
		}
	}

	return 0
}

func getFirstDigitWithCharacters(line string) int {
	for index, letter := range line {
		if isDigit(byte(letter)) {
			return int(letter - '0')
		} else {
			if val := getNumber(line[:index+1], strings.HasSuffix); val != 0 {
				return val
			}
		}
	}

	return 0
}

func getLastDigitWithCharacters(line string) int {
	for index := len(line) - 1; index >= 0; index-- {
		letter := line[index]
		if isDigit(letter) {
			return int(letter - '0')
		} else {
			if val := getNumber(line[index:], strings.HasPrefix); val != 0 {
				return val
			}
		}
	}

	return 0
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func getInputLines() []string {
	input := readFile("2023/1_input.txt")

	return strings.Split(input, "\n")
}

func readFile(fileName string) string {
	file, err := os.ReadFile(fileName)
	if err != nil {
		return ""
	}

	return string(file)
}

func getNumber(word string, fn func(string, string) bool) int {
	for key, val := range numbersMap {
		if fn(word, key) {
			return val
		}
	}

	return 0
}
