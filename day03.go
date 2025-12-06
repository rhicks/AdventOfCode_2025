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

	// find largest number in the string
	// mark the index of that number
	// from the index forward, find the next highest number

	lines := strings.SplitSeq(strings.TrimSpace(input), "\n")
	total := 0
	for line := range lines {
		// fmt.Println(line)
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

func day03_2(input string) {
	fmt.Println("Day 3_2")

	// 987654321111111 - Org
	// 987654321111    - Good
	// 811111111111119
	// 811111111119
	// 234234234234278
	// 434234234278
	// 818181911112111
	// 888911112111

	maxSizeofLine := 12

	lines := strings.SplitSeq(strings.TrimSpace(input), "\n")
	for line := range lines {
		battery := 0
		println()
		println(line)
		for i := 0; i < len(line); i++ {
			shortString := getStringOfSize(line, maxSizeofLine+1) // get a 13 char line
			biggestNumber := getLargestString(shortString)        // find largest 12 digit number
			println("Comparing:", battery, biggestNumber)
			if biggestNumber > battery {
				battery = biggestNumber
			}
			// battery = biggestNumber // store largest 12 digit number in library
			line = shiftRightDropFirst(line)
		}
		println("battery:", battery)
	}
}

func getStringOfSize(s string, size int) string {
	if size > len(s) {
		return s
	} else {
		return s[:size]
	}
}

func getLargestString(s string) int {
	if len(s) != 13 {
		// handle error however you like; for now, just panic
		panic("input must be exactly 13 characters long")
	}
	maxStr := ""
	for i := 0; i < len(s); i++ {
		// remove the character at index i
		candidate := s[:i] + s[i+1:]

		// first candidate, or larger than current max?
		if candidate > maxStr {
			maxStr = candidate
		}
	}
	// convert the largest 12-digit string to an int
	result, err := strconv.Atoi(maxStr)
	if err != nil {
		panic(err) // or handle error
	}

	return result
}

func removeAtIndex(s string, index int) string {
	return s[:index] + s[index+1:]
}

func shiftRightDropFirst(s string) string {
	if len(s) <= 13 {
		return s
	}
	return s[1:]
}
