package main

import (
	"fmt"
	"strconv"
	"strings"
)

func day05_1(input string) {
	lines := strings.Split(input, "\n")
	freshItems := map[int]struct{}{} // set of ints
	startOfIng := 0                  // start of Ingredients
	total := 0
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

			if startNumber > endNumber {
				fmt.Println(startNumber, endNumber)
			}

			for startNumber <= endNumber {
				// fmt.Println(startNumber)
				freshItems[startNumber] = struct{}{}
				startNumber++
			}
		}
	}
	for i := startOfIng; i < len(lines); i++ {
		number, _ := strconv.Atoi(lines[i])
		_, exists := freshItems[number]
		if exists {
			// fmt.Println(number)
			total++
		}
	}
	fmt.Println("Fresh:", total)
}
