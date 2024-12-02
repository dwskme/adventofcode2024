package Day1

import (
	"os"
	"sort"
	"strconv"
	"strings"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Day1_1() int {
	var left, right []int
	Sum := 0

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

	sort.Ints(left)
	sort.Ints(right)

	for i := range left {
		Sum += abs(left[i] - right[i])
	}
	return Sum
}
