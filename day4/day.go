package main

import (
	"strings"

	tj "github.com/tjhowse/tjgo"
)

type board struct {
	b [5][5]int
}

func (b *board) print() {
	for q := 0; q < 5; q++ {
		for h := 0; h < 5; h++ {
			print(b.b[q][h], " ")
		}
		println()
	}
}

func (b *board) number(i int) {
	for q := 0; q < 5; q++ {
		for h := 0; h < 5; h++ {
			if b.b[q][h] == i {
				b.b[q][h] = -1
			}
		}
	}
}

func (b *board) sum() int {
	total := 0
	for q := 0; q < 5; q++ {
		for h := 0; h < 5; h++ {
			if b.b[q][h] > 0 {
				total += b.b[q][h]
			}
		}
	}
	return total
}

func (b *board) checkWin() bool {

	for q := 0; q < 5; q++ {
		win := true
		for h := 0; h < 5; h++ {
			if b.b[q][h] > 0 {
				win = false
			}
		}
		if win {
			return true
		}
	}
	for q := 0; q < 5; q++ {
		win := true
		for h := 0; h < 5; h++ {
			if b.b[h][q] > 0 {
				win = false
			}
		}
		if win {
			return true
		}
	}
	return false
}

func main() {
	a := tj.FileToSlice("input")

	numbers := []int{}

	for _, n := range strings.Split(a[0], ",") {
		numbers = append(numbers, tj.Str2int(n))
	}

	boards := []board{}

	for i := 2; i < len(a); i += 6 {
		newBoard := board{}
		for q := 0; q < 5; q++ {
			split := strings.Fields(a[i+q])
			for h, j := range split {
				newBoard.b[q][h] = tj.Str2int(string(j))
			}
		}
		boards = append(boards, newBoard)
	}

	for _, number := range numbers {
		losers := []board{}
		for i := 0; i < len(boards); i++ {
			boards[i].number(number)
			if boards[i].checkWin() {
				if len(boards) == 1 {
					println("Last winner's score: ", boards[0].sum(), number, ": ", boards[0].sum()*number)
					return
				}
			}
			if !boards[i].checkWin() {
				losers = append(losers, boards[i])
			}
		}
		boards = losers
	}
}
