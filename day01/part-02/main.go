package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func CountSimilars(arr []int, value int) int {
	total := 0
	for _, v := range arr {
		if v == value {
			total++
		}
	}
	return total
}

func main() {
	filePath := "./input.txt"

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Failed to open file:", err)
		return
	}
	defer file.Close()

	var c1 []int
	var c2 []int

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		v := strings.Fields(line)

		n1, _ := strconv.Atoi(v[0])
		n2, _ := strconv.Atoi(v[1])

		c1 = append(c1, n1)
		c2 = append(c2, n2)

	}

	sort.Ints(c1)
	sort.Ints(c2)

	r := 0

	for _, v := range c1 {
		s := CountSimilars(c2, v)
		r += s * v
	}

	println(r)
}
