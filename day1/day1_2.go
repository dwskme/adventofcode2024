package Day1

import (
	"os"
	"strconv"
	"strings"
)

func sumInsideSlice(v []int) int {
	sum := 0
	for _, value := range v {
		sum += value
	}
	return sum
}

func Day1_2() int {
	var left, right, result []int

	data, _ := os.ReadFile("day1/input.txt")
	lines := strings.Split(string(data), "\n")

	for _, line := range lines {
		parts := strings.Fields(line)
		if len(parts) == 2 {
			l, _ := strconv.Atoi(parts[0])
			left = append(left, l)
			r, _ := strconv.Atoi(parts[1])
			right = append(right, r)
		}
	}

	for _, l := range left {
		count := 0
		for _, r := range right {
			if l == r {
				count++
			}
		}
		result = append(result, l*count)
	}

	return sumInsideSlice(result)
}
