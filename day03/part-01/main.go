package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	filePath := "./input.txt"

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Failed to open file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var t int

	for scanner.Scan() {
		line := scanner.Text()

		regex := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
		matches := regex.FindAllStringSubmatch(line, -1)

		for _, v := range matches {
			n1, _ := strconv.Atoi(v[1])
			n2, _ := strconv.Atoi(v[2])
			r := n1 * n2
			t += r
		}

	}
	fmt.Println(t)

}
