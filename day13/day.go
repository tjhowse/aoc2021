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

	println("s.maxX: ", s.maxX)
	println("s.maxY: ", s.maxY)
	println("len X: ", len(s.s[0]))
	println("len Y: ", len(s.s))
	println("count: ", count)
}

func (s *sheet) print() {
	for i := 0; i <= s.maxY; i++ {
		for j := 0; j <= s.maxX; j++ {
			// for i := 0; i <= 14; i++ {
			// for j := 0; j <= 10; j++ {
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
	println("Counting:")
	println("s.maxX:", s.maxX)
	println("s.maxY:", s.maxY)
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

func (s *sheet) fold2(x, y int) {
	if x != 0 {
		for j := 0; j <= s.maxY; j++ {
			for i := 0; i < x; i++ {
				s.s[j][i] = s.s[j][i] || s.s[j][s.maxX-i]
				s.s[j][s.maxX-i] = false
			}
		}
		s.maxX = x - 1
	} else {
		for j := 0; j < y; j++ {
			for i := 0; i <= s.maxX; i++ {
				s.s[j][i] = s.s[j][i] || s.s[s.maxY-j][i]
				s.s[s.maxY-j][i] = false
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

type better struct {
	s map[[2]int]bool
}

func (b *better) init() {
	b.s = make(map[[2]int]bool)
}

func (b *better) set(x, y int) {
	b.s[[2]int{x, y}] = true
}
func (b *better) count() int {
	c := 0
	for _, v := range b.s {
		if v {
			c++
		}
	}
	return c
}
func (b *better) flip(x, y int) {
	if x != 0 {
		// flipping on x
		for k, v := range b.s {
			if !v || k[0] < x {
				continue
			}
			b.s[k] = false
			rightOfSplit := k[0] - x
			k[0] = x - rightOfSplit
			b.s[k] = true
		}
	} else {
		// flipping on y
		for k, v := range b.s {
			if !v || k[1] < y {
				continue
			}
			b.s[k] = false
			downOfSplit := k[1] - y
			k[1] = y - downOfSplit
			b.s[k] = true
		}
	}
}
func (b *better) draw() {
	var mX, mY int

	for k, v := range b.s {
		if !v {
			continue
		}
		if k[0] > mX {
			mX = k[0]
		}
		if k[1] > mY {
			mY = k[1]
		}
	}
	for y := 0; y <= mY; y++ {
		for x := 0; x <= mX; x++ {
			if b.s[[2]int{x, y}] {
				print("#")
			} else {
				print(" ")

			}
		}
		println()
	}
	println("--------------------")
}

func run2(filename string) better {
	a := tj.FileToSlice(filename)
	rand.Seed(time.Now().UnixNano())

	dotsList := []spot{}
	// dotSet := map[spot]bool{}
	hello := better{}
	hello.init()

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
		hello.set(s.x, s.y)
		// dotSet[s] = true
	}
	return hello
}

func mainPart1() {
	hello := run2("input")

	println(hello.count())
	hello.draw()
	hello.flip(0, 7)
	hello.draw()
	hello.flip(5, 0)
	hello.draw()
	println(hello.count())
	hello = run2("input_real")

	println(hello.count())
	hello.flip(655, 0)
	hello.flip(0, 447)
	hello.flip(327, 0)
	hello.flip(0, 223)
	hello.flip(163, 0)
	hello.flip(0, 111)
	hello.flip(81, 0)
	hello.flip(0, 55)
	hello.flip(40, 0)
	hello.flip(0, 27)
	hello.flip(0, 13)
	hello.flip(0, 6)
	hello.draw()

	println(hello.count())
	// for k, v := range dotSet {
	// 	println("k: ", k.x, ",", k.y, " v: ", v)
	// }

	// return s
	return
	s := run("input")

	// s.mark(5, 4)
	// s.mark(5, 5)
	// s.mark(5, 6)
	// s.mark(8, 4)
	s.print()
	return
	s.fold2(0, 7)
	s.print()
	s.fold2(5, 0)
	s.print()
	println(s.count())

	println()
	return

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
