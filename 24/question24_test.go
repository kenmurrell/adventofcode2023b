package main

import (
	"testing"
)

func TestPartOne(t *testing.T) {
	ansTest := 2
	rTest := PartOne("test.txt", 7, 27)
	if rTest != ansTest {
		t.Errorf("Got %d; want %d", rTest, ansTest)
	}
	
	ansFull := 31921
	rFull := PartOne("data.txt", 200000000000000, 400000000000000)
	if rFull != ansFull {
		t.Errorf("Got %d; want %d", rFull, ansFull)
	}
}
