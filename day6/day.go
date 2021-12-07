package main

import (
	"strings"
	"time"

	tj "github.com/tjhowse/tjgo"
)

type fish struct {
	c int
}

func (f *fish) tick() bool {
	// Returns true if this creates another fish
	if f.c == 0 {
		f.c = 6
		return true
	}
	f.c--
	return false

}

func main() {

	mainPart1()
	// mainPart1v2()
}

func mainPart1() {
	a := tj.FileToSlice("input")

	start := time.Now()
	f := []fish{}

	s := strings.Split(a[0], ",")
	for _, i := range s {
		f = append(f, fish{c: tj.Str2int(i)})
	}
	lenSeq := []int{}

	for i := 0; i < 80; i++ {
		// for i := 0; i < 256; i++ { // Erp...
		newFish := []fish{}
		for j := 0; j < len(f); j++ {
			if f[j].tick() {
				newFish = append(newFish, fish{c: 8})
			}
		}
		f = append(f, newFish...)
		lenSeq = append(lenSeq, len(f))
	}
	println("Part 1:", len(f))

	b := 8
	q := 6
	for i := len(lenSeq); i < 256; i += 1 {
		lenSeq = append(lenSeq, lenSeq[i-1]+(lenSeq[i-b-1]-lenSeq[i-b-2])+(lenSeq[i-q-1]-lenSeq[i-q-2]))
	}
	println("Part 2:", lenSeq[len(lenSeq)-1])
	end := time.Now()
	println("Elapsed: ", end.Sub(start).Microseconds())
}

func mainPart1v2() {
	a := tj.FileToSlice("input")

	start := time.Now()

	f := []fish{}

	s := strings.Split(a[0], ",")
	for _, i := range s {
		f = append(f, fish{c: tj.Str2int(i)})
	}
	lenSeq := []int{}

	for i := 0; i < 10; i++ {
		newFish := []fish{}
		for j := 0; j < len(f); j++ {
			if f[j].tick() {
				newFish = append(newFish, fish{c: 8})
			}
		}
		f = append(f, newFish...)
		lenSeq = append(lenSeq, len(f))
	}
	// println("Part 1:", len(f)) // Not accurate since we're not running for the full 80 days

	b := 8
	q := 6
	for i := len(lenSeq); i < 256; i += 1 {
		lenSeq = append(lenSeq, lenSeq[i-1]+(lenSeq[i-b-1]-lenSeq[i-b-2])+(lenSeq[i-q-1]-lenSeq[i-q-2]))
	}
	println("Part 2:", lenSeq[len(lenSeq)-1])

	end := time.Now()
	println("Elapsed: ", end.Sub(start).Microseconds())
}
