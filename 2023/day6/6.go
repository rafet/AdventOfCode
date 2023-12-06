package main

func main() {
	times := []int{49, 97, 94, 94}
	distances := []int{263, 1532, 1378, 1851}
	multi := 1
	for i := 0; i < len(times); i++ {
		multi *= validCount(times[i], distances[i])
	}

	println("answer part1:", multi)

	// Part 2

	t := 49979494
	d := 263153213781851

	println("answer part2:", validCount(t, d))
}

func validCount(t, d int) int {
	cnt := 0
	for a := 0; a < t; a++ {
		if -a*a+a*t > d {
			cnt += 1
		}
	}

	return cnt
}
