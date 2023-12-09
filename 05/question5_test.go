package main

import (
	"testing"
)

func TestPartOne(t *testing.T) {
	ansTest := 35
	rTest := PartOne("test.txt")
	if rTest != ansTest {
		t.Errorf("Got %d; want %d", rTest, ansTest)
	}
	ansFull := 227653707
	rFull := PartOne("data.txt")
	if rFull != ansFull {
		t.Errorf("Got %d; want %d", rFull, ansFull)
	}
}

func TestPartTwo(t *testing.T) {
	ansTest := 46
	rTest := PartTwo("test.txt")
	if rTest != ansTest {
		t.Errorf("Got %d; want %d", rTest, ansTest)
	}
	ansFull := 78775052
	rFull := PartTwo("data.txt")
	if rFull != ansFull {
		t.Errorf("Got %d; want %d", rFull, ansFull)
	}
}
