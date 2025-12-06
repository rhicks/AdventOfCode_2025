package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/samyakbardiya/advent-of-code/2025/util"
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

	data := []string{
		"987654321111111",
		"811111111111119",
		"234234234234278",
		"818181911112111",
	}

	log.SetLevel(log.InfoLevel)

	totalJoltage := 0

	for _, bank := range data {
		joltage := ""
		startIdx := 0
		needed := 12
		fmt.Println(bank)

		for i := range needed {
			maxIdx := startIdx
			remaining := needed - i - 1
			maxSearchIdx := len(bank) - remaining
			fmt.Println(maxIdx, remaining, maxSearchIdx)

			for j := startIdx; j < maxSearchIdx; j++ {
				fmt.Println(bank[j], bank[maxIdx])
				if bank[j] > bank[maxIdx] {
					maxIdx = j
				}
			}

			joltage += string(bank[maxIdx])
			startIdx = maxIdx + 1
		}
		log.Debug("", "joltage", joltage)

		totalJoltage += util.Atoi(joltage)
	}

	fmt.Println(totalJoltage)
}
