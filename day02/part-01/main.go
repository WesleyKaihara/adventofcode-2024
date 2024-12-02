package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isSafe(levels []int) bool {
	increasing := levels[1] > levels[0]
	for i := 1; i < len(levels); i++ {
		diff := levels[i] - levels[i-1]

		if diff < -3 || diff > 3 || diff == 0 {
			return false
		}

		if increasing && diff < 0 {
			return false
		} else if !increasing && diff > 0 {
			return false
		}
	}

	return true
}

func main() {
	filePath := "./input.txt"

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Failed to open file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	t := 0

	for scanner.Scan() {
		line := scanner.Text()
		v := strings.Fields(line)

		levels := make([]int, len(v))
		for i, level := range v {
			levels[i], _ = strconv.Atoi(level)
		}

		if isSafe(levels) {
			t++
		}
	}

	fmt.Println(t)
}
