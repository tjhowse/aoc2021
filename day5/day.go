package main

import (
	"strings"

	tj "github.com/tjhowse/tjgo"
)

type floorMap struct {
	m    [1000][1000]int
	maxX int
	maxY int
}

func (f *floorMap) line(s string) {
	sp := strings.Split(s, " ")
	sX := tj.Str2int(strings.Split(sp[0], ",")[0])
	sY := tj.Str2int(strings.Split(sp[0], ",")[1])
	eX := tj.Str2int(strings.Split(sp[2], ",")[0])
	eY := tj.Str2int(strings.Split(sp[2], ",")[1])

	f.maxX = 1000
	f.maxY = 1000

	if !((sX == eX) || (sY == eY)) {
		// return // Part 1
		stepX := 1
		stepY := 1
		if sX > eX {
			stepX = -1
		}
		if sY > eY {
			stepY = -1
		}
		// Diagonal
		for j := sY; j != (eY + stepY); j += stepY {
			for i := sX; i != (eX + stepX); i += stepX {
				if (sY-j)-(sX-i) == 0 || (eY-j)+(eX-i) == 0 {
					f.m[i][j]++
				}
			}
		}
	} else {
		// Horizontal/vertical
		if sX > eX {
			sX, eX = eX, sX
		}
		if sY > eY {
			sY, eY = eY, sY
		}
		for j := sY; j <= eY; j++ {
			for i := sX; i <= eX; i++ {
				f.m[i][j]++
			}
		}
	}
}

func (f *floorMap) draw() {
	for i := 0; i <= f.maxY; i++ {
		for j := 0; j <= f.maxX; j++ {
			print(f.m[j][i])
		}
		println()
	}
}

func (f *floorMap) count() int {
	ret := 0
	for i := 0; i < f.maxY; i++ {
		for j := 0; j < f.maxX; j++ {
			if f.m[j][i] > 1 {
				ret++
			}
		}
	}
	return ret
}

func main() {
	a := tj.FileToSlice("input")

	// s := strings.Split(a[0], ",")

	// numbers := []int{}

	m := floorMap{}

	for _, n := range a {
		m.line(n)
	}
	// m.line(a[0])
	// m.line(a[1])

	// m.draw()
	println(m.count())

}
