package main

import (
	"fmt"
	"strconv"
	"strings"
)

func day02_1(input string) {
	fmt.Println("Day 2_1")
	total := 0

	sequences := strings.SplitSeq(strings.TrimSpace(input), ",")
	for sequence := range sequences {
		startNumber, _ := strconv.Atoi(strings.Split(sequence, "-")[0])
		endNumber, _ := strconv.Atoi(strings.Split(sequence, "-")[1])

		for startNumber < (endNumber + 1) {
			numberString := strconv.Itoa(startNumber)
			numberLength := len(numberString)
			if numberLength%2 == 0 {
				firstHalf := numberString[:numberLength/2]
				secondHalf := numberString[numberLength/2:]

				if firstHalf == secondHalf {
					total += startNumber
				}
			}
			startNumber++
		}

	}
	fmt.Println(total)
}

func day02_2(input string) {
	fmt.Println()
	fmt.Println("Day 2_2")
	total := 0

	sequences := strings.SplitSeq(strings.TrimSpace(input), ",")
	for sequence := range sequences {
		startNumber, _ := strconv.Atoi(strings.Split(sequence, "-")[0])
		endNumber, _ := strconv.Atoi(strings.Split(sequence, "-")[1])

		for startNumber < (endNumber + 1) {
			startNumberString := strconv.Itoa(startNumber)
			divisors := divisors(len(startNumberString))
			for i, value := range divisors {
				if i != 0 {
					parts, _ := splitIntoParts(startNumberString, value)
					if allEqual(parts) {
						// fmt.Println(parts, startNumberString, value)
						total = total + startNumber
						break
					}
				}
			}
			startNumber++
		}
	}
	fmt.Println(total)
}

func divisors(n int) []int {
	result := []int{}
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			result = append(result, i)
		}
	}
	return result
}

func splitIntoParts(s string, parts int) ([]string, error) {
	if len(s)%parts != 0 {
		return nil, fmt.Errorf("cannot split: string length %d not evenly divisible by %d", len(s), parts)
	}

	size := len(s) / parts
	result := make([]string, parts)

	for i := 0; i < parts; i++ {
		start := i * size
		end := start + size
		result[i] = s[start:end]
	}

	return result, nil
}

func allEqual(values []string) bool {
	for _, v := range values {
		if v != values[0] {
			return false
		}
	}
	return true
}
