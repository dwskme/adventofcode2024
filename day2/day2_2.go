package Day2

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func withDampener(values []int) bool {
	if isSafe(values) {
		return true
	}

	for i := 0; i < len(values); i++ {
		newValues := append([]int{}, values[:i]...)
		newValues = append(newValues, values[i+1:]...)
		if isSafe(newValues) {
			return true
		}
	}

	return false
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Day2_2() int {
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

		if withDampener(values) {
			count++
		}
	}

	return count
}
