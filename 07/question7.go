package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Hand struct {
	cards string
	bet   int
	S     Score
}

type Score struct {
	a int
	b int
}

func CalculateScore1(cards string) Score {
	handranks := make(map[string]int)
	handranks["2"] = 0
	handranks["3"] = 1
	handranks["4"] = 2
	handranks["5"] = 3
	handranks["6"] = 4
	handranks["7"] = 5
	handranks["8"] = 6
	handranks["9"] = 7
	handranks["T"] = 8
	handranks["J"] = 9
	handranks["Q"] = 10
	handranks["K"] = 11
	handranks["A"] = 12

	p := make(map[string]int, 1)
	r := 0
	for i, s := range strings.Split(cards, "") {
		r += handranks[s] * int(math.Pow(100, float64(5-1-i)))
		p[s] += 1
	}
	d := 0
	t := false
	for _, v := range p {
		if v == 5 || v == 4 {
			return Score{v + 1, r}
		}
		if v == 3 {
			t = true
		} else if v == 2 {
			d += 1
		}
	}
	if t {
		if d > 0 {
			return Score{4, r}
		}
		return Score{3, r}
	}
	return Score{d, r}
}

func CalculateScore2(cards string) Score {
	handranks := make(map[string]int)
	handranks["2"] = 2
	handranks["3"] = 3
	handranks["4"] = 4
	handranks["5"] = 5
	handranks["6"] = 6
	handranks["7"] = 7
	handranks["8"] = 8
	handranks["9"] = 9
	handranks["T"] = 10
	handranks["J"] = 1
	handranks["Q"] = 11
	handranks["K"] = 12
	handranks["A"] = 13

	p := make(map[string]int, 1)
	r := 0
	maxkey := ""
	maxval := 0

	for i, s := range strings.Split(cards, "") {
		r += handranks[s] * int(math.Pow(100, float64(5-1-i)))
		p[s] += 1
		if s != "J" && p[s] > maxval {
			maxkey = s
			maxval = p[s]
		}
	}
	if maxval > 0 {
		p[maxkey] += p["J"]
		p["J"] = 0
	}
	d := 0
	t := false
	for _, v := range p {
		if v == 5 || v == 4 {
			return Score{v + 1, r}
		}
		if v == 3 {
			t = true
		} else if v == 2 {
			d += 1
		}
	}
	if t {
		if d > 0 {
			return Score{4, r}
		}
		return Score{3, r}
	}
	return Score{d, r}
}

func ParseInput(path string, fn Calc) []Hand {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var hands []Hand
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		integer, _ := strconv.Atoi(line[1])
		score := fn(line[0])
		hands = append(hands, Hand{
			cards: line[0],
			bet:   integer,
			S:     score,
		})
	}
	return hands
}

type Calc func(string) Score

func PartOne(path string) int {
	hands := ParseInput(path, CalculateScore1)
	sort.SliceStable(hands, func(i, j int) bool {
		if hands[i].S.a == hands[j].S.a {
			return hands[i].S.b < hands[j].S.b
		}
		return hands[i].S.a < hands[j].S.a
	})
	total := 0
	for i, h := range hands {
		total += (i + 1) * h.bet
	}
	return total
}

func PartTwo(path string) int {
	hands := ParseInput(path, CalculateScore2)
	sort.SliceStable(hands, func(i, j int) bool {
		if hands[i].S.a == hands[j].S.a {
			return hands[i].S.b < hands[j].S.b
		}
		return hands[i].S.a < hands[j].S.a
	})
	total := 0
	for i, h := range hands {
		total += (i + 1) * h.bet
	}
	return total
}

func main() {
	r1 := PartOne("data.txt")
	fmt.Printf("ANSWER ONE: %d\n", r1)
	r2 := PartTwo("data.txt")
	fmt.Printf("ANSWER TWO: %d\n", r2)
}
