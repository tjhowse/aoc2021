package main

import (
	"strings"

	tj "github.com/tjhowse/tjgo"
)

func main() {

	mainPart1()
	// mainPart2()
}

// func checkIn(c rune, s string) bool {
// 	for _, i := range s {
// 		if i == s {
// 			return true
// 		}
// 	}
// 	return false
// }

func overlap(s1, s2 string) string {
	var result string
	for _, c := range s1 {
		for _, b := range s2 {
			if c == b {
				result += string(b)
			}
		}
	}

	return result
}

func decode(s string) int {
	switch len(s) {
	case 2:
		return 1
	case 3:
		return 7
	case 4:
		return 4
	case 7:
		return 8
	}

	return -1
}

// 1: ab
// 4: eafb
// 6: cdfgeb
// 7: dab
// 8: acedgfb

// 2,3,5
// cdfbe
// gcdfa
// fbcad

// 9, 0
// cefabd 9,0
// cagedb 9,0

// 1: be
// 4: cgeb
// 7: edb
// 8: cfbegad

// 0: fgaecd

// 7: edb

// 6:
// agebfd
// cbdgef

// fdcge
// fecdb
// fabcd

func mainPart1() {
	a := tj.FileToSlice("input")

	sum := 0

	grandTotal := 0

	for _, i := range a {
		set := make(map[int]string)
		b := strings.Split(i, "|")
		signals := strings.Fields(b[0])
		outputs := strings.Fields(b[1])
		for _, i := range outputs {
			if decode(i) >= 0 {
				sum++
			}
		}
		for _, i := range signals {
			decoded := decode(i)
			if decoded >= 0 {
				set[decoded] = i
			}
		}
		// 1, 4, 7, 8 known
		for _, i := range signals {
			decoded := decode(i)
			if decoded < 0 {
				// Decode them otherwise:
				if len(i) == 6 {
					// 0, 6, 9
					if len(overlap(i, set[7])) == 2 {
						set[6] = i
						continue
					}
					if len(overlap(i, set[4])) == 3 {
						set[0] = i
						continue
					}
					set[9] = i
				} else {
					//2,3,5
					if len(overlap(i, set[7])) == 3 {
						set[3] = i
						continue
					}

					if len(overlap(i, set[4])) == 2 {
						set[2] = i
						continue
					}
					set[5] = i
				}
			}
		}

		var total int
		index := 1000
		for _, i := range outputs {
			for number, s := range set {
				if len(s) == len(i) && len(overlap(s, i)) == len(s) {
					total += number * index
					index /= 10
				}
			}
		}
		grandTotal += total
	}
	println("Sum:", sum)
	println("grandTotal:", grandTotal)

}
