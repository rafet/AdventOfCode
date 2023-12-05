package main

import (
	"math"
	"os"
	"strconv"
	"strings"
)

type info struct {
	destinationStart int
	sourceStart      int
	rangeLength      int
}

var maps = map[string][]info{}

var converters []string

func main() {
	lines := getInputLines()
	seeds := getSeeds(lines[0])

	mapLines := lines[2:]
	readMaps(mapLines)

	lowestLocation := math.MaxInt
	for _, seed := range seeds {
		val := seed
		for _, converter := range converters {
			val = convert(converter, val)
		}

		lowestLocation = min(lowestLocation, val)
	}

	println("answer part1:", lowestLocation)

	seedPairs := getSeedPairs(lines[0])

	lowestLocation2 := math.MaxInt
	for _, pair := range seedPairs {
		for val := pair[0]; val < pair[0]+pair[1]; val++ {
			v := val
			for _, converter := range converters {
				v = convert(converter, v)
			}

			lowestLocation2 = min(lowestLocation2, v)
		}

	}

	println("answer part2:", lowestLocation2)
}
func convert(converter string, val int) int {
	infos := maps[converter]

	for _, i := range infos {
		if val >= i.sourceStart && val <= i.sourceStart+i.rangeLength {
			offset := val - i.sourceStart
			return i.destinationStart + offset
		}
	}
	return val
}

func readMaps(lines []string) {
	var batches [][]string
	var batch []string

	for _, line := range lines {
		if line == "" {
			batches = append(batches, batch)
			batch = []string{}
			continue
		}
		batch = append(batch, line)
	}

	for _, b := range batches {
		mapName := strings.Fields(b[0])[0]

		converters = append(converters, mapName)

		for _, i := range b[1:] {
			values := strings.Fields(i)

			destinationStart, _ := strconv.Atoi(values[0])
			sourceStart, _ := strconv.Atoi(values[1])
			rangeLength, _ := strconv.Atoi(values[2])
			maps[mapName] = append(maps[mapName], info{destinationStart, sourceStart, rangeLength})
		}
	}

	println()
}

func getSeeds(line string) []int {
	seeds := strings.Split(line, " ")[1:]

	var seedsInt []int
	for i := 0; i < len(seeds); i++ {
		val, _ := strconv.Atoi(seeds[i])
		seedsInt = append(seedsInt, val)
	}

	return seedsInt
}

func getSeedPairs(line string) [][]int {
	fields := strings.Fields(line)
	var pairs [][]int
	for i := 1; i < len(fields[1:]); i += 2 {
		seed, _ := strconv.Atoi(fields[i])
		seedRange, _ := strconv.Atoi(fields[i+1])
		pairs = append(pairs, []int{seed, seedRange})
	}

	return pairs
}

func getInputLines() []string {
	file, err := os.ReadFile("2023/day5/input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(file), "\n")

	return lines
}
