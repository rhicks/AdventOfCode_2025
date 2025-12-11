package main

import (
	"fmt"
	"strings"
)

func day07_1(input string) {
	lines := strings.Split(input, "\n")
	totalLines := len(lines)
	lineLength := len(lines[totalLines-1])
	splitCount := 0

	println(totalLines, lineLength) // output: 16 15

	// make an empty grid of the correct size
	grid := make([][]string, totalLines)
	for i := range totalLines {
		grid[i] = make([]string, lineLength)
	}
	// PrintGrid7(grid) // output: empty grid of 16x15

	// loop over the input and populate the grid
	for i, line := range lines {
		for idx, char := range line {
			switch char {
			case '.':
				grid[i][idx] = "."
			case 'S':
				grid[i][idx] = "S"
			case '^':
				grid[i][idx] = "^"
			}
		}
	}
	PrintGrid7(grid) //print full grid before splitting
	fmt.Println()

	// loop over the grid and add splits
	for r, row := range grid {
		for c, value := range row {
			counted := false
			if value == "S" {
				grid[r+1][c] = "|"
			}
			if value == "." && r > 0 && grid[r-1][c] == "|" {
				grid[r][c] = "|"
			}
			if value == "^" && grid[r][c-1] != "|" {
				grid[r][c-1] = "|"
				// splitCount++
				counted = true
			}
			if value == "^" && grid[r][c+1] != "|" {
				grid[r][c+1] = "|"
				if !counted {
					// splitCount++
				}
			}
		}
	}

	// loop over the grid and count splits
	for r, row := range grid {
		for c, value := range row {
			if value == "^" && grid[r-1][c] == "|" {
				splitCount++
			}
		}
	}
	PrintGrid7(grid) //print full grid after splitting
	fmt.Println()
	fmt.Println("Grand Total:", splitCount)
}

func PrintGrid7(grid [][]string) {
	for i, row := range grid {
		fmt.Printf("%02d %v\n", i, row)
	}
}

func UpdateGride(grid [][]int) {
	// do nothing yet
}
