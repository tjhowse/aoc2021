package main

import (
	tj "github.com/tjhowse/tjgo"
)

func main() {
	a := tj.FileToSlice("input")
	b := []int{}

	for _, d := range a {
		b = append(b, tj.Str2int(d))
	}
	prev := 99999999
	increase := 0

	for i := 0; i < len(b)-2; i++ {
		avg := b[i] + b[i+1] + b[i+2]
		if avg > prev {
			increase++
		}
		prev = avg
	}
	println(increase)

}

// Part 1
// func main() {
// 	a := tj.FileToSlice("input")
// 	b := []int{}

// 	for _, d := range a {
// 		b = append(b, tj.Str2int(d))
// 	}

// 	prev := 99999999
// 	increase := 0
// 	for _, d := range b {
// 		if d > prev {
// 			increase++
// 		}
// 		prev = d
// 	}
// 	println(increase)

// }

// func main() {
// 	min := 347312
// 	max := 805915

// FirstLoop:
// 	for i := min; i <= max; i++ {
// 		s := tj.Int2str(i)
// 		dupe := false
// 		var prev rune
// 		for _, c := range s {
// 			if c == prev {
// 				dupe = true
// 				break
// 			}
// 			if c < prev {
// 				break FirstLoop
// 			}
// 			prev = c
// 		}
// 		if !dupe {
// 			continue
// 		}
// 		println(i)
// 	}
// 	a := tj.FileToSlice("input")

// 	for _, l := range a {
// 		println(l)
// 	}

// }
