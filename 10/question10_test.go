package main

import (
	"testing"
)

func TestPartOne(t *testing.T) {
	ansTest1 := 4
	rTest1 := PartOne("test.txt")
	if rTest1 != ansTest1 {
		t.Errorf("Got %d; want %d", rTest1, ansTest1)
	}

	ansTest2 := 8
	rTest2 := PartOne("test2.txt")
	if rTest2 != ansTest2 {
		t.Errorf("Got %d; want %d", rTest2, ansTest2)
	}	
	
	ansFull := 6886
	rFull := PartOne("data.txt")
	if rFull != ansFull {
		t.Errorf("Got %d; want %d", rFull, ansFull)
	}
}
