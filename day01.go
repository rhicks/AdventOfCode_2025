package main

import (
	"fmt"
	"strconv"
	"strings"
)

func day01_1(input string) {
	currentPosition := 50
	zeroHit := 0
	number := 0
	lines := strings.SplitSeq(strings.TrimSpace(input), "\n")

	for line := range lines {
		if strings.Contains(line, "R") {
			// fmt.Println("ADD")
			number, _ = strconv.Atoi(line[1:])
			for number > 100 {
				number = number - 100
			}
			currentPosition, zeroHit = getDelta(currentPosition, number, zeroHit)
		} else if strings.Contains(line, "L") {
			// .Println("SUBTRACT")
			number, _ = strconv.Atoi(line[1:])
			for number > 100 {
				number = number - 100
			}
			number = number * -1
			currentPosition, zeroHit = getDelta(currentPosition, number, zeroHit)
		} else {
			fmt.Println("ERROR")
		}
		//fmt.Println(currentPosition)
	}
	fmt.Println("Day 01 01:", zeroHit)
}

func day01_2(input string) {
	currentPosition := 50
	zeroHit := 0
	number := 0
	lines := strings.SplitSeq(strings.TrimSpace(input), "\n")

	for line := range lines {
		if strings.Contains(line, "R") {
			// fmt.Println("ADD")
			number, _ = strconv.Atoi(line[1:])
			for number > 100 {
				number = number - 100
			}
			currentPosition, zeroHit = getDelta(currentPosition, number, zeroHit)
		} else if strings.Contains(line, "L") {
			// .Println("SUBTRACT")
			number, _ = strconv.Atoi(line[1:])
			for number > 100 {
				number = number - 100
			}
			number = number * -1
			currentPosition, zeroHit = getDelta(currentPosition, number, zeroHit)
		} else {
			fmt.Println("ERROR")
		}
		// fmt.Println(currentPosition)
	}
	fmt.Println(zeroHit)
}

func getDelta(current int, new int, hitCounter int) (int, int) {
	delta := current + new

	if delta == 100 || delta == 0 {
		hitCounter++
		return 0, hitCounter
	}
	if delta < 0 {
		delta = 100 + delta
	}
	if delta > 100 {
		delta = delta - 100
	}

	return delta, hitCounter

}
