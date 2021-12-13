package main

import (
	"math/rand"
	"strings"
	"time"

	tj "github.com/tjhowse/tjgo"
)

type spot struct {
	x int
	y int
}

type sheet struct {
	s    [][]bool
	maxX int
	maxY int
}

func (s *sheet) setSize(x, y int) {
	s.maxX = x
	s.maxY = y
	count := 0
	for i := 0; i <= s.maxY; i++ {
		s.s = append(s.s, []bool{})
		for j := 0; j <= s.maxX; j++ {
			s.s[i] = append(s.s[i], false)
			count++
		}
	}
	println("count: ", count)
}

func (s *sheet) print() {
	for i := 0; i <= s.maxY; i++ {
		for j := 0; j <= s.maxX; j++ {
			if s.s[i][j] {
				print("#")
			} else {
				print(".")
			}

		}
		println()
	}
	println("-----------------")
}

func (s *sheet) mark(x, y int) {
	s.s[y][x] = true
}
func (s *sheet) count() int {
	c := 0
	for i := 0; i <= s.maxY; i++ {
		for j := 0; j <= s.maxX; j++ {
			if s.s[i][j] {
				c++
				// println(j, ",", i)
			}
		}
	}
	return c

}

func (s *sheet) fold(x, y int) {
	if x == 0 {
		println("Folding along y:", y)
		for i := y + 1; i <= s.maxY; i++ {
			for j := 0; j <= s.maxX; j++ {
				s.s[s.maxY-i][j] = s.s[i][j] || s.s[s.maxY-i][j]
				s.s[i][j] = false
			}
		}
		s.maxY = y - 1
	} else {
		println("Folding along x:", x)
		for i := 0; i <= s.maxY; i++ {
			for j := x + 1; j <= s.maxX; j++ {
				s.s[i][s.maxX-j] = s.s[i][j] || s.s[i][s.maxX-j]
				s.s[i][j] = false
			}
		}
		s.maxX = x - 1
	}
}

func (s *sheet) fold2(x, y int) {
	if x != 0 {
		for j := 0; j <= s.maxY; j++ {
			for i := 0; i <= s.maxX; i++ {
				s.s[j][i] = s.s[j][i] || s.s[j][s.maxX-i]
			}
		}
		s.maxX = x - 1
	} else {
		for j := 0; j <= s.maxY; j++ {
			for i := 0; i <= s.maxX; i++ {
				s.s[j][i] = s.s[j][i] || s.s[s.maxY-j][i]
			}
		}
		s.maxY = y - 1
	}
}

func main() {

	// mainPart1()
	// start := time.Now()
	mainPart1()
	// end := time.Now()
	// println(end.Sub(start).Milliseconds())
}

func run(filename string) sheet {
	a := tj.FileToSlice(filename)
	rand.Seed(time.Now().UnixNano())

	dotsList := []spot{}

	var maxX, maxY int

	for _, a := range a {
		// println(a)
		if a == "" {
			break
		}
		split := strings.Split(a, ",")
		var s spot
		s.x = tj.Str2int(split[0])
		s.y = tj.Str2int(split[1])
		if s.x > maxX {
			maxX = s.x
		}
		if s.y > maxY {
			maxY = s.y
		}
		dotsList = append(dotsList, s)
	}
	var s sheet
	s.setSize(maxX, maxY)

	// println(len(dotsList))
	for i := 0; i < len(dotsList); i++ {
		s.mark(dotsList[i].x, dotsList[i].y)
	}
	return s
}

func mainPart1() {
	s := run("input")

	s.print()
	s.fold2(0, 7)
	s.print()
	s.fold2(5, 0)
	s.print()
	println(s.count())

	println()

	s = run("input_real")
	println("maxX: ", s.maxX, "maxY: ", s.maxY)

	s.fold2(655, 0)
	println(s.count())
	if s.count() != 731 {
		println("Failure!")
	} else {
		println("Success!")

	}

	// Not 866 or 865
	//731 was answer for part 1

}
