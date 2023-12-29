package main

import (
	"testing"
)

func TestPartOne(t *testing.T) {
	ansTest := 19114
	rTest := PartOne("test.txt")
	if rTest != ansTest {
		t.Errorf("Got %d; want %d", rTest, ansTest)
	}

	ansFull := 406849
	rFull := PartOne("data.txt")
	if rFull != ansFull {
		t.Errorf("Got %d; want %d", rFull, ansFull)
	}
}
