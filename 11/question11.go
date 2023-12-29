package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
	"time"
)

type Galaxy struct {
	x int
	y int
}


func ParseInput(path string) *[][]bool {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var grid [][]bool
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		var row []bool
		for _, item := range strings.Split(text, ""){
			if item == "#" {
				row = append(row, true)
			} else {
				row = append(row, false)
			}
		}
		grid = append(grid, row)
	}
	return &grid
}

func PartOne(path string) int {
	grid := ParseInput(path)
	// find the rows to expand
	rowSkips := findRowSkips(grid)
	// find the columns to expand
	colSkips := findColSkips(grid)
	// tag galaxies
	galaxies:= tagGalaxies(grid, rowSkips, colSkips, 1)
	//measure distances
	return int(measureDistances(galaxies))
}

func findColSkips(grid *[][]bool) *[]int {
	var colskips []int
	for n:=0; n<len(*grid); n++ {
		var col []bool
		for r:=0; r<len((*grid)[0]); r++ {
			col = append(col, (*grid)[r][n])
		}
		if isEmpty(&col){
			colskips = append(colskips, n)
		}
	}
	return &colskips
}

func findRowSkips(grid *[][]bool) *[]int {
	var rowskips []int
	for n:=0; n<len(*grid); n++ {
		if isEmpty(&(*grid)[n]){
			rowskips = append(rowskips, n)
		}
	}
	return &rowskips
}

func isEmpty(arr *[]bool) bool {
	for _, i := range *arr {
		if i {
			return false
		}
	}
	return true
}

func measureDistances(galaxies *[]Galaxy) int64 {
	fmt.Print("Measuring distances...")
	var total int64 =0
	for n, gax := range *galaxies {
		for g2:=n+1; g2<len(*galaxies); g2++ {
			gax2 := (*galaxies)[g2]
			xdistance := math.Abs(float64(gax.x) - float64(gax2.x))
			ydistance := math.Abs(float64(gax.y) - float64(gax2.y))
			total += int64(xdistance) + int64(ydistance)
		}
	}
	fmt.Println("DONE")
	return total
}

func tagGalaxies(grid *[][]bool, rowSkips *[]int, colSkips *[]int, skip int) *[]Galaxy {
	fmt.Print("Tagging galaxies...")
	var galaxies []Galaxy
	for row:=0; row <len(*grid); row++ {
		for col:=0; col<len((*grid)[0]); col++ {
			if (*grid)[row][col]{
				rowsSkipped := 0
				colsSkipped := 0
				for _, r := range *rowSkips {
					if r < row {
						rowsSkipped += skip
					}
				}
				for _, c := range *colSkips {
					if c < col {
						colsSkipped += skip
					}
				}
				galaxies = append(galaxies, Galaxy{row+rowsSkipped, col+colsSkipped})
			}
		}
	}
	fmt.Println("DONE")
	return &galaxies
}

func PartTwo(path string, exp int) int64 {
	grid := ParseInput(path)
	// find the rows to expand
	rowSkips := findRowSkips(grid)
	// find the columns to expand
	colSkips := findColSkips(grid)
	// tag galaxies
	galaxies:= tagGalaxies(grid, rowSkips, colSkips, exp -1)
	//measure distances
	return measureDistances(galaxies)
}

func main() {
	start1 := time.Now()
	r1 := PartOne("data.txt")
	elapsed1 := time.Since(start1).Seconds()
	fmt.Printf("ANSWER ONE: %d; elapsed %fs\n", r1, elapsed1) 
	start2 := time.Now()
	r2 := PartTwo("data.txt", 1000000)
	elapsed2 := time.Since(start2).Seconds()
	fmt.Printf("ANSWER TWO: %d; elapsed %fs\n", r2, elapsed2)
}
