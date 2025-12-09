package main

import (
	"fmt"
	"strconv"
	"strings"
)

func day05_1(input string) {
	lines := strings.Split(input, "\n")
	// freshItems := map[int]struct{}{} // set of ints
	startOfIng := 0 // start of Ingredients
	total := 0

	type FreshRange struct {
		Start int
		End   int
	}

	var listOfFreshRanges []FreshRange

	// one := 436639876131680
	// two := 436867778393908
	// fmt.Printf("%T", one-two)
	for i, line := range lines {
		if line == "" {
			println("Ingredients begin at line:", i+1)
			startOfIng = i + 1
			break
		}
		if line != "" {
			parts := strings.Split(line, "-")
			startNumber, _ := strconv.Atoi(parts[0])
			endNumber, _ := strconv.Atoi(parts[1])

			// create a list of fresh ranges
			yar := FreshRange{startNumber, endNumber}
			listOfFreshRanges = append(listOfFreshRanges, yar)

			fmt.Println(yar)
			// if startNumber > endNumber {
			// 	fmt.Println(startNumber, endNumber)
			// }

			// for startNumber <= endNumber {
			// 	// fmt.Println(startNumber)
			// 	freshItems[startNumber] = struct{}{}
			// 	startNumber++
			// }
		}
	}
	for i := startOfIng; i < len(lines); i++ {
		number, _ := strconv.Atoi(lines[i])
		// _, exists := freshItems[number]
		// if exists {
		// 	// fmt.Println(number)
		// 	total++
		// }
		for _, r := range listOfFreshRanges {
			if number >= r.Start && number <= r.End {
				total++
				break
			}
		}
	}
	fmt.Println("Fresh:", total)
}
