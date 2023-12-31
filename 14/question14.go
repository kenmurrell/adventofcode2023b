package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func ParseInput(path string) *[][]string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var board [][]string
	level := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		lineSplit := strings.Split(text, "")
		if level == 0 {
			board = make([][]string, len(lineSplit))
		}
		for i:=0; i< len(lineSplit); i++ {
			board[i] = append(board[i], lineSplit[i])
		}
		level++
	}

	return &board
}

func PartOne(path string) int {
	b := ParseInput(path)
	tiltboard(b)
	r := calculate(b)
	return r
}

func PartTwo(path string) int {
	b := ParseInput(path)
	for i := 0; i < 1000; i++ {
		for j:=0; j<4; j++{
			tiltboard(b)
			rotateRight(b) //this is
			rotateRight(b) //incredibly dumb
			rotateRight(b) //but I've forgotten how to rotate CC
		}
		if i > 0 && i % 100 == 0 {
			fmt.Printf("iter: %d\n", i)
		}
    }
	r := calculate(b)
	return r
}

func tiltboard(board *[][]string) {
	for col:=0; col<len(*board); col++ {
		run := (*board)[col]
		for row := 0; row<len(run); row++ {
			if run[row] != "O"{
				continue
			}
			row_itr := row
			for row_itr >= 0 {
				if (row_itr == 0 || run[row_itr-1] == "#" || run[row_itr-1] == "O") {
					if run[row_itr] == "." {
						run[row_itr] = "O"
						run[row] = "."
						break
					}else if row == row_itr {
						// block already in place
						break
					}
				}
				row_itr -=1
			}
		}
		(*board)[col] = run
	}
}

func calculate(board *[][]string) int{
	r := 0
	for col:=0; col<len(*board); col++ {
		run := (*board)[col]
		for row := 0; row < len(run); row++ {
			if run[row]=="O"{
				r+= len(run) - row
			}
		}
	}
	return r
}
func rotateRight(board *[][]string) {
	ln := len(*board)
	for row:=0; row< ln; row++ {
		for col:=row; col < ln; col++ {
			temp1 := (*board)[row][col]
			temp2 := (*board)[col][row]
			(*board)[col][row] = temp1
			(*board)[row][col] = temp2
		}
	}
	for row:=0; row< ln; row++ {
		run := (*board)[row]
		for i, j := 0, len(run)-1; i < j; i, j = i+1, j-1 {
			run[i], run[j] = run[j], run[i]
		}
		(*board)[row] = run
	}
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

func printBoard(x *[][]string){
	for col := 0; col<len(*x);col++ {
		for row := 0; row<len((*x)[0]);row++ {
			fmt.Print((*x)[row][col])
		}
		fmt.Print("\n")
	}
	fmt.Print("\n")
}
