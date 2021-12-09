package main

import (
	tj "github.com/tjhowse/tjgo"
)

func main() {

	mainPart1()
	// mainPart2()
}

// const mapSizeX = 10
// const mapSizeY = 5

const mapSizeX = 100
const mapSizeY = 100

func robustGet(x int, y int, floorMap [mapSizeY][mapSizeX]int, d int) int {
	if x < 0 || x >= mapSizeX || y < 0 || y >= mapSizeY {
		return d
	}
	return floorMap[y][x]

}

type point struct {
	x int
	y int
}

func mainPart1() {
	a := tj.FileToSlice("input")

	intSlice := [mapSizeY][mapSizeX]int{}

	for i, line := range a {
		for j, v := range line {
			// print(tj.Str2int(string(v)))
			intSlice[i][j] = tj.Str2int(string(v))
		}
	}

	offset := 99999999

	sum := 0

	lowPoints := []point{}

	for i := 0; i < mapSizeY; i++ {
		for j := 0; j < mapSizeX; j++ {
			if robustGet(j-1, i, intSlice, offset) <= intSlice[i][j] {
				continue
			}
			if robustGet(j+1, i, intSlice, offset) <= intSlice[i][j] {
				continue
			}
			if robustGet(j, i-1, intSlice, offset) <= intSlice[i][j] {
				continue
			}
			if robustGet(j, i+1, intSlice, offset) <= intSlice[i][j] {
				continue
			}
			// println("i: ", i, "j:", j, "val:", intSlice[i][j])
			sum += (intSlice[i][j] + 1)
			lowPoints = append(lowPoints, point{x: j, y: i})
		}
	}
	// Not 1842
	println("Sum: ", sum)

	for _, p := range lowPoints {
		var copied [mapSizeY][mapSizeX]int
		for i := 0; i < mapSizeY; i++ {
			for j := 0; j < mapSizeX; j++ {
				copied[i][j] = intSlice[i][j]
			}
		}
		copied[p.y][p.x] = -1
		for {
			changed := 0
			for i := 0; i < mapSizeY; i++ {
				for j := 0; j < mapSizeX; j++ {
					if copied[i][j] == 9 || copied[i][j] == -1 {
						continue
					}
					if robustGet(j-1, i, copied, offset) == -1 {
						copied[i][j] = -1
						changed++
						continue
					}
					if robustGet(j+1, i, copied, offset) == -1 {
						copied[i][j] = -1
						changed++
						continue
					}
					if robustGet(j, i-1, copied, offset) == -1 {
						copied[i][j] = -1
						changed++
						continue
					}
					if robustGet(j, i+1, copied, offset) == -1 {
						copied[i][j] = -1
						changed++
						continue
					}
				}
			}
			if changed == 0 {
				pitSize := 0
				for i := 0; i < mapSizeY; i++ {
					for j := 0; j < mapSizeX; j++ {
						if copied[i][j] < 0 {
							pitSize++
						}
					}
					// println()
				}
				println(pitSize)
				break
			}
		}
		// for i := 0; i < mapSizeY; i++ {
		// 	for j := 0; j < mapSizeX; j++ {
		// 		print(copied[i][j])
		// 	}
		// 	println()
		// }

		// fmt.Println("%v", p)
	}

	// intSlice := []int{}
	// strSlice := []string{}
	// woo := strings.Split(a[0], ",")
	// println("\"", woo[0], "\"")

	// for _, line := range a {
	// 	println(line)
	// }

}
