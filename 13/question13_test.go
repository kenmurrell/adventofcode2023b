package main

import (
	"testing"
)

func TestPartOne(t *testing.T) {
	ansTest := 405
	rTest := PartOne("test.txt")
	if rTest != ansTest {
		t.Errorf("Got %d; want %d", rTest, ansTest)
	}
	ansTest2 := 1400
	rTest2 := PartOne("test2.txt")
	if rTest2 != ansTest2 {
		t.Errorf("Got %d; want %d", rTest2, ansTest2)
	}
	ansFull := 28651
	rFull := PartOne("data.txt")
	if rFull != ansFull {
		t.Errorf("Got %d; want %d", rFull, ansFull)
	}
}

func TestPartTwo(t *testing.T) {
	ansTest1 := 400
	rTest1 := PartTwo("test.txt")
	if rTest1 != ansTest1 {
		t.Errorf("Got %d; want %d", rTest1, ansTest1)
	}
	ansTest2 := 900
	rTest2 := PartTwo("test3.txt")
	if rTest2 != ansTest2 {
		t.Errorf("Got %d; want %d", rTest2, ansTest2)
	}

	ansFull := 25450
	rFull := PartTwo("data.txt")
	if rFull != ansFull {
		t.Errorf("Got %d; want %d", rFull, ansFull)
	}
}
