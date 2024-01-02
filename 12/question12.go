package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

type Cond int

const (
	Good Cond = iota
	Bad
	Unkn
)

type Spring struct {
	data   []Cond
	groups []int
}

func ParseInput(path string) *[]Spring {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var springs []Spring
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		textSplit := strings.Split(text, " ")
		var data []Cond
		for _, d := range strings.Split(textSplit[0], "") {
			if d == "." {
				data = append(data, Good)
			} else if d == "#" {
				data = append(data, Bad)
			} else if d == "?" {
				data = append(data, Unkn)
			}
		}
		var groups []int
		for _, d := range strings.Split(textSplit[1], ",") {
			val, _ := strconv.Atoi(d)
			groups = append(groups, val)
		}
		springs = append(springs, Spring{data, groups})
	}
	return &springs
}

func PartOne(path string) int {
	totalcount := 0
	springs := ParseInput(path)
	for _, s := range *springs {
		totalcount += combosPerSpring(&s)
	}
	return totalcount
}

func PartTwo(path string) int {
	totalcount := 0
	springs := ParseInput(path)
	unFoldedSprings := make([]Spring, len(*springs), len(*springs))
	for i, s := range *springs {
		var data2 []Cond
		for x := 0; x < 5; x++ {
			if x > 0 {
				data2 = append(data2, Unkn)
			}
			data2 = append(data2, s.data...)
		}
		var group2 []int
		for x := 0; x < 5; x++ {
			group2 = append(group2, s.groups...)
		}
		unFoldedSprings[i] = Spring{data2, group2}
	}
	for i, s := range unFoldedSprings {
		totalcount += combosPerSpring(&s)
		fmt.Printf("Completed %d\n", i)
	}
	return totalcount
}

func combosPerSpring(s *Spring) int {
	comboCount := 0
	unknCnt := 0
	for _, val := range s.data {
		if val == Unkn {
			unknCnt += 1
		}
	}
	possibilities := int(math.Pow(2, float64(unknCnt)))
	for itr := 0; itr < possibilities; itr++ {
		mutated := make([]Cond, len(s.data), len(s.data))
		unknownId := 0
		for index, val := range s.data {
			if val == Unkn {
				x := (itr>>unknownId)&0x1 == 1
				if x {
					mutated[index] = Good
				} else {
					mutated[index] = Bad
				}
				unknownId += 1
			} else {
				mutated[index] = val
			}
		}
		if isEqual(count(&mutated), &s.groups) {
			comboCount += 1
		}
	}
	return comboCount
}

func isEqual(group1 *[]int, group2 *[]int) bool {
	if len(*group1) != len(*group2) {
		return false
	}
	for i := 0; i < len(*group1); i++ {
		if (*group1)[i] != (*group2)[i] {
			return false
		}
	}
	return true
}

func count(data *[]Cond) *[]int {
	count := 0
	var output []int
	for _, d := range *data {
		if d == Bad {
			count += 1
		}
		if d == Good {
			if count > 0 {
				output = append(output, count)
			}
			count = 0
		}
	}
	if count > 0 {
		output = append(output, count)
	}
	return &output
}

func main() {
	start1 := time.Now()
	r1 := PartOne("data.txt")
	elapsed1 := time.Since(start1).Seconds()
	fmt.Printf("ANSWER ONE: %d; elapsed %fs\n", r1, elapsed1)
	start2 := time.Now()
	r2 := PartTwo("data.txt")
	elapsed2 := time.Since(start2).Seconds()
	fmt.Printf("ANSWER TWO: %d; elapsed %fs\n", r2, elapsed2)
}
