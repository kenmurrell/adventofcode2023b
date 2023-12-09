package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Race struct {
	Distance int
	Time     int
}

func ParseInput(path string) []Race {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var races []Race
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ":")
		info := strings.Split(line[1], " ")
		for i := len(info) - 1; i >= 0; i-- {
			if info[i] == "" || info[i] == " " {
				info = append(info[:i], info[i+1:]...)
			}
		}
		if strings.Contains(line[0], "Time") {
			for _, t := range info {
				integer, _ := strconv.Atoi(t)
				races = append(races, Race{
					Distance: 0,
					Time:     integer,
				})

			}
		}
		if strings.Contains(line[0], "Distance") {
			for i, t := range info {
				integer, _ := strconv.Atoi(t)
				x := races[i]
				races[i] = Race{
					Distance: integer,
					Time:     x.Time,
				}
			}
		}
	}
	return races
}

func ParseInput2(path string) Race {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var t int
	var d int
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ":")
		info := strings.ReplaceAll(line[1], " ", "")

		if strings.Contains(line[0], "Time") {
			t, _ = strconv.Atoi(info)
		}
		if strings.Contains(line[0], "Distance") {
			d, _ = strconv.Atoi(info)
		}
	}
	return Race{
		Distance: d,
		Time:     t,
	}
}

type scoreCalc func(int) string

func PartOne(path string) int {
	races := ParseInput(path)
	total := 1
	for _, r := range races {
		waystowin := 0
		for c := 1; c < r.Time; c++ {
			dis := c * (r.Time - c)
			if dis > r.Distance {
				waystowin++
			}
		}
		total *= waystowin
	}
	return total
}

func PartTwo(path string) int {
	r := ParseInput2(path)
	waystowin := 0
	for c := 1; c < r.Time; c++ {
		dis := c * (r.Time - c)
		if dis > r.Distance {
			waystowin++
		}
	}
	return waystowin
}

func main() {
	r1 := PartOne("data.txt")
	fmt.Printf("ANSWER ONE: %d\n", r1)
	r2 := PartTwo("data.txt")
	fmt.Printf("ANSWER TWO: %d\n", r2)
}
