package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

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

	t := 0

	for i := 0; i < len(c1); i++ {
		d := c1[i] - c2[i]

		if d < 0 {
			t += -d
		} else {
			t += d
		}
	}

	fmt.Println(t)
}
