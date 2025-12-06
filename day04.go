package main

import (
	"fmt"
	"strings"
)

func day04_1(input string) {
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
	// fmt.Println(input)
	lines := strings.SplitSeq(strings.TrimSpace(input), "\n")
	// println(data[0])
	totalLines := 0
	lineLength := 0
	for line := range lines {
		println(len(line))
		totalLines++
		lineLength = len(line)
	}
	println(totalLines, lineLength)
	// create grid based on line width, and number lines
	// populate grid
	// count adjacencies
	// Create a 10x10 grid of ints
	grid := make([][]int, totalLines)

	for i := 0; i < lineLength; i++ {
		grid[i] = make([]int, lineLength)
	}
	grid[3][4] = 7

	for _, row := range grid {
		fmt.Println(row)
	}
}
