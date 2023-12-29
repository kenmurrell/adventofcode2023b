package main

import (
	"testing"
)

func TestPartOne(t *testing.T) {
	ansTest := 374
	rTest := PartOne("test.txt")
	if rTest != ansTest {
		t.Errorf("Got %d; want %d", rTest, ansTest)
	}
	
	ansFull := 9609130
	rFull := PartOne("data.txt")
	if rFull != ansFull {
		t.Errorf("Got %d; want %d", rFull, ansFull)
	}
}

func TestPartTwo(t *testing.T) {
	var ansTest1 int64 = 1030
	rTest1 := PartTwo("test.txt", 10)
	if rTest1 != ansTest1 {
		t.Errorf("Got %d; want %d", rTest1, ansTest1)
	}

	var ansTest2 int64 = 8410
	rTest2 := PartTwo("test.txt", 100)
	if rTest2 != ansTest2 {
		t.Errorf("Got %d; want %d", rTest2, ansTest2)
	}
	
	var ansFull int64 = 702152204842
	rFull := PartTwo("data.txt", 1000000)
	if rFull != ansFull {
		t.Errorf("Got %d; want %d", rFull, ansFull)
	}
}