package main

import (
	tj "github.com/tjhowse/tjgo"
)

func ox(a []string, index int, flip bool) []string {

	r := make([]string, 0)
	oneCount := []int{}

	for j := 0; j < len(a[0]); j++ {
		oneCount = append(oneCount, 0)
		for _, b := range a {
			if b[j] == byte('1') {
				oneCount[j]++
			}
		}
	}

	var popular byte
	if float64(oneCount[index]) >= float64(len(a))/2 {
		// 1 most poular
		if flip {
			popular = byte('1')
		} else {
			popular = byte('0')
		}
	} else {

		if flip {
			popular = byte('0')
		} else {
			popular = byte('1')
		}
	}
	// println(oneCount[index], " of ", float64(len(a))/2, " are ones")
	// println("Popular: ", string(popular), "index: ", index)
	for _, l := range a {
		if l[index] == popular {
			r = append(r, l)
		}
	}

	return r
}

func main() {
	a := tj.FileToSlice("input")

	for i := 0; i < len(a[0]); i++ {
		a = ox(a, i, true)
		if len(a) == 1 {
			println("FINAL: ", a[0])
			break
		}
	}
	a = tj.FileToSlice("input")

	for i := 0; i < len(a[0]); i++ {
		a = ox(a, i, false)
		if len(a) == 1 {
			println("FINAL: ", a[0])
			break
		}
	}

}

// func main1() {
// 	a := tj.FileToSlice("input")

// 	gamma := 0
// 	epsilon := 0

// 	oneCount := []int{}

// 	for j := 0; j < len(a[0]); j++ {
// 		oneCount = append(oneCount, 0)
// 		for _, b := range a {
// 			if b[j] == byte('1') {
// 				oneCount[j]++
// 			}
// 		}
// 	}

// 	for i, k := range oneCount {
// 		if k >= (len(a) / 2) {
// 			gamma = gamma + int(math.Pow(2, float64(len(a[0])-i)))
// 		} else {
// 			epsilon = epsilon + int(math.Pow(2, float64(len(a[0])-i)))
// 		}
// 	}
// 	gamma /= 2 // HAHAHAH WHY
// 	epsilon /= 2
// 	println(gamma * epsilon)
// 	// println(epsilon)

// }
