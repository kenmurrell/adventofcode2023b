package main

import (
	"testing"
)

func TestPartOne(t *testing.T) {
	ansTest := 21
	rTest := PartOne("test.txt")
	if rTest != ansTest {
		t.Errorf("Got %d; want %d", rTest, ansTest)
	}

	ansFull := 7361
	rFull := PartOne("data.txt")
	if rFull != ansFull {
		t.Errorf("Got %d; want %d", rFull, ansFull)
	}
}

func TestPartTwo(t *testing.T) {
	ansTest1 := 525152
	rTest1 := PartTwo("test.txt")
	if rTest1 != ansTest1 {
		t.Errorf("Got %d; want %d", rTest1, ansTest1)
	}
}
