package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var left, right []int
	sum := 0

	data, _ := os.ReadFile("input.txt")
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
		sum += abs(left[i] - right[i])
	}

	fmt.Println(sum)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
