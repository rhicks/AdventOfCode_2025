package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func mnml() {
	// Read entire contents of input.txt into memory as a byte slice.
	// Ignoring error for simplicity.
	input, _ := os.ReadFile("inputs/day1.txt")

	// Mapping for direction commands:
	// 'L' subtracts 1 from dial,
	// 'R' adds 1 to dial.
	delta := map[byte]int{'L': -1, 'R': 1}

	// Starting dial value
	dial := 50

	// Counters for two modes of counting revolutions or transitions
	part1, part2 := 0, 0

	// Split input by whitespace — each field should be of form "L10", "R200", etc.
	for _, s := range strings.Fields(string(input)) {

		// Parse the numeric portion after the first character.
		// Example: "L20" -> s[0] = 'L', s[1:] = "20"
		n, _ := strconv.Atoi(s[1:])

		// Loop that runs `n` times.
		// This is a Go idiom: `for range n` loops n times but discards the index.
		// Each iteration moves the dial one unit left or right.
		for range n {

			// Apply the directional change:
			// `dial +=` modifies in place.
			// After incrementing, check if dial hits a multiple of 100.
			dial += delta[s[0]]
			if dial%100 == 0 {
				part2++
			}
		}

		// After all individual steps for this entry are complete,
		// check again if dial is exactly on a 100 boundary.
		if dial%100 == 0 {
			part1++
		}
	}

	// Print results
	fmt.Println(part1)
	fmt.Println(part2)
}
