package main

import (
	"os"
	"strings"
)

func main() {
	lines := getInputLines()
	sum := 0
	for _, line := range lines {
		sum += 10*getFirstDigit(line) + getLastDigit(line)
	}

	println(sum)
}

func getFirstDigit(s string) int {
	for i := 0; i < len(s); i++ {
		c := s[i]
		if isDigit(c) {
			return int(c - '0')
		}
	}

	return 0
}

func getLastDigit(s string) int {
	for i := len(s) - 1; i >= 0; i-- {
		c := s[i]
		if isDigit(c) {
			return int(c - '0')
		}
	}

	return 0
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func getInputLines() []string {
	file, err := os.ReadFile("2023/1_input.txt")
	if err != nil {
		return []string{}
	}

	return strings.Split(string(file), "\n")
}
