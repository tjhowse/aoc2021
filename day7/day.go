package main

import (
	"strings"

	tj "github.com/tjhowse/tjgo"
)

func main() {

	mainPart1()
	// mainPart2()
}

func mainPart1() {
	a := tj.FileToSlice("input")

	intSlice := []int{}
	// blah := []int{}

	s := strings.Split(a[0], ",")
	for _, i := range s {
		intSlice = append(intSlice, tj.Str2int(i))
	}

	for i := 0; i < len(intSlice); i++ {
		println(intSlice[i])
	}
}
