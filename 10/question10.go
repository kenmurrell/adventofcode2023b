package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"time"
)

type Action func(*Point, *int)

type Point struct {
	Column int
	Row int
}

type Empty struct {

}

func (p1 *Point) Add(p2 *Point) Point{
	return Point{p1.Column+p2.Column, p1.Row+p2.Row}
}

func CreateChar2Point() *map[string][]Point{
	char2Point := make(map[string][]Point)
	char2Point["|"] = []Point{{0, 1}, {0, -1}}
	char2Point["-"] = []Point{{1, 0}, {-1, 0}}
	char2Point["7"] = []Point{{0, 1}, {-1, 0}}
	char2Point["F"] = []Point{{0, 1}, {1, 0}}
	char2Point["L"] = []Point{{0, -1}, {1, 0}}
	char2Point["J"] = []Point{{0, -1}, {-1, 0}}
	char2Point["S"] = []Point{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	return &char2Point
}

func ParseInput(path string) (*Point, *map[Point]string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	pipeMap := make(map[Point]string)
	i := 0
	var c Point
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		lineSplit := strings.Split(text, "")
		for j:=0; j<len(lineSplit); j++{
			p := Point{j, i}
			letter := lineSplit[j]
			if letter == "S" {
				c = p
			}
			pipeMap[p] = letter
		}
		i++
	}
	return &c, &pipeMap
}

func PartOne(path string) int {
	start, pipeMap := ParseInput(path)
	char2Point := CreateChar2Point()
	var wg sync.WaitGroup
	var mu sync.Mutex
	lenMap := make(map[Point]int)

	wg.Add(1)
	var addToMap Action = func(p *Point, i *int){
		if val, ok := lenMap[*p]; !ok || val > *i{
			lenMap[*p] = *i+1
		} 
	  };
	go traverse(&mu, &wg, 0, *start, *start, pipeMap, addToMap, char2Point)
	wg.Wait()
	highest := 0
	for _, v := range lenMap {
		if v > highest{
			highest = v
		}
	}
	return highest
}

func traverse(
	mu *sync.Mutex,
	wg *sync.WaitGroup,
	currLen int,
	prevp Point,
	p Point,
	pipeMap *map[Point]string,
	act Action,
	char2Point *map[string][]Point) {
		defer wg.Done()
		letter := (*pipeMap)[p]
		directions := (*char2Point)[letter]
		for _, dir := range directions{
			nextPoint := p.Add(&dir)
			if nextPoint == prevp {
				continue
			}
			nextLetter := (*pipeMap)[nextPoint]
			if nextLetter == "S" {
				return
			}
			nextletterdirs := (*char2Point)[nextLetter]
			for _, nextLetterDir := range nextletterdirs{
				if nextPoint.Add(&nextLetterDir) != p {
					continue
				}
				mu.Lock()
				act(&nextPoint, &currLen)
				mu.Unlock()
				wg.Add(1)
				go traverse(mu, wg, currLen+1, p, nextPoint, pipeMap, act, char2Point)
			}
		}			
}

func PartTwo(path string) int {
	return 0
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
