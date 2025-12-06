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

	// We want exactly 12 digits in the final number for each line.
	const digitsToKeep = 12

	var total int64

	// Split the whole input into separate lines (each line is one bank of batteries).
	lines := strings.SplitSeq(strings.TrimSpace(input), "\n")

	// Go through each line one by one.
	for line := range lines {
		// Skip empty lines just in case.
		if line == "" {
			continue
		}

		// For this line, build the largest possible number
		// by turning on exactly `digitsToKeep` batteries.
		best := buildLargestNumber(line, digitsToKeep)

		// Turn the digit string into a number so we can add it to the total.
		val, err := strconv.ParseInt(best, 10, 64)
		if err != nil {
			panic(err)
		}

		total += val
		println()
	}

	// Print the total joltage from all lines.
	fmt.Println(total)
}

// buildLargestNumber takes a string of digits (line) and a number of digits to keep
// (digitsToKeep). It returns the largest possible number (as a string) you can make
// by removing some digits but keeping the remaining digits in the same order.
func buildLargestNumber(line string, digitsToKeep int) string {
	// How many digits we are allowed to remove.
	toRemove := len(line) - digitsToKeep

	// This slice is our "stack". We will push digits onto it and sometimes pop from it.
	stack := make([]byte, 0, len(line))

	// Look at each digit from left to right.
	for i := 0; i < len(line); i++ {
		currentDigit := line[i]

		// While:
		// 1. We are still allowed to remove digits (toRemove > 0).
		// 2. The stack is not empty.
		// 3. The last digit we kept is smaller than the current digit.
		// Then: removing the smaller last digit makes the number bigger,
		// so we pop it from the stack.
		for toRemove > 0 && len(stack) > 0 {
			lastDigit := stack[len(stack)-1]

			if lastDigit < currentDigit {
				// Pop the last digit off the stack.
				stack = stack[:len(stack)-1]
				toRemove--
			} else {
				// The last digit is not smaller, so stop removing here.
				break
			}
		}

		// Add (push) the current digit onto the stack.
		stack = append(stack, currentDigit)
		fmt.Println("stack now:", string(stack))
		fmt.Println("toRemove:", toRemove)
	}

	// If we still need to remove digits after processing all characters,
	// remove them from the end of the stack. These are the least valuable
	// digits to keep.
	if toRemove > 0 {
		stack = stack[:len(stack)-toRemove]
	}

	// At this point, stack should have at least digitsToKeep digits.
	// If it has more, just keep the first digitsToKeep.
	if len(stack) > digitsToKeep {
		stack = stack[:digitsToKeep]
	}

	// Turn the stack of bytes back into a string of digits.
	return string(stack)
}
