package main

import (
	"fmt"
	"strconv"
	"strings"
)

func day06_1(input string) {
	var nums [][]any
	grandTotal := 0

	type Problem struct {
		ID        int
		Numbers   []any
		Operation string
	}

	problemsByID := make(map[int]*Problem)
	var problems []*Problem

	lines := strings.Split(input, "\n")

	for _, line := range lines {
		fields := strings.Fields(line)
		for i, field := range fields {
			nums = append(nums, []any{i, field})
		}
		// fmt.Println(i)
		// create an array of len(lines)
		// only do that on the first run
		// fmt.Println(nums)
	}
	for _, row := range nums {
		id := row[0].(int)
		numOrOp := row[1]

		prob, exists := problemsByID[id]
		if !exists {
			// Create new struct if ID not seen before
			prob = &Problem{
				ID:        id,
				Numbers:   []any{},
				Operation: "",
			}
			problemsByID[id] = prob
			problems = append(problems, prob)
		}
		// Update existing struct (or the one we just created)
		prob.Numbers = append(prob.Numbers, numOrOp)
	}
	for _, r := range problems {
		// fmt.Printf("ID %d: %v\n", r.ID, r.Numbers)
		r.Operation = r.Numbers[len(r.Numbers)-1].(string)
		r.Numbers = r.Numbers[:len(r.Numbers)-1]
	}
	for _, r := range problems {
		fmt.Printf("ID %d: %v %s\n", r.ID, r.Numbers, r.Operation)
		total := 0
		if r.Operation == "*" {
			total = 1
			for _, num := range r.Numbers {
				temp, _ := strconv.Atoi(num.(string))
				total = total * temp
			}
		}
		if r.Operation == "+" {
			total = 0
			for _, num := range r.Numbers {
				temp, _ := strconv.Atoi(num.(string))
				total = total + temp
			}
		}
		fmt.Println(total)
		grandTotal += total
	}
	fmt.Println(grandTotal)

}

// func day06_2(input string) {
// 	lines := strings.Split(input, "\n")

// }
