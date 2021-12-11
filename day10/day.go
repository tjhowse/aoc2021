package main

import (
	tj "github.com/tjhowse/tjgo"
)

func isOpen(r rune) bool {
	return r == '{' || r == '(' || r == '[' || r == '<'
}

func getType(r rune) int {
	switch r {
	case '{':
		return 1
	case '}':
		return 1
	case '(':
		return 2
	case ')':
		return 2
	case '[':
		return 3
	case ']':
		return 3
	case '<':
		return 4
	case '>':
		return 4

	}
	return 0
}

func score(r rune) int {

	switch r {
	case '}':
		return 1197
	case ')':
		return 3
	case ']':
		return 57
	case '>':
		return 25137

	}
	return 0
}

func score2(r rune) int {

	switch r {
	case '{':
		return 3
	case '(':
		return 1
	case '[':
		return 2
	case '<':
		return 4

	}
	return 0
}

// func flip(r rune) rune {

// 	switch r {
// 	case '{':
// 		return '}'
// 	case '(':
// 		return ')'
// 	case '[':
// 		return ']'
// 	case '<':
// 		return '>'

// 	}
// 	return 0
// }

func main() {

	mainPart1()
	// mainPart2()
}

func mainPart1() {
	a := tj.FileToSlice("input")

	// intSlice := []int{}\
	incomplete := []string{}

	totalScore := 0
	for _, line := range a {
		depth := 0
		// var opener byte
		openers := []byte{}
		corrupt := false
		for i := 0; i < len(line); i++ {
			if isOpen(rune(line[i])) {
				// println("Open:", string(line[i]))
				depth++
				// opener = line[i]
				openers = append(openers, line[i])
			} else {
				// println("Close:", string(line[i]))
				if getType(rune(line[i])) != getType(rune(openers[len(openers)-1])) {
					// println("Mismatch on line ", x)
					// println(line)
					// println(line[:i])
					// println("Closed with", string(rune(line[i])), "expected ", string(rune(openers[len(openers)-1])))
					// println("Score: ", score(rune(line[i])))
					totalScore += score(rune(line[i]))
					corrupt = true
					break
				} else {
					openers = openers[:len(openers)-1]
				}
			}
		}
		if !corrupt {
			incomplete = append(incomplete, line)
			part2Score := 0
			for i := len(openers) - 1; i >= 0; i-- {
				line := openers[i]
				part2Score *= 5
				part2Score += score2(rune(line))
				// println(part2Score)
				// print(string(line))

			}
			// for _, line := range openers {

			// }
			println(part2Score)
		}
	}
	println("total score", totalScore)

	// for _, line := range incomplete {
	// 	println(line)
	// }

}
