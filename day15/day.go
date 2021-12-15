package main

import (
	"math"
	"math/rand"
	"sort"
	"time"

	tj "github.com/tjhowse/tjgo"
)

func main() {

	rand.Seed(time.Now().UnixNano())
	mainPart1()
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
	println("--------------------")
}
func (m *subMap) printWithPath(path [][2]int) {

	// for y := 0; y < len(m.m); y++ {
	// 	for x := 0; x < len(m.m[0]); x++ {
	for y := 0; y < int(math.Min(50, float64(len(m.m)))); y++ {
		for x := 0; x < int(math.Min(50, float64(len(m.m[0])))); x++ {
			if checkOnPath(x, y, path) {
				print("+")
			} else {
				print(m.m[y][x])
			}
		}
		println()
	}
	println("--------------------")
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

func checkOnPath(x, y int, path [][2]int) bool {
	for _, step := range path {
		if step[0] == x && step[1] == y {
			// This spot it alread on our path, no backtracking!
			return true
		}
	}
	return false
}

func (m *subMap) getAdjacent(x, y int, path [][2]int) [][3]int {
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
			if checkOnPath(dX, dY, path) {
				continue
			}
			adjacent = append(adjacent, [3]int{dX, dY, m.m[dY][dX]})
		}
	}
	// Sort the adjacent cells in-place according to destination cost.
	sort.Slice(adjacent, func(i, j int) bool {
		return adjacent[i][2] < adjacent[j][2]
	})

	return adjacent
}

func (m *subMap) getAdjacent2(pos [2]int) [][2]int {
	adjacent := [][2]int{}
	for _, oY := range []int{-1, 0, 1} {
		for _, oX := range []int{-1, 0, 1} {
			if !(oX == 0 || oY == 0) || ((oX == 0) && (oY == 0)) {
				continue
			}

			dX := pos[0] + oX
			dY := pos[1] + oY
			if !m.checkInBounds(dX, dY) {
				continue
			}
			adjacent = append(adjacent, [2]int{dX, dY})
		}
	}
	return adjacent
}

func (m *subMap) walk(x, y int, path [][2]int, total int, best int) (int, [][2]int) {
	// m.printWithPath(path)
	total += m.m[y][x]
	if total > best {
		return 99999999, path
	}

	path = append(path, [2]int{x, y})
	if x == len(m.m[0])-1 && y == len(m.m)-1 {
		// Destination!
		// println("Destination! ", total)

		return total, path
	}
	adj := m.getAdjacent(x, y, path)
	if len(adj) == 0 {
		// Dead end!
		// println("Dead end vvv")
		// m.printWithPath(path)
		// println("Dead end ^^^")
		// println("Dead end")
		return 99999999, path
	}
	leastCost := best
	shortestPath := [][2]int{}
	for _, a := range adj {
		nextCost, newPath := m.walk(a[0], a[1], path, total, leastCost)
		// if nextCost != 99999999 {
		// 	m.printWithPath(path)
		// 	panic("FUCK!")
		// }
		if nextCost < leastCost {
			// println("New lowst cost chosen:", nextCost)
			leastCost = nextCost
			shortestPath = newPath
		}
	}
	return leastCost, shortestPath
}

// Returns the approximate cost of getting from x,y to the destination
func (m *subMap) heuristic(pos [2]int) int {
	// return int(math.Sqrt(math.Pow(float64(len(m.m[0])-pos[0]), 2)+math.Pow(float64(len(m.m)-pos[1]), 2)))
	return len(m.m[0]) - pos[0] + len(m.m) - pos[1]
}

func getLowest(openSet *map[[2]int]bool, fScore *map[[2]int]int) [2]int {
	min := 999999999
	lowest := [2]int{}
	for node, val := range *openSet {
		if !val {
			continue
		}
		fScore := getScore(*fScore, node)
		if fScore < min {
			min = fScore
			lowest = node
		}
	}
	return lowest
}

func (m *subMap) getCost(pos [2]int) int {
	return m.m[pos[1]][pos[0]]
}

// func (m *subMap) reconstructPath(cameFrom map[[2]int][2]int, current [2]int, goal [2]int) [][2]int {
func (m *subMap) reconstructPath(cameFrom map[[2]int][2]int, current [2]int) [][2]int {

	totalPath := [][2]int{current}
	for {
		_, ok := cameFrom[current]
		if !ok {
			break
		}
		current = cameFrom[current]
		totalPath = append(totalPath, current)
	}
	// current = cameFrom[current]

	// totalPath = append(totalPath, current)
	return totalPath
}

func getScore(m map[[2]int]int, p [2]int) int {
	score, ok := m[p]
	if !ok {
		return 99999999
	}
	return score
}

func (m *subMap) aStar() int {
	start := [2]int{0, 0}
	goal := [2]int{len(m.m[0]) - 1, len(m.m) - 1}

	openSet := make(map[[2]int]bool)
	openSet[start] = true

	cameFrom := make(map[[2]int][2]int)

	gScore := make(map[[2]int]int)
	gScore[start] = 0

	fScore := make(map[[2]int]int)
	fScore[start] = m.heuristic(start)

	for len(openSet) != 0 {

		current := getLowest(&openSet, &fScore)

		// if current[0] == len(m.m[0])-1 && current[1] == len(m.m)-1 {
		if current == goal {
			cost := 0
			path := m.reconstructPath(cameFrom, current)
			for _, s := range path {
				cost += m.getCost(s)
			}
			cost -= m.getCost(start)
			// m.printWithPath(path)
			// println("Path made!")
			return cost
		}
		delete(openSet, current)

		for _, neighbour := range m.getAdjacent2(current) {
			// tentativeGScore := m.heuristic(current) + m.getCost(neighbour)
			tentativeGScore := getScore(gScore, current) + m.getCost(neighbour)
			if tentativeGScore < getScore(gScore, neighbour) {
				cameFrom[neighbour] = current
				gScore[neighbour] = tentativeGScore
				fScore[neighbour] = tentativeGScore + m.heuristic(neighbour)
				openSet[neighbour] = true
			}
		}
	}
	return -1
}

// ..........||||||||||..........||||||||||..........
// 11637517422274862853338597396444961841755517295286 eg top

// 67554889357866599146897761125791887223681299833479 egbot
// ..........||||||||||..........||||||||||..........
// 67554889356755488935675548893567554889356755488935
func (m *subMap) scaleUp(scale int) {
	// newMap := [len(m.m[0]) * 5][len(m.m) * 5]int{}
	if scale == 1 {
		return
	}

	newMap := [][]int{}
	for y := 0; y < len(m.m)*scale; y++ {
		newMap = append(newMap, []int{})
		for x := 0; x < len(m.m[0])*scale; x++ {
			newMap[y] = append(newMap[y], 0)

			if y < len(m.m) {
				// First row of duplications
				if x < len(m.m[0]) {
					// First set of duplications
					newMap[y][x] = m.m[y][x]
				} else {
					newMap[y][x] = newMap[y][x-len(m.m[0])] + 1
				}
			} else {
				if x < len(m.m[0]) {
					// First set of duplications
					newMap[y][x] = newMap[y-len(m.m)][x] + 1
				} else {
					newMap[y][x] = newMap[y-len(m.m)][x] + 1
				}
			}

			if newMap[y][x] > 9 {
				newMap[y][x] = 1
			}
		}
	}
	m.m = newMap
}

func run(filename string, scale int) int {
	var a subMap
	a.load(filename)
	a.scaleUp(scale)
	return a.aStar()

}

func mainPart1() {
	// answer := run("input", 1)
	// answer := run("input", 5)
	// answer := run("input_real", 1)
	answer := run("input_real", 5)
	println(answer)

	// not 2798 or 2800, or 2807 too low
	// 2809 was right... wtf...
	// I'm putting this down to A* being an approximately-best finder, and my implementation is a bit fucky -
	// it doesn't always seem to reconstruct the path all the way to the start? And it doesn't find the
	// same path as the one provided in the example. I would've expected this to give me paths that are longer??? though???

}
