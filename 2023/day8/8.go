package main

import (
	"math/big"
	"os"
	"strings"
)

type path struct {
	d     string
	left  string
	right string
}

var pathMap = map[string]path{}

func main() {
	lines := getInputLines()
	directions := []rune(lines[0])

	for i := 2; i < len(lines); i++ {
		p := getPath(lines[i])
		pathMap[p.d] = p
	}

	cnt := calc("AAA", "ZZZ", directions)

	println("answer 1:", cnt)

	var startPoints []path
	for _, p := range pathMap {
		if p.d[len(p.d)-1] == 'A' {
			startPoints = append(startPoints, p)
		}
	}

	var cnts []int
	for _, s := range startPoints {
		cnts = append(cnts, calc(s.d, "Z", directions))
	}

	println("answer 2:", lcm(cnts))
}

func gcd(a, b *big.Int) *big.Int {
	for b.Cmp(big.NewInt(0)) != 0 {
		a, b = b, new(big.Int).Mod(a, b)
	}
	return a
}

func lcm(numbers []int) int {
	result := big.NewInt(1)
	for _, num := range numbers {
		bigNum := big.NewInt(int64(num))
		gcdResult := gcd(result, bigNum)
		result.Mul(result, bigNum)
		result.Div(result, gcdResult)
	}
	return int(result.Int64())
}

func calc(start, end string, directions []rune) int {
	curr := pathMap[start]
	dirIndex := 0
	for {
		ind := dirIndex % len(directions)

		to := directions[ind]
		if to == 'L' {
			curr = pathMap[curr.left]
		} else {
			curr = pathMap[curr.right]
		}

		dirIndex += 1

		if len(end) == 1 && curr.d[len(curr.d)-1] == end[0] || curr.d == end {
			break
		}
	}

	return dirIndex
}

func getPath(line string) path {
	fields := strings.FieldsFunc(line, func(r rune) bool {
		return r == '(' || r == ')' || r == ',' || r == ' ' || r == '='
	})

	return path{
		d:     fields[0],
		left:  fields[1],
		right: fields[2],
	}
}
func getInputLines() []string {
	file, err := os.ReadFile("2023/day8/input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(file), "\n")
	lines = lines[:len(lines)-1]

	return lines
}
