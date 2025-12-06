package main

import (
	"fmt"
	"slices"
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

	lines := strings.SplitSeq(strings.TrimSpace(input), "\n")
	for line := range lines {
		fmt.Println()
		fmt.Println(line)
		battery := ""
		desiredBatteryLength := 12

		batteryMap := map[int]int{}
		for i, char := range line {
			num := int(char - '0')
			batteryMap[i] = num
		}
		printMapDesc(batteryMap)
		// var newLine []int
		// fmt.Println(len(newLine))
		for i := 9; i >= 0; i-- {
			currentSlice := keysByValue(batteryMap, i)
			if len(currentSlice) > 0 {
				if currentSlice[0] < 88
			}
		}
		// 	getSlice(line, 0, len(line))
		// 	// currentSlice := keysByValue(batteryMap, i)
		// 	// if len(currentSlice) > 0 && len(line) > 12 {
		// 	// }
		// }
		// fmt.Println(i, keysByValue(batteryMap, i))
		// currentSlice := keysByValue(batteryMap, i)
		// if len(currentSlice) > 0 {
		// 	if currentSlice[0] < 88 { // this is suffcient to remove any first numbers above 88
		// 		// fmt.Println(i, keysByValue(batteryMap, i))
		// 		// fmt.Println(currentSlice[0], "Index") // index first number of the BatteryPack
		// 		// fmt.Println(i, "Starting Value")
		// 		newLine := line[currentSlice[0]:] // actul first number of the BatteryPack
		// 		if len(newLine) == 13 {
		// 			fmt.Println(line)
		// 			fmt.Println(newLine)
		// 			break
		// 		}
		// 		//break
		// 	}
		// }
	}
}

func getSlice(line string, start int, end int) {
	// start := 0       //defaults
	// end := len(line) //defaults
	fmt.Println(line[start:end])
}

func keysByValue(m map[int]int, value int) []int {
	var keys []int
	for k, v := range m {
		if value == v {
			keys = append(keys, k)
		}
	}
	slices.Sort(keys)
	return keys
}

func printMapDesc(m map[int]int) {
	// print a map of ints in descending order
	// the map will show the index number of the numbers 0-9
	for i := 9; i >= 0; i-- {
		fmt.Println(i, keysByValue(m, i))
	}
}
