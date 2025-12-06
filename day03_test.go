package main

import (
	"testing"
)

func TestLineIsString(t *testing.T, line string) {
	if line != string {
		t.Error("Line is not of type String")
	}
}
