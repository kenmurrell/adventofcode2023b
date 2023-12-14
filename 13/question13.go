package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func ParseInput(path string) *[][][]bool {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var allLavas [][][]bool
	var lava [][]bool
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		if len(text) == 0 {
			allLavas = append(allLavas, lava)
			lava = nil
		}else {
			lineSplit := strings.Split(text, "")
			var line []bool
			for j:=0; j<len(lineSplit); j++{
				if lineSplit[j] == "#"{
					line = append(line, true)
				} else {
					line = append(line, false)
				}
			}
			lava = append(lava, line)
		}
	}
	if len(lava) > 0 {
		allLavas = append(allLavas, lava)
	}
	return &allLavas
}

func PartOne(path string) int {
	allLavas := ParseInput(path)
	total := 0
	for p := 0; p < len(*allLavas); p++ {
		pattern := (*allLavas)[p]
		col := -1
		row := -1 

		col, row = run(&pattern, -1, -1)
		if col >= 0 {
			total += (col + 1)
		} else if row >=0 {
			total += 100 * (row + 1)
		}
	}
	
	return total
}

func PartTwo(path string) int {
	allLavas := ParseInput(path)
	total := 0
	for p := 0; p < len(*allLavas); p++ {
		pattern := (*allLavas)[p]
		nmutations := len(pattern) * len(pattern[0]) 
		col, row := run(&pattern, -1, -1)
		for n:=0; n<nmutations; n++ {
			mpattern := mutate(&pattern, n)
			c2, r2 := run(mpattern, col, row)
			if c2 >= 0 {
				total += (c2 + 1)
				break
			} else if r2 >=0 {
				total += (r2 + 1) *100
				break
			} 
		}
	}
	
	return total
}

func mutate(arr *[][]bool, mnumber int) *[][] bool {
	ctr:=0
	newArr := make([][]bool, len(*arr))
	for row:=0; row<len(*arr); row++ {
		newLine := make([]bool, len((*arr)[0]))
		for col:=0; col<len((*arr)[0]); col++{
			val := (*arr)[row][col]
			if ctr == mnumber {
				val = !val
			}
			newLine[col] = val
			ctr+=1
		}
		newArr[row] = newLine
	}

	return &newArr
} 

func run(pattern *[][]bool, col int, row int) (int, int) {
	c2 := -1
	r2 := -1
	for idx:=0; idx<len((*pattern)[0]); idx++ {
		if testIdx(pattern, idx) {
			if idx != col {
				return idx, r2
			}
		}
	}
	patternRot := rotateArr(pattern)
	for idx:=0; idx<len((*patternRot)[0]); idx++ {
		if testIdx(patternRot, idx) {
			if idx != row {
				return c2, idx
			}
		}
	}
	return c2, r2
}

func rotateArr(arr *[][]bool) *[][]bool {
	var newArr [][]bool
	for i:=0; i<len((*arr)[0]); i++ {
		line := make([]bool, len(*arr))
		for j:=0; j<len(*arr); j++ {
			x := (*arr)[j][i]
			line[j] = x
		}
		newArr = append(newArr, line)
	}
	return &newArr
}

func testIdx(arr *[][]bool, i int) bool {
	for _, line := range *arr {
		if !testIdxOnArr(&line, i) {
			return false
		}
	}
	return true
}

func testIdxOnArr(arr *[]bool, i int) bool {
	L := i
	R := i + 1 
	if L < 0 || R >= len(*arr) {
		return false
	}
	for L >=0 && R < len(*arr) {
		if (*arr)[L] != (*arr)[R] {
			return false
		}
		L -=1
		R +=1
	}
	return true
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
