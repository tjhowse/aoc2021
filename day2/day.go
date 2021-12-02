package main

import (
	"strings"

	tj "github.com/tjhowse/tjgo"
)

func main() {
	a := tj.FileToSlice("input")

	// b := []int{}
	intSlice := []int{}
	stringSlice := []string{}

	for _, d := range a {
		command := strings.Split(d, " ")
		stringSlice = append(stringSlice, command[0])
		intSlice = append(intSlice, tj.Str2int(command[1]))
	}

	distance := 0
	depth := 0
	aim := 0

	for i := 0; i < len(stringSlice); i++ {
		switch stringSlice[i] {
		case "forward":
			distance += intSlice[i]
			depth += intSlice[i] * aim // Part 2
		case "down":
			aim += intSlice[i] // Part 2
			// depth += intSlice[i] // Part1
		case "up":
			aim -= intSlice[i] // Part 2
			// depth -= intSlice[i] // Part1
		}
	}
	println("Depth: ", depth, "Distance: ", distance, "Aim: ", aim)
	println("tot: ", depth*distance)

}
