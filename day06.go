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

func day06_2(input string) {
	lines := strings.Split(strings.TrimRight(input, "\n"), "\n")
	if len(lines) == 0 {
		return
	}

	lastLine := lines[len(lines)-1]
	var widthList []int

	// 1) Find the *end index* of each column by looking at the operator line
	for i, char := range lastLine {
		if char != ' ' && i != 0 {
			widthList = append(widthList, i)
		}
	}
	widthList = append(widthList, len(lastLine))

	type Problem struct {
		ID        int
		Numbers   []string
		Operation string
		Width     int
	}

	// 2) Pre-create one Problem per column
	problems := make([]*Problem, len(widthList))
	prev := 0
	for idx, end := range widthList {
		width := end - prev
		problems[idx] = &Problem{
			ID:        idx,
			Numbers:   []string{},
			Operation: "",
			Width:     width,
		}
		prev = end
	}

	// 3) For all lines *except* the last, treat them as numbers
	for _, line := range lines[:len(lines)-1] {
		nums := SplitByFieldWidths(line, widthList)
		for i, field := range nums {
			problems[i].Numbers = append(problems[i].Numbers, field)
		}
	}

	// 4) Last line contains operators
	opFields := SplitByFieldWidths(lastLine, widthList)
	for i, field := range opFields {
		op := ""
		for _, r := range field {
			if r != ' ' {
				op = string(r)
				break
			}
		}
		problems[i].Operation = op
	}

	// Helper: for a given Problem, split its Numbers into vertical digit columns.
	splitProblemIntoColumns := func(p *Problem) [][]int {
		var cols [][]int

		// If there are no rows or width is zero, nothing to do.
		if len(p.Numbers) == 0 || p.Width == 0 {
			return cols
		}

		// Iterate over each character position in the field width.
		for col := 0; col < p.Width; col++ {
			var digits []int

			// Walk down the rows and collect digits at this column position.
			for _, row := range p.Numbers {
				if col >= len(row) {
					continue
				}
				ch := row[col]
				if ch >= '0' && ch <= '9' {
					digits = append(digits, int(ch-'0'))
				}
			}

			// Only keep columns that actually contain digits.
			if len(digits) > 0 {
				cols = append(cols, digits)
			}
		}

		return cols
	}

	grandTotal := 0

	// 5) Use the columns to build full numbers and apply the operator
	for _, r := range problems {
		fmt.Printf("ID %d (width %d): %q %s\n", r.ID, r.Width, r.Numbers, r.Operation)

		// Get digit-columns for this Problem
		cols := splitProblemIntoColumns(r)

		// Convert each digit slice into an int: [3 5 6] -> 356
		var fullNums []int
		for _, digits := range cols {
			val := 0
			for _, d := range digits {
				val = val*10 + d
			}
			fullNums = append(fullNums, val)
		}

		// Apply the operator across the full numbers
		total := 0
		switch r.Operation {
		case "*":
			total = 1
			for _, n := range fullNums {
				total *= n
			}
		case "+":
			total = 0
			for _, n := range fullNums {
				total += n
			}
		default:
			// If an unexpected operator shows up
			fmt.Printf("  Unknown op %q, skipping\n", r.Operation)
			continue
		}

		fmt.Printf("  Columns: %v -> numbers: %v %s => %d\n", cols, fullNums, r.Operation, total)
		grandTotal += total
	}

	fmt.Println("Grand total:", grandTotal)
}

func SplitByFieldWidths(line string, widths []int) []string {
	runes := []rune(line)
	var fields []string

	pos := 0
	for _, w := range widths {
		if pos >= len(runes) {
			// No more characters left; append empty field
			fields = append(fields, "")
			continue
		}

		end := w
		startPos := pos

		// end := pos + w
		// if end > len(runes) {
		// 	end = len(runes)
		// }

		fields = append(fields, string(runes[startPos:end]))
		pos = end
		// fmt.Println(w)
	}

	return fields
}
