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

type Spring struct {
	data   string
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
		var groups []int
		for _, d := range strings.Split(textSplit[1], ",") {
			val, _ := strconv.Atoi(d)
			groups = append(groups, val)
		}
		springs = append(springs, Spring{textSplit[0], groups})
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

func combosPerSpring(s *Spring) int {
	comboCount := 0
	unknCnt := 0
	for _, val := range s.data {
		if val == '?' {
			unknCnt += 1
		}
	}
	possibilities := int(math.Pow(2, float64(unknCnt)))
	for itr := 0; itr < possibilities; itr++ {
		var mutated string
		unknownId := 0
		for _, val := range s.data {
			if val == '?' {
				x := (itr>>unknownId)&0x1 == 1
				if x {
					mutated += "."
				} else {
					mutated += "#"
				}
				unknownId += 1
			} else {
				mutated += string(val)
			}
		}
		if check(mutated, &s.groups) {
			comboCount += 1
		}
	}
	return comboCount
}

func check(data string, group *[]int) bool {
	count := 0
	groupIdx := 0
	for _, d := range data {
		if d == '#' {
			count += 1
		} else if d == '.' && count > 0 {
			if groupIdx >= len(*group) || (*group)[groupIdx] != count {
				return false
			}
			groupIdx++
			count = 0
		}
	}
	if count > 0 {
		if groupIdx >= len(*group) || (*group)[groupIdx] != count {
			return false
		}
		groupIdx++
	}
	return groupIdx == len(*group)
}

func PartTwo(path string) int {
	total := 0
	var comboMap = make(map[string]int)
	for _, spring := range *ParseInput(path) {
		dupedData := strings.Repeat("?"+spring.data, 5)[1:]
		var dupedGroups []int
		for i := 0; i < 5; i++ {
			dupedGroups = append(dupedGroups, spring.groups...)
		}
		total += calculate(dupedData, dupedGroups, &comboMap, -1)
	}
	return total
}

func calculate(data string, groups []int, comboMap *map[string]int, idx int) int {
	key := fmt.Sprintf("%s %s %d", data, arr2Str(&groups), idx)
	if m, ok := (*comboMap)[key]; ok {
		return m
	}

	if len(data) == 0 && len(groups) == 0 && idx <= 0 {
		return 1
	} else if len(data) == 0 {
		return 0
	}
	output := 0
	if data[0] == '#' {
		if idx == 0 || (idx == -1 && len(groups) == 0) {
			return 0
		} else if idx == -1 {
			idx = groups[0]
			groups = groups[1:]
		}
		output = calculate(data[1:], groups, comboMap, idx-1)
	} else if data[0] == '.' {
		if idx <= 0 {
			output = calculate(data[1:], groups, comboMap, -1)
		}
	} else if data[0] == '?' {
		// calculate if it were bad
		baddata := "#" + data[1:]
		r1 := calculate(baddata, groups, comboMap, idx)
		// calculate if it were good
		opdata := "." + data[1:]
		r2 := calculate(opdata, groups, comboMap, idx)
		output = r1 + r2
	} else {
		//uh oh
	}

	(*comboMap)[key] = output
	return output
}

func arr2Str(arr *[]int) string {
	b := make([]string, len(*arr))
	for i, v := range *arr {
		b[i] = strconv.Itoa(v)
	}
	return strings.Join(b, ",")
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
