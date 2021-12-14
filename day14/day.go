package main

import (
	"math/rand"
	"sort"
	"time"

	tj "github.com/tjhowse/tjgo"
)

func main() {

	rand.Seed(time.Now().UnixNano())
	mainPart1()
}

func mainPart1() {
	// hello := run("input_real")
	run("input")

}

func step(input string, rules map[string]string) string {
	result := string(input[0])
	for i := 0; i < len(input)-1; i++ {
		sub := rules[input[i:i+2]]
		if sub != "" {
			result = result + sub + string(input[i+1])
		}
	}
	return result
}

func count(input string) map[rune]int {
	result := make(map[rune]int)
	for _, letter := range input {
		result[letter]++
	}
	return result
}

func process(pairs map[string]int, rules map[string]string, totalCount *map[string]int) (map[string]int, map[string]int) {
	additions := make(map[string]int)
	removals := make(map[string]int)
	for k, v := range pairs {
		if v <= 0 {
			continue
		}
		for k2, v2 := range rules {
			if k == k2 {
				// We have a match between a pair and rules. Add new pairs for this match.
				removals[k] += v
				newKey1 := string(k[0]) + v2
				newKey2 := v2 + string(k[1])
				// println("Replacing ", k, " with: ", newKey1, ",", newKey2)
				(*totalCount)[v2] += v
				additions[newKey1] += v
				additions[newKey2] += v
			}
		}
	}
	return additions, removals
}

func run(filename string) {
	a := tj.FileToSlice(filename)

	starter := a[0]

	rules := make(map[string]string)

	for i := 2; i < len(a); i++ {
		rules[a[i][:2]] = string(a[i][6])
	}

	// Now we have a set of leter pairs.
	pairs := make(map[string]int)
	totalElementCount := make(map[string]int)

	for _, v := range starter {
		totalElementCount[string(v)]++
	}

	for i := 0; i < len(starter)-1; i++ {
		pairs[starter[i:i+2]]++
	}

	for i := 0; i < 40; i++ {
		println("Step ", i+1)
		add, remove := process(pairs, rules, &totalElementCount)

		for k, v := range add {
			pairs[k] += v
		}
		for k, v := range remove {
			pairs[k] -= v
		}

	}
	sum := 0
	for k, v := range totalElementCount {
		println(string(k), ",", v)
		sum += v
	}

	counts := []int{}
	for _, v := range totalElementCount {
		counts = append(counts, v)
	}
	sort.Ints(counts)
	most := ""
	least := ""
	print("Most abundant: ")
	for k, v := range totalElementCount {
		if counts[len(counts)-1] == v {
			println(string(k))
			most = string(k)
		}
	}
	print("Least abundant: ")
	for k, v := range totalElementCount {
		if counts[0] == v {
			println(string(k))
			least = string(k)
		}
	}
	println("Total: ", totalElementCount[most]-totalElementCount[least])

	return

	// result := starter
	// for i := 0; i < 15; i++ {
	// 	// println("Step: ", i)
	// 	result = step(result, rules)
	// 	counted := count(result)
	// 	// println("#F: ", counted['F'])
	// 	// println("#N: ", counted['N'])
	// 	print(counted['F'], ",")
	// 	println(counted['N'])
	// }

	// counts := []int{}
	// counted := count(result)
	// for _, v := range counted {
	// 	counts = append(counts, v)
	// }
	// sort.Ints(counts)
	// print("Most abundant: ")
	// for k, v := range counted {
	// 	if counts[len(counts)-1] == v {
	// 		println(string(k))
	// 	}
	// }
	// print("Least abundant: ")
	// for k, v := range counted {
	// 	if counts[0] == v {
	// 		println(string(k))
	// 	}
	// }
	// println("Result: ", counts[len(counts)-1]-counts[0])

	// 10 steps:
	// K:  S v: 1343
	// K:  C v: 1669
	// K:  V v: 2772
	// K:  H v: 1913
	// K:  F v: 3296
	// K:  N v: 706
	// K:  P v: 2004
	// K:  B v: 1482
	// K:  O v: 1932
	// K:  K v: 2340

}
