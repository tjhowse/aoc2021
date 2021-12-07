package main

import (
	"strings"

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
	// mainPart2()
}

func mainPart2() {
	a := tj.FileToSlice("input")

	f := []int{}

	s := strings.Split(a[0], ",")
	for _, i := range s {
		f = append(f, tj.Str2int(i))
	}

	days := 80
	total := len(f)
	for i := 0; i < days; i++ {
		// for j := 0;

	}
	println(total)
}

func mainPart1() {
	a := tj.FileToSlice("input")

	f := []fish{}
	simple := []int{}

	s := strings.Split(a[0], ",")
	for _, i := range s {
		f = append(f, fish{c: tj.Str2int(i)})
		simple = append(simple, tj.Str2int(i))
	}

	// for _, i := range f {
	// 	println(i.c)
	// }

	lenSeq := []int{}
	// prevLen := len(f)

	for i := 0; i < 80; i++ {
		// for i := 0; i < 256; i++ {
		newFish := []fish{}
		for j := 0; j < len(f); j++ {
			if f[j].tick() {
				newFish = append(newFish, fish{c: 8})
			}
		}
		f = append(f, newFish...)
		// print("Day: ", i, " : ")
		// for k := 0; k < len(f); k++ {
		// 	print(f[k].c, ",")
		// }
		// println()
		// lenSeq = append(lenSeq, len(f)-prevLen)
		lenSeq = append(lenSeq, len(f))
		// println("Len: ", len(f), "step:", len(f)-prevLen)
		// prevLen = len(f)
	}
	println("Length:", len(f))

	b := 8
	q := 6

	// for i := b + 1; i < len(lenSeq); i += 1 {
	// 	print((lenSeq[i] - lenSeq[i-1]), " , ", (lenSeq[i-b]-lenSeq[i-b-1])+(lenSeq[i-q]-lenSeq[i-q-1]))
	// 	println()
	// }
	for i := len(lenSeq); i < 256; i += 1 {
		lenSeq = append(lenSeq, lenSeq[i-1]+(lenSeq[i-b-1]-lenSeq[i-b-2])+(lenSeq[i-q-1]-lenSeq[i-q-2]))
		// print((lenSeq[i] - lenSeq[i-1]), " , ", (lenSeq[i-b]-lenSeq[i-b-1])+(lenSeq[i-q]-lenSeq[i-q-1]))
		// println()
	}
	println(lenSeq[len(lenSeq)-1])
	// println("-1:", lenSeq[len(lenSeq)-1])
	// println("-2:", lenSeq[len(lenSeq)-2])
	// println("-3:", lenSeq[len(lenSeq)-3])

	// numbers := []int{}

	// for i, n := range a {
	// 	if i < 10 {
	// 		s := strings.Split(n, " ")
	// 		// println(s)
	// 		fmt.Printf("%v\n", s)

	// 	}
	// 	// println(n)
	// 	// m.line(n)
	// }
	// m.line(a[0])
	// m.line(a[1])

	// m.draw()
	// println(m.count())

}
