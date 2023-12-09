package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func PartOne(path string) int {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		first := 0
		last := 0
		line := strings.Split(scanner.Text(), "")
		for i := 0; i < len(line); i++ {
			if i, err := strconv.Atoi(line[i]); err == nil {
				if first == 0 {
					first = i
				}
				last = i
			}
		}
		sum += (first * 10) + last
	}
	return sum
}

func PartTwo(path string) int {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	m := make(map[string]int)
	m["one"] = 1
	m["two"] = 2
	m["three"] = 3
	m["four"] = 4
	m["five"] = 5
	m["six"] = 6
	m["seven"] = 7
	m["eight"] = 8
	m["nine"] = 9

	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var sb strings.Builder
		first := 0
		last := 0
		line := strings.Split(scanner.Text(), "")
		for i := 0; i < len(line); i++ {
			if integer, err := strconv.Atoi(line[i]); err == nil {
				if first == 0 {
					first = integer
				}
				last = integer
			} else {
				sb.WriteString(line[i])
				for k, v := range m {
					index := strings.LastIndex(sb.String(), k)
					if index >= 0 && index+len(k) == sb.Len() {
						if first == 0 {
							first = v
						}
						last = v
					}
				}
			}
		}
		sum += (first * 10) + last
	}
	return sum
}

func main() {
	least1 := PartOne("data.txt")
	fmt.Printf("ANSWER ONE: %d\n", least1)
	least2 := PartTwo("data.txt")
	fmt.Printf("ANSWER TWO: %d\n", least2)
}
