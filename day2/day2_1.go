package Day2

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func isSafe(values []int) bool {
	incr, decr := false, false

	for i := 1; i < len(values); i++ {
		left := values[i-1]
		right := values[i]

		if left < right {
			incr = true
		} else if left > right {
			decr = true
		}

		if incr && decr {
			return false
		}

		if abs(left-right) < 1 || abs(left-right) > 3 {
			return false
		}
	}
	return true
}

func Day2_1() int {
	file, _ := os.Open("day2/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	count := 0

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)

		values := make([]int, len(parts))
		for i, part := range parts {
			values[i], _ = strconv.Atoi(part)
		}

		if isSafe(values) {
			count++
		}
	}

	return count
}
