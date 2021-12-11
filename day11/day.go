package main

import (
	"fmt"

	tj "github.com/tjhowse/tjgo"
)

const MAXX = 10
const MAXY = 10

type oct struct {
	m          [MAXX][MAXY]int
	flashed    [MAXX][MAXY]bool
	flashCount int
}

func (o *oct) checkInside(x int, y int) bool {
	if x < 0 || x >= MAXX || y < 0 || y >= MAXY {
		return false
	}
	return true

}

func (o *oct) flash(x, y int) bool {
	if o.flashed[y][x] {
		return false
	}

	o.flashed[y][x] = true
	o.flashCount++
	for j := -1; j < 2; j++ {
		for i := -1; i < 2; i++ {
			if o.checkInside(x+i, y+j) {
				o.m[y+j][x+i]++
				if o.m[y+j][x+i] > 9 {
					o.flash(x+i, y+j)
				}
			}
		}
	}
	return true
}

func (o *oct) increase() {
	for y := 0; y < MAXY; y++ {
		for x := 0; x < MAXX; x++ {
			o.m[y][x]++
		}
	}
}

func (o *oct) checkFlash() bool {
	flashHappened := false
	for y := 0; y < MAXY; y++ {
		for x := 0; x < MAXX; x++ {
			if o.m[y][x] > 9 {
				flashHappened = o.flash(x, y)
			}
		}
	}
	return flashHappened
}

func (o *oct) reset() {
	for y := 0; y < MAXY; y++ {
		for x := 0; x < MAXX; x++ {
			if o.flashed[y][x] {
				o.flashed[y][x] = false
				o.m[y][x] = 0
			}
		}
	}
}

func (o *oct) print() {
	for y := 0; y < MAXY; y++ {
		for x := 0; x < MAXX; x++ {
			fmt.Printf("%2d", o.m[y][x])
		}
		println()
	}
	println("------------------")
}

func main() {

	mainPart1()
	// mainPart2()
}

func mainPart1() {
	a := tj.FileToSlice("input")

	var woo oct

	for y, line := range a {
		for x := 0; x < len(line); x++ {
			woo.m[y][x] = tj.Str2int(string(line[x]))
		}
	}
	prevFlashCount := 0

	for i := 1; i < 10000; i++ {
		woo.increase()
		limit := 0
		for woo.checkFlash() {
			if limit > 10 {
				break
			}
		}
		woo.reset()
		if (woo.flashCount - prevFlashCount) == 100 {
			println("First turn all flash: ", i)
			break
		}
		if i == 100 {
			println("Flashes at turn 100: ", woo.flashCount)
		}
		prevFlashCount = woo.flashCount
	}
}
