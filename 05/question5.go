package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

type ParseSeeds func(string, *[]int)

type Converter struct {
	dest   []int
	src    []int
	length []int
}

func (c *Converter) Convert(num int) int {
	for i, _ := range c.dest {
		if num >= c.src[i] && num < c.src[i]+c.length[i] {
			return c.dest[i] + num - c.src[i]
		}
	}

	return num
}

func ParseInput(path string, fn ParseSeeds) ([]int, *[]Converter) { //stupid but whatever
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var seeds []int
	var cons []Converter

	var t_dest []int
	var t_src []int
	var t_length []int

	for scanner.Scan() {
		text := scanner.Text()
		if strings.Contains(text, "seeds:") {
			cline := strings.Split(text, ": ")
			numbers := strings.Split(cline[1], " ")
			for _, n := range numbers {
				integer, _ := strconv.Atoi(n)
				seeds = append(seeds, integer)
			}
			continue
		}
		if strings.Contains(text, "-to-") {
			continue
		}
		if len(strings.Trim(text, "\n")) == 0 {
			if len(t_dest) == 0 {
				continue
			}
			CreateAndReset(&cons, &t_dest, &t_src, &t_length)
		} else {
			numbers := strings.Split(text, " ")
			d, _ := strconv.Atoi(numbers[0])
			t_dest = append(t_dest, d)
			s, _ := strconv.Atoi(numbers[1])
			t_src = append(t_src, s)
			l, _ := strconv.Atoi(numbers[2])
			t_length = append(t_length, l)
		}
	}
	if len(t_dest) > 0 {
		CreateAndReset(&cons, &t_dest, &t_src, &t_length)
	}
	return seeds, &cons
}

func CreateAndReset(cons *[]Converter, t_dest *[]int, t_src *[]int, t_length *[]int) {
	*cons = append(*cons, Converter{
		dest:   *t_dest,
		src:    *t_src,
		length: *t_length,
	})
	*t_dest = make([]int, 2)
	*t_src = make([]int, 2)
	*t_length = make([]int, 2)
}

func ParseSeeds2(line string, seeds *[]int) {
	starting := -1
	cline := strings.Split(line, ": ")
	numbers := strings.Split(cline[1], " ")
	for _, n := range numbers {
		integer, _ := strconv.Atoi(n)
		if starting < 0 {
			starting = integer
		} else {
			for i := 0; i < integer; i++ {
				*seeds = append(*seeds, starting+i)
			}
			starting = -1
		}
	}
}

func ParseSeeds1(line string, seeds *[]int) {
	cline := strings.Split(line, ": ")
	numbers := strings.Split(cline[1], " ")
	for _, n := range numbers {
		integer, _ := strconv.Atoi(n)
		*seeds = append(*seeds, integer)
	}
}

func PartOne(path string) int {
	least := 0
	var num int
	seeds, cons := ParseInput(path, ParseSeeds1)
	for _, s := range seeds {
		num = s
		for _, c := range *cons {
			num = c.Convert(num)
		}
		if num < least || least == 0 {
			least = num
		}
	}
	return least
}

func PartTwo(path string) int {
	var wg sync.WaitGroup
	var mu sync.Mutex
	least := math.MaxInt32
	seeds, cons := ParseInput(path, ParseSeeds2)
	seedChan := make(chan int)

	//Add one for seed iter
	wg.Add(1)
	go func() {
		defer wg.Done()
		start := 0
		cnt := false
		for _, n := range seeds {

			if cnt {
				wg.Add(n)
				for i := 0; i < n; i++ {
					seedChan <- start + i
				}
				cnt = false
			} else {
				start = n
				cnt = true
			}
		}
	}()
	go func() {
		i := 0
		for n := range seedChan {
			i += 1
			go func() {
				RunTwo(&wg, &mu, n, &least, cons)
			}()
			if i%10000000 == 0 {
				fmt.Printf("itr: %dM\n", i/1000000)
			}
		}
	}()

	wg.Wait()
	time.Sleep(time.Second * 10)
	close(seedChan)
	return least
}

func RunTwo(wg *sync.WaitGroup, mu *sync.Mutex, num int, least *int, cons *[]Converter) {
	defer wg.Done()
	for i := range *cons {
		mu.Lock()
		num = (*cons)[i].Convert(num)
		mu.Unlock()
	}
	mu.Lock()
	if num < *least {
		*least = num
		fmt.Printf("Least: %d\n", *least)
	}
	mu.Unlock()
}

func main() {
	r1 := PartOne("data.txt")
	fmt.Printf("ANSWER ONE: %d\n", r1)
	r2 := PartTwo("data.txt")
	fmt.Printf("ANSWER TWO: %d\n", r2)
}
