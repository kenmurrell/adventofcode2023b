package main

import (
	"testing"
)

func TestPartOne(t *testing.T) {
	ansTest := 297
	rTest := PartOne("test.txt")
	if rTest != ansTest {
		t.Errorf("Got %d; want %d", rTest, ansTest)
	}
	ansFull := 52974
	rFull := PartOne("data.txt")
	if rFull != ansFull {
		t.Errorf("Got %d; want %d", rFull, ansFull)
	}
}

func TestPartTwo(t *testing.T) {
	ansTest := 363
	rTest := PartTwo("test.txt")
	if rTest != ansTest {
		t.Errorf("Got %d; want %d", rTest, ansTest)
	}
	ansFull := 53340
	rFull := PartTwo("data.txt")
	if rFull != ansFull {
		t.Errorf("Got %d; want %d", rFull, ansFull)
	}
}
