package main

import (
	"bufio"
	"fmt"
	"os"
)

func wasVisited(visited [][]int, x int, y int) bool {
	for _, item := range visited {
		if item[0] == x && item[1] == y {
			return true
		}
	}
	return false
}

func main() {
	filePath := "./input.txt"

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Failed to open file:", err)
		return
	}
	defer file.Close()

	var mp [][]int32
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		mp = append(mp, []int32(line))
	}

	var x, y, dir int
	for i := 0; i < len(mp); i++ {
		for j := 0; j < len(mp[i]); j++ {
			if mp[i][j] == '^' {
				x, y, dir = j, i, 0
				break
			}
		}
	}

	visited := [][]int{{x, y}}

	directions := [][]int{
		{0, -1}, // up
		{1, 0},  // right
		{0, 1},  // down
		{-1, 0}, // left
	}

	for {
		nx := x + directions[dir][0]
		ny := y + directions[dir][1]

		if ny < 0 || ny >= len(mp) || nx < 0 || nx >= len(mp[ny]) {
			break
		}

		if mp[ny][nx] == '#' {
			dir++
			if dir == 4 {
				dir = 0
			}
		} else {
			x = nx
			y = ny

			if !wasVisited(visited, nx, ny) {
				visited = append(visited, []int{x, y})
			}
		}
	}

	fmt.Println(len(visited))
}
