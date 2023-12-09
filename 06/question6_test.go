package main

import (
	"testing"
)

func TestPartOne(t *testing.T) {
	ansTest := 288
	rTest := PartOne("test.txt")
	if rTest != ansTest {
		t.Errorf("Got %d; want %d", rTest, ansTest)
	}
	ansFull := 4568778
	rFull := PartOne("data.txt")
	if rFull != ansFull {
		t.Errorf("Got %d; want %d", rFull, ansFull)
	}
}

func TestPartTwo(t *testing.T) {
	ansTest := 71503
	rTest := PartTwo("test.txt")
	if rTest != ansTest {
		t.Errorf("Got %d; want %d", rTest, ansTest)
	}
	ansFull := 28973936
	rFull := PartTwo("data.txt")
	if rFull != ansFull {
		t.Errorf("Got %d; want %d", rFull, ansFull)
	}
}
