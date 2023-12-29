package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type HailStone struct {
	x int
	y int
	z int 
	xv int
	yv int
	zv int
}

type Empty struct {}

func ParseInput(path string) *[]HailStone {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var stones []HailStone
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		line := strings.Split(text, "@")
		positions := strings.Split(line[0], ", ")
		velocities := strings.Split(line[1], ", ")
		hx, _ := strconv.Atoi(strings.Trim(positions[0], " "))
		hy, _ := strconv.Atoi(strings.Trim(positions[1], " "))
		hz, _ := strconv.Atoi(strings.Trim(positions[2], " "))
		hxv, _ := strconv.Atoi(strings.Trim(velocities[0], " "))
		hyv, _ := strconv.Atoi(strings.Trim(velocities[1], " "))
		hzv, _ := strconv.Atoi(strings.Trim(velocities[2], " "))
		hailstone := HailStone{hx,hy,hz,hxv,hyv,hzv}
		stones = append(stones, hailstone)
	}
	return &stones
}



func PartOne(path string, min float64, max float64) int {
	total := 0 
	stones := ParseInput(path)
	for i:=0; i<len(*stones); i++ {
		stone := (*stones)[i]
		for j:=i+1; j<len(*stones); j++{
			stone2 := (*stones)[j]
			x, y, err := interceptXY(stone, stone2)
			if err != nil {
				continue
			}
			if boundarycheckXY(min, max, x, y) {
				total += 1
			}
		}
	}
	return total
}

func boundarycheckXY(min float64, max float64, x float64, y float64) bool {
	return min < x && max > x && min < y && max > y
}

func interceptXY(h1 HailStone, h2 HailStone) (float64, float64, error){
	if h1.xv == 0 || h1.yv == 0 || h2.xv == 0 || h2.yv == 0 {
		return 0, 0, errors.New("speed is zero")
	}
	rateyx1 := float64(h1.yv) / float64(h1.xv)
	ystart1 := float64(h1.y) - rateyx1*float64(h1.x)

	rateyx2 := float64(h2.yv) / float64(h2.xv)
	ystart2 := float64(h2.y) - rateyx2*float64(h2.x)
	if rateyx1 - rateyx2 == 0 {
		return 0, 0, errors.New("divide by zero")
	}
	x := (ystart2 - ystart1) / (rateyx1 - rateyx2)

	y := rateyx1 * x + ystart1
	//past check:
	t1 := (x - float64(h1.x)) / float64(h1.xv)
	t2 := (x - float64(h2.x)) / float64(h2.xv)
	if t1 < 0 || t2 < 0{
		return 0, 0, errors.New("hailstones crossed in the past")
	}
	return x, y, nil
}

func main() {
	start1 := time.Now()
	r1 := PartOne("data.txt", 200000000000000, 400000000000000)
	elapsed1 := time.Since(start1).Seconds()
	fmt.Printf("ANSWER ONE: %d; elapsed %fs\n", r1, elapsed1) 
}
