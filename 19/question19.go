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

type Workflow struct {
	cat    string
	comp   string
	val    int
	result string
}

type Part map[string]int

func ParseInput(path string) (*map[string][]Workflow, *[]Part) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var parts []Part
	workflows := make(map[string][]Workflow)
	pattern := regexp.MustCompile("([a-z]+)([><])?([0-9]+)\\:([a-zA-Z]+)")
	scanner := bufio.NewScanner(file)
	workflowsDone := false
	for scanner.Scan() {
		text := scanner.Text()
		if len(text) == 0 {
			workflowsDone = true
			continue
		}
		if workflowsDone {
			text = strings.Trim(text, "{}")
			part := Part{}
			for _, item := range strings.Split(text, ",") {
				temp := strings.Split(item, "=")
				integer, _ := strconv.Atoi(temp[1])
				part[temp[0]] = integer
			}
			parts = append(parts, part)

		} else {
			temp := strings.Split(text, "{")
			name := temp[0]
			var workflowlist []Workflow
			for _, item := range strings.Split(temp[1], ",") {
				m := pattern.FindAllStringSubmatch(item, -1)
				if len(m) == 1 {
					val, _ := strconv.Atoi(m[0][3])
					workflowlist = append(workflowlist, Workflow{
						cat:    m[0][1],
						comp:   m[0][2],
						val:    val,
						result: m[0][4],
					})
				} else {
					trimmed := strings.Trim(item, "}")
					workflowlist = append(workflowlist, Workflow{
						cat:    "",
						comp:   "",
						val:    0,
						result: trimmed,
					})
				}
			}
			workflows[name] = workflowlist
		}
	}
	return &workflows, &parts
}

func PartOne(path string) int {
	workflows, parts := ParseInput(path)
	sum := 0
	for _, p := range *parts {
		w := (*workflows)["in"]
		for true {
			r := runWorkflow(&p, &w)
			if r == "A" {
				for _, v := range p {
					sum += v
				}
				break
			} else if r == "R" {
				break
			} else {
				w = (*workflows)[r]
			}
		}
	}
	return sum
}

func runWorkflow(p *Part, ws *[]Workflow) string {
	for _, w := range *ws {
		empty := len(w.comp) == 0
		greater := w.comp == ">" && (*p)[w.cat] > w.val
		lesser := w.comp == "<" && (*p)[w.cat] < w.val
		if empty || greater || lesser {
			return w.result
		}
	}
	return "" //wont ever happen
}

func main() {
	start1 := time.Now()
	r1 := PartOne("data.txt")
	elapsed1 := time.Since(start1).Seconds()
	fmt.Printf("ANSWER ONE: %d; elapsed %fs\n", r1, elapsed1)
}
