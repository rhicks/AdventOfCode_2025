package main

import (
	"fmt"
	"sort"
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

type FreshRange struct {
	Start int
	End   int
}

func day05_2(input string) {
	// too big: 682053047137940
	// too ???: 439108493356772
	// too ???: 682690004244977
	// too ???: 865853398114375
	// too ???: 352556672963024
	// too ???: 352556672963116
	lines := strings.Split(input, "\n")
	total := 0

	var listOfFreshRanges []FreshRange

	for i, line := range lines {
		if line == "" {
			println("Ingredients begin at line:", i+1)
			// startOfIng = i + 1
			break
		}
		if line != "" {
			parts := strings.Split(line, "-")
			startNumber, _ := strconv.Atoi(parts[0])
			endNumber, _ := strconv.Atoi(parts[1])

			yar := FreshRange{startNumber, endNumber}
			listOfFreshRanges = append(listOfFreshRanges, yar)
		}
	}
	listOfFreshRanges = mergeFreshRanges(listOfFreshRanges)

	for _, ranges := range listOfFreshRanges {
		fmt.Println(ranges)
		total += ranges.End - ranges.Start + 1
	}
	fmt.Println(total)
}

func mergeFreshRanges(ranges []FreshRange) []FreshRange {
	if len(ranges) == 0 {
		return nil
	}

	// 1) Sort by Start
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].Start < ranges[j].Start
	})

	merged := []FreshRange{ranges[0]}

	for _, r := range ranges[1:] {
		last := &merged[len(merged)-1]

		if r.Start <= last.End {
			if r.End > last.End {
				last.End = r.End
			}
		} else {
			merged = append(merged, r)
		}
	}

	return merged
}
