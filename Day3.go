package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	file, err := os.Open("inputs/3.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var treeMap []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		//fmt.Println(scanner.Text())
		treeMap = append(treeMap, scanner.Text())
	}
	width := len(treeMap[0])

	// Solution to part 1, right 3, down 1
	// index is x, y, second step will be x+3, y+1, etc.
	index := [2]int{0, 0}
	treeCount := 0

	for index[1] < len(treeMap) {
		if treeMap[index[1]][index[0]] == '#' {
			treeCount++
		}
		index[0] = (index[0] + 3) % width
		index[1] += 1
	}

	println(treeCount)

	// Solution to part 2, multiple slopes multiplied together
	// Basically same as above, but we'll loop through a slice of slopes
	slopes := [][]int{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}

	treeMultiple := 1
	for _, slope := range slopes {
		index = [2]int{0, 0}
		treeCount = 0
		for index[1] < len(treeMap) {
			if treeMap[index[1]][index[0]] == '#' {
				treeCount++
			}
			index[0] = (index[0] + slope[0]) % width
			index[1] += slope[1]
		}

		treeMultiple *= treeCount

	}

	println(treeMultiple)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
