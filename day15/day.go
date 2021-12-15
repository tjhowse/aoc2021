package main

import (
	"math/rand"
	"sort"
	"time"

	tj "github.com/tjhowse/tjgo"
)

func main() {

	rand.Seed(time.Now().UnixNano())
	mainPart1()
}

func mainPart1() {
	// hello := run("input_real")
	run("input")

}

type subMap struct {
	m    [][]int
	been map[[2]int]bool
}

func (m *subMap) load(filename string) {
	a := tj.FileToSlice(filename)

	for j, line := range a {
		m.m = append(m.m, []int{})
		for _, char := range line {
			m.m[j] = append(m.m[j], tj.Str2int(string(char)))
			// m.m[j][i] = tj.Str2int(string(char))
		}
	}
	m.been = make(map[[2]int]bool)
}
func (m *subMap) print() {

	for y := 0; y < len(m.m); y++ {
		for x := 0; x < len(m.m[0]); x++ {
			if m.checkBeen(x, y) {
				print("X")
			} else {
				print(m.m[y][x])
			}
		}
		println()
	}
}

func (m *subMap) checkBeen(x, y int) bool {
	return m.been[[2]int{x, y}]
}

func (m *subMap) markBeen(x, y int) {
	m.been[[2]int{x, y}] = true
}

func (m *subMap) checkInBounds(x, y int) bool {
	if x >= len(m.m[0]) || x < 0 || y >= len(m.m) || y < 0 {
		return false
	}
	return true
}

func (m *subMap) getAdjacent(x, y int) [][3]int {
	adjacent := [][3]int{}
	for _, oY := range []int{-1, 0, 1} {
		for _, oX := range []int{-1, 0, 1} {
			if !(oX == 0 || oY == 0) || ((oX == 0) && (oY == 0)) {
				continue
			}

			dX := x + oX
			dY := y + oY
			if !m.checkInBounds(dX, dY) {
				continue
			}
			// This is within bounds
			if m.checkBeen(dX, dY) {
				continue
			}
			// We haven't been here before.
			adjacent = append(adjacent, [3]int{dX, dY, m.m[dY][dX]})
		}
	}
	// Sort the adjacent cells in-place according to destination cost.
	sort.Slice(adjacent, func(i, j int) bool {
		return adjacent[i][2] < adjacent[j][2]
	})

	return adjacent
}

func walk(m subMap, x, y int, total int, depth int) int {
	// if total > 10 {
	// 	return total
	// }
	depth++
	// m.print()
	// println("Walking to: ", x, y, "Total:", total, "Depth: ", depth)
	total += m.m[y][x]
	m.markBeen(x, y)
	if x == len(m.m[0])-1 && y == len(m.m)-1 {
		// Destination!
		println("Destination! ", total)
		return total
	}
	adj := m.getAdjacent(x, y)
	if len(adj) == 0 {
		// Dead end!
		println("Dead end")
		return 9999999999999
	}
	leastCost := 999999999999999
	fuck := [2]int{}
	for _, a := range adj {
		// for i, a := range adj {
		// println("Taking path ", depth, i)
		nextCost := walk(m, a[0], a[1], total, depth)
		if nextCost < leastCost {
			// println("New lowst cost chosen:", nextCost)
			leastCost = nextCost
			fuck[0], fuck[1] = a[0], a[1]
		}
	}
	return leastCost
}

func run(filename string) {
	var a subMap
	a.load(filename)
	a.print()

	// adj := a.getAdjacent(2, 2)
	// for _, x := range adj {
	// 	println(x[0], x[1], x[2])
	// }
	t := walk(a, 0, 0, 0, 0)
	println(t)
}
