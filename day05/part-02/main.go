package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isValid(update []int, rules []string) bool {
	position := make(map[int]int)
	for i, page := range update {
		position[page] = i
	}

	for _, rule := range rules {
		pageRole := strings.Split(rule, "|")
		n1, _ := strconv.Atoi(pageRole[0])
		n2, _ := strconv.Atoi(pageRole[1])

		p1, x := position[n1]
		if !x {
			continue
		}

		p2, y := position[n2]
		if !y {
			continue
		}

		if p1 > p2 {
			return false
		}
	}

	return true
}

func findRuleInUpdate(update []int, rulePage int) int {
	for i, page := range update {
		if page == rulePage {
			return i
		}
	}
	return -1
}

func fixOrder(update []int, rules []string) []int {
	for {
		updated := false

		for _, rule := range rules {
			parts := strings.Split(rule, "|")
			n1, _ := strconv.Atoi(parts[0])
			n2, _ := strconv.Atoi(parts[1])

			i1, i2 := findRuleInUpdate(update, n1), findRuleInUpdate(update, n2)

			if i1 != -1 && i2 != -1 && i1 > i2 {
				update[i1], update[i2] = update[i2], update[i1]
				updated = true
			}
		}

		if !updated {
			break
		}
	}

	return update
}

func main() {
	filePath := "./input.txt"

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Failed to open file:", err)
		return
	}
	defer file.Close()

	var rules []string
	var updates [][]int
	isRule := true

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" {
			isRule = false
			continue
		}

		if isRule {
			rules = append(rules, line)
		} else {
			pages := strings.Split(line, ",")

			var updatePages []int

			for _, page := range pages {
				n, _ := strconv.Atoi(page)
				updatePages = append(updatePages, n)
			}

			updates = append(updates, updatePages)
		}
	}

	var t int

	for _, update := range updates {
		if !isValid(update, rules) {
			orderedUpdate := fixOrder(update, rules)

			v := orderedUpdate[len(update)/2]
			t += v
		}
	}

	fmt.Println(t)
}
