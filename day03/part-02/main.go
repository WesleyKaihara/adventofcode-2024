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
	enabled := true

	for scanner.Scan() {
		line := scanner.Text()

		regex := regexp.MustCompile(`do\(\)|don't\(\)|mul\((\d+),(\d+)\)`)
		matches := regex.FindAllStringSubmatch(line, -1)

		for _, v := range matches {
			if v[0] == "do()" {
				enabled = true
			} else if v[0] == "don't()" {
				enabled = false
				continue
			}

			if enabled && v[1] != "" && v[2] != "" {
				n1, _ := strconv.Atoi(v[1])
				n2, _ := strconv.Atoi(v[2])
				r := n1 * n2
				t += r
			}

		}

	}
	fmt.Println(t)

}
