package main

import (
	"math"
	"strings"
	"time"

	tj "github.com/tjhowse/tjgo"
)

func main() {

	mainPart1()
	// mainPart2()
}

func calcFuel(i int) int {
	fuel := 0
	for j := 1; j <= i; j++ {
		fuel += j
	}
	return fuel
}

func mainPart1() {
	a := tj.FileToSlice("input")
	t := time.Now()

	intSlice := []int{}
	// blah := []int{}

	s := strings.Split(a[0], ",")
	for _, i := range s {
		intSlice = append(intSlice, tj.Str2int(i))
	}

	leastFuel := 99999999999999

	for i := 0; i < len(intSlice); i++ {
		fuel := 0
		for _, j := range intSlice {
			// fuel += int(math.Abs(float64(i - j))) // Part 1
			fuel += calcFuel(int(math.Abs(float64(i - j)))) // Part 2

		}
		if fuel < leastFuel {
			leastFuel = fuel
		}
	}
	println(leastFuel)
	t2 := time.Now()

	elapsed := t2.Sub(t)
	println(elapsed.Milliseconds())
}
