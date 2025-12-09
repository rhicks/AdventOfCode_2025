package main

import (
	"fmt"
	"strings"
)

func day04_1(input string) string {
	//data := []string{
	//	"..@@.@@@@.\n@@@.@.@.@@\n@@@@@.@.@@\n@.@@@@..@.\n@@.@@@@.@@\n.@@@@@@@.@\n.@.@.@.@@@\n@.@@@.@@@@\n.@@@@@@@@.\n@.@.@@@.@.",
	//}

	// The forklifts can only access a roll of paper if there are fewer than four
	// rolls of paper in the eight adjacent positions. If you can figure out which
	// rolls of paper the forklifts can access, they'll spend less time looking and
	// more time breaking down the wall to the cafeteria.
	//
	//In this example, there are 13 rolls of paper that can be accessed by a forklift (marked with x):

	// ..@@.@@@@.  ..xx.xx@x.
	// @@@.@.@.@@  x@@.@.@.@@
	// @@@@@.@.@@  @@@@@.x.@@
	// @.@@@@..@.  @.@@@@..@.
	// @@.@@@@.@@  x@.@@@@.@x
	// .@@@@@@@.@  .@@@@@@@.@
	// .@.@.@.@@@  .@.@.@.@@@
	// @.@@@.@@@@  x.@@@.@@@@
	// .@@@@@@@@.  .@@@@@@@@.
	// @.@.@@@.@.  x.x.@@@.x.

	fmt.Println("Day 04_1")
	fmt.Println(input)
	lines := strings.SplitSeq(strings.TrimSpace(input), "\n")
	// println(data[0])
	totalLines := 0
	lineLength := 0
	for line := range lines {
		// println(len(line)) output: 10
		totalLines++
		lineLength = len(line)
	}
	// println(totalLines, lineLength) output: 10 10
	// create grid based on line width, and number lines
	// populate grid
	// count adjacencies
	// Create a 10x10 grid of ints
	grid := make([][]int, totalLines)
	for i := 0; i < lineLength; i++ {
		grid[i] = make([]int, lineLength)
	}
	PrintGrid(grid) // should be empty grid
	println("---")
	lines2 := strings.Split(input, "\n")
	for i, line := range lines2 {
		for idx, char := range line {
			if char == 46 {
				grid[i][idx] = 0
			} else {
				grid[i][idx] = 1
			}

		}
	}
	PrintGrid(grid) // should be full grid now
	// total := calcGrid(grid, totalLines, lineLength, grid)
	//println(total)
	return input
}

func day04_2(input string) string {
	fmt.Println("Day 04_2")
	fmt.Println(input)
	lines := strings.SplitSeq(strings.TrimSpace(input), "\n")
	totalLines := 0
	lineLength := 0
	for line := range lines {
		// println(len(line)) output: 10
		totalLines++
		lineLength = len(line)
	}
	// println(totalLines, lineLength) output: 10 10
	grid := make([][]int, totalLines)
	grid2 := make([][]int, totalLines)
	for i := 0; i < lineLength; i++ {
		grid[i] = make([]int, lineLength)
		grid2[i] = make([]int, lineLength)
	}
	PrintGrid(grid) // should be empty grid
	println("---")
	lines2 := strings.Split(input, "\n")
	for i, line := range lines2 {
		for idx, char := range line {
			if char == 46 {
				grid[i][idx] = 0
				grid2[i][idx] = 0
			} else {
				grid[i][idx] = 1
				grid2[i][idx] = 1
			}

		}
	}
	PrintGrid(grid2) // should be full grid now
	total := calcGrid(grid, totalLines, lineLength, grid2)
	grandTotal := total
	println(total)
	for total > 0 {
		grid = grid2
		total = calcGrid(grid, totalLines, lineLength, grid2)
		grandTotal += total
	}
	PrintGrid(grid2)
	println(grandTotal)
	return input
}

func PrintGrid(grid [][]int) {
	for _, row := range grid {
		fmt.Println(row)
	}
}

func isThisTP(grid [][]int, r int, c int) bool {
	if grid[r][c] == 1 {
		return true
	} else {
		return false
	}
}

func accessible(grid [][]int, r int, c int, maxRows int, maxCols int) bool {
	total := 0
	topRow := false
	bottomRow := false
	firstCol := false
	lastCol := false
	if r-1 < 0 {
		topRow = true
	}
	if r+1 > maxRows-1 {
		bottomRow = true
	}
	if c-1 < 0 {
		firstCol = true
	}
	if c+1 > maxCols-1 {
		lastCol = true
	}
	if !topRow && !firstCol && grid[r-1][c-1] == 1 {
		fmt.Printf("+1:grid[%d][%d]\n", r, c)
		total++
	}
	if !topRow && grid[r-1][c] == 1 {
		fmt.Printf("+1:grid[%d][%d]\n", r, c)
		total++
	}
	if !topRow && !lastCol && grid[r-1][c+1] == 1 {
		fmt.Printf("+1:grid[%d][%d]\n", r, c)
		total++
	}
	if !firstCol && grid[r][c-1] == 1 {
		fmt.Printf("+1:grid[%d][%d]\n", r, c)
		total++
	}
	if !lastCol && grid[r][c+1] == 1 {
		fmt.Printf("+1:grid[%d][%d]\n", r, c)
		total++
	}
	if !bottomRow && !firstCol && grid[r+1][c-1] == 1 {
		fmt.Printf("+1:grid[%d][%d]\n", r, c)
		total++
	}
	if !bottomRow && grid[r+1][c] == 1 {
		fmt.Printf("+1:grid[%d][%d]\n", r, c)
		total++
	}
	if !bottomRow && !lastCol && grid[r+1][c+1] == 1 {
		fmt.Printf("+1:grid[%d][%d]\n", r, c)
		total++
	}
	println("total from acc:", total)
	if total > 3 {
		return false
	} else {
		fmt.Printf("ACCESSIBLE:grid[%d][%d]\n", r, c)
		return true
	}
}

func calcGrid(grid [][]int, maxRows int, maxCols int, grid2 [][]int) int {
	total := 0
	for r, row := range grid {
		// println(maxCols) // output: 10
		// println(maxRows) // output: 10
		for c, value := range row {
			// total = 0
			if isThisTP(grid, r, c) {
				fmt.Printf("TP AT: grid[%d][%d] = %d\n", r, c, value)
				if accessible(grid, r, c, maxRows, maxCols) {
					grid2[r][c] = 0
					total++
				}
			}
		}
	}
	return total
}
