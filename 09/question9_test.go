package main

import (
	"testing"
)

func TestPartOne(t *testing.T) {
	ansTest := 114
	rTest := PartOne("test.txt")
	if rTest != ansTest {
		t.Errorf("Got %d; want %d", rTest, ansTest)
	}
	ansFull := 1938800261
	rFull := PartOne("data.txt")
	if rFull != ansFull {
		t.Errorf("Got %d; want %d", rFull, ansFull)
	}
}

func TestPartTwo(t *testing.T) {
	ansTest := 2
	rTest := PartTwo("test.txt")
	if rTest != ansTest {
		t.Errorf("Got %d; want %d", rTest, ansTest)
	}
	ansFull := 1112
	rFull := PartTwo("data.txt")
	if rFull != ansFull {
		t.Errorf("Got %d; want %d", rFull, ansFull)
	}
}
