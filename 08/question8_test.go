package main

import (
	"testing"
)

func TestPartOne(t *testing.T) {
	ansTest1 := 2
	rTest1 := PartOne("test.txt")
	if rTest1 != ansTest1 {
		t.Errorf("Got %d; want %d", rTest1, ansTest1)
	}
	ansTest2 := 6
	rTest2 := PartOne("test2.txt")
	if rTest2 != ansTest2 {
		t.Errorf("Got %d; want %d", rTest2, ansTest2)
	}
	ansFull := 13019
	rFull := PartOne("data.txt")
	if rFull != ansFull {
		t.Errorf("Got %d; want %d", rFull, ansFull)
	}
}

func TestPartTwo(t *testing.T) {
	ansTest := 6
	rTest := PartTwo("test3.txt")
	if rTest != uint64(ansTest) {
		t.Errorf("Got %d; want %d", rTest, ansTest)
	}
	var ansFull uint64 = 13524038372771
	rFull := PartTwo("data.txt")
	if rFull != ansFull {
		t.Errorf("Got %d; want %d", rFull, ansFull)
	}
}
