package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func ParseInput(path string) *[]*[]int {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var sequencelist []*[]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		var newlist []int
		for _, s := range strings.Split(text, " ") {
			i, _ := strconv.Atoi(s)
			newlist = append(newlist, i)
		}
		sequencelist = append(sequencelist, &newlist)
	}
	return &sequencelist
}

func DoWork(seq *[]int) *[]*[]int {
	var work []*[]int
	level := 0
	work = append(work, seq)
	prevLevel := work[level]
	for !AllZero(prevLevel) {
		var newLevel []int
		for i := 0; i < len(*prevLevel)-1; i++ {
			newLevel = append(newLevel, (*prevLevel)[i+1]-(*prevLevel)[i])
		}
		work = append(work, &newLevel)
		level++
		prevLevel = work[level]
	}
	return &work
}

func PartOne(path string) int {
	total := 0
	sl := ParseInput(path)
	for _, seq := range *sl {
		work := DoWork(seq)
		n := 0
		for i := range *work {
			lastIdx := len(*work) - 1 - i
			levellist := (*work)[lastIdx]
			n = (*levellist)[len(*levellist)-1] + n
		}
		total += n
	}

	return total
}

func AllZero(list *[]int) bool {
	for _, l := range *list {
		if l != 0 {
			return false
		}
	}
	return true
}

func PartTwo(path string) int {
	total := 0
	sl := ParseInput(path)
	for _, seq := range *sl {
		work := DoWork(seq)
		n := 0
		for i := range *work {
			lastIdx := len(*work) - 1 - i
			levellist := (*work)[lastIdx]
			n = (*levellist)[0] - n
		}
		total += n
	}
	return total
}

func main() {
	r1 := PartOne("data.txt")
	fmt.Printf("ANSWER ONE: %d\n", r1)
	r2 := PartTwo("data.txt")
	fmt.Printf("ANSWER TWO: %d\n", r2)
}
