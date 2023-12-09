package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func ParseInput(path string) (*[]string, *map[string]*map[string]string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var instr []string
	pathways := make(map[string]*map[string]string)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		if len(instr) == 0 {
			instr = append(instr, strings.Split(text, "")...)
			continue
		}
		if len(text) == 0 {
			continue
		}

		line := strings.Split(text, " = ")
		lineTrim := line[1][1 : len(line[1])-1]
		line1Splt := strings.Split(lineTrim, ", ")
		path := map[string]string{
			"L": line1Splt[0],
			"R": line1Splt[1],
		}
		pathways[line[0]] = &path

	}
	return &instr, &pathways
}

func PartOne(path string) int {
	instr, pathways := ParseInput(path)
	i := 0
	cur := "AAA"
	for true {
		paths := (*pathways)[cur]
		idx := i % len(*instr)
		i++
		dir := (*instr)[idx]
		cur = (*paths)[dir]
		if cur == "ZZZ" {
			break
		}
	}
	return i
}

func PartTwo(path string) uint64 {
	instr, pathways := ParseInput(path)
	var startings []string
	for k, _ := range *pathways {
		if strings.LastIndex(k, "A") == 2 {
			startings = append(startings, k)
		}
	}
	var hits []uint64
	for _, s := range startings {
		count := 0
		cur := s
		for true {
			idx := count % len(*instr)
			dir := (*instr)[idx]
			paths := (*pathways)[cur]
			cur = (*paths)[dir]
			if strings.LastIndex(cur, "Z") == 2 {
				break
			}
			count++
		}
		hits = append(hits, uint64(count+1))
	}

	lm := LowestMultiple(hits[0], hits[1], hits[2:]...)
	return lm
}

func LowestMultiple(a uint64, b uint64, n ...uint64) uint64 {
	r := (a * b) / GreatestDemoninator(a, b)
	for i := 0; i < len(n); i++ {
		r = LowestMultiple(r, n[i])
	}
	return r
}

func GreatestDemoninator(num1 uint64, num2 uint64) uint64 {
	for num2 != 0 {
		t := num2
		num2 = num1 % num2
		num1 = t
	}
	return num1
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
