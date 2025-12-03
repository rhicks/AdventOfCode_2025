package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func getInput(day int) string {
	inputDir := "inputs"
	inputFile := filepath.Join(inputDir, fmt.Sprintf("day%d.txt", day))

	if _, err := os.Stat(inputFile); err == nil {
		content, _ := os.ReadFile(inputFile)
		return string(content)
	}

	return ""
}

func main() {
	day := 3
	input := getInput(day)

	// day01_1(input)
	// day01_2(input)
	day03_1(input)
	//day02_2(input)

}
