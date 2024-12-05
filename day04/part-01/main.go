package main

import (
	"bufio"
	"fmt"
	"os"
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

	scanner := bufio.NewScanner(file)

	var ws [][]string

	for scanner.Scan() {
		line := scanner.Text()
		ws = append(ws, strings.Split(line, ""))
	}

	var t int

	rows := len(ws)
	columns := len(ws[0])

	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			// right
			if j+3 < columns &&
				ws[i][j] == "X" && ws[i][j+1] == "M" && ws[i][j+2] == "A" && ws[i][j+3] == "S" {
				t++
			}

			// left
			if j-3 >= 0 &&
				ws[i][j] == "X" && ws[i][j-1] == "M" && ws[i][j-2] == "A" && ws[i][j-3] == "S" {
				t++
			}

			// down
			if i+3 < rows &&
				ws[i][j] == "X" && ws[i+1][j] == "M" && ws[i+2][j] == "A" && ws[i+3][j] == "S" {
				t++
			}

			// up
			if i-3 >= 0 &&
				ws[i][j] == "X" && ws[i-1][j] == "M" && ws[i-2][j] == "A" && ws[i-3][j] == "S" {
				t++
			}

			// Diagonal down \
			if i+3 < rows && j+3 < columns &&
				ws[i][j] == "X" && ws[i+1][j+1] == "M" && ws[i+2][j+2] == "A" && ws[i+3][j+3] == "S" {
				t++
			}

			// Diagonal down /
			if i+3 < rows && j-3 >= 0 &&
				ws[i][j] == "X" && ws[i+1][j-1] == "M" && ws[i+2][j-2] == "A" && ws[i+3][j-3] == "S" {
				t++
			}

			// Diagonal up /
			if i-3 >= 0 && j+3 < columns &&
				ws[i][j] == "X" && ws[i-1][j+1] == "M" && ws[i-2][j+2] == "A" && ws[i-3][j+3] == "S" {
				t++
			}

			// Diagonal up \
			if i-3 >= 0 && j-3 >= 0 &&
				ws[i][j] == "X" && ws[i-1][j-1] == "M" && ws[i-2][j-2] == "A" && ws[i-3][j-3] == "S" {
				t++
			}

		}
	}

	fmt.Println(t)

}
