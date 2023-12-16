package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Step struct {
	code string
	add bool
	number int
}

func ParseInput2(path string) *[]Step {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var steps []Step
	pattern := regexp.MustCompile("([a-z]+)([=-])([0-9]*)")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		lineSplit := strings.Split(text, ",")
		for _, s := range lineSplit {
			m := pattern.FindAllStringSubmatch(s, -1)
			if len(m) <1 || len(m[0]) < 3{
				panic("fuck")
			}
			add := false
			num := -1
			if m[0][2] == "=" {
				add = true
			}
			if len(m[0][3]) > 0 {
				num, _ = strconv.Atoi(m[0][3])
			}
			step := Step{
				code: m[0][1],
				add: add,
				number: num,
			}
			steps = append(steps, step)
		} 
	}
	return &steps
}

func ParseInput1(path string) *[]string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var steps []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		lineSplit := strings.Split(text, ",")
		steps = append(steps, lineSplit...)
	}
	return &steps
}

func PartOne(path string) int {
	steps := ParseInput1(path)
	sum := 0
	for _, s := range *steps {
		sum += hash(s)
	}
	
	return sum
}

func PartTwo(path string) int {
	steps := ParseInput2(path)
	boxes := make([][]Step, 256)
	for _, step := range *steps {
		boxnum := hash(step.code)
		box := boxes[boxnum]
		box = AddOrReplace(&box, step)
		boxes[boxnum] = box
	}
	return calculate(&boxes)
}

func AddOrReplace(box *[]Step, newStep Step) []Step{
	index := -1
	for i, step :=  range *box {
		if step.code == newStep.code {
			index = i
			break
		}
	}
	if index >= 0 {
		if newStep.add {
			(*box)[index] = newStep
			return *box
		} else {
			return append((*box)[:index], (*box)[index+1:]...)
		}
	} else {
		if newStep.add{
			return append(*box, newStep)
		} else {
			return *box
		}
	}
}

func hash(str string) int {
	value := 0
	for i := range str {
		ascii := int(str[i])
		value += ascii
		value = value * 17
		value = value % 256
	}
	return value
}

func calculate(boxes *[][]Step) int {
	sum := 0
	for box_num, b := range *boxes {
		for step_num, s := range b {
			sum += (box_num +1) * (step_num + 1) * s.number
		}
	}
	return sum
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
