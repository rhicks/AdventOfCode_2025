package main

import (
	"fmt"
	"strconv"
	"strings"
)

func day03_1(input string) {
	fmt.Println("Day 3_1")

	// In 987654321111111, you can make the largest joltage possible, 98, by turning on the first two batteries.
	// In 811111111111119, you can make the largest joltage possible by turning on the batteries labeled 8 and 9, producing 89 jolts.
	// In 234234234234278, you can make 78 by turning on the last two batteries (marked 7 and 8).
	// In 818181911112111, the largest joltage you can produce is 92.

	lines := strings.SplitSeq(strings.TrimSpace(input), "\n")
	total := 0
	for line := range lines {
		fmt.Println(line)
		largestNumber := 0
		secondLargest := 0
		newIndex := 0
		for i, char := range line {
			num := int(char - '0')
			if num > largestNumber && i != (len(line)-1) {
				largestNumber = num
				newIndex = i
			}
		}
		for i, char := range line {
			num := int(char - '0')
			if num > secondLargest && i > newIndex {
				secondLargest = num
			}
		}
		temp, _ := strconv.Atoi(strconv.Itoa(largestNumber) + strconv.Itoa(secondLargest))
		total = total + temp
	}
	fmt.Println(total)
}

// func findLargestNumber(startIndex int, line string) int {
// 	return 0
// }

// find largest number in the string
// mark the index of that number
// from the index forward, find the next highest number
