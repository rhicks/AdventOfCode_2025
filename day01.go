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
			number, _ = strconv.Atoi(line[1:])
			for number > 100 {
				number = number - 100
			}
			currentPosition, zeroHit = getDelta(currentPosition, number, zeroHit)
		} else if strings.Contains(line, "L") {
			number, _ = strconv.Atoi(line[1:])
			for number > 100 {
				number = number - 100
			}
			number = number * -1
			currentPosition, zeroHit = getDelta(currentPosition, number, zeroHit)
		} else {
			fmt.Println("ERROR")
		}
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
			number, _ = strconv.Atoi(line[1:])
			currentPosition, zeroHit = spinDial("right", number, currentPosition, zeroHit)
		} else if strings.Contains(line, "L") {
			number, _ = strconv.Atoi(line[1:])
			currentPosition, zeroHit = spinDial("left", number, currentPosition, zeroHit)
		} else {
			fmt.Println("ERROR")
		}
	}
	fmt.Println("Day 01 02:", zeroHit)
}

func spinDial(dir string, count int, position int, hitCounter int) (int, int) {
	hit := 0
	full := 0
	start := 0

	if dir == "right" {
		hit = 0
		full = 100
		start = 0
	}
	if dir == "left" {
		hit = 100
		full = 0
		start = 100
	}

	if position == hit {
		hitCounter++
	}
	for count > 0 {
		if position == full {
			position = start
			hitCounter++
		}
		if dir == "right" {
			position++
		}
		if dir == "left" {
			position--
		}
		count--
	}
	return position, hitCounter
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
