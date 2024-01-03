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
	ansTest1 := 1
	rTest1 := PartTwo("test2.txt")
	if rTest1 != ansTest1 {
		t.Errorf("Got %d; want %d", rTest1, ansTest1)
	}

	ansTest2 := 525152
	rTest2 := PartTwo("test.txt")
	if rTest2 != ansTest2 {
		t.Errorf("Got %d; want %d", rTest2, ansTest2)
	}

	ansFull := 83317216247365
	rFull := PartTwo("data.txt")
	if rFull != ansFull {
		t.Errorf("Got %d; want %d", rFull, ansFull)
	}
}
