package main

import (
	"math/rand"
	"strings"
	"time"
	"unicode"

	tj "github.com/tjhowse/tjgo"
)

func nextCaves(current string, paths [][2]string, visitCount map[string]int) []string {
	result := []string{}

	for _, path := range paths {
		if path[0] == current {
			if path[1] == "start" {
				continue
			}
			if unicode.IsUpper(rune(path[1][0])) || visitCount[path[1]] == 0 {
				// We can't visit a lowercase cave more than once
				result = append(result, path[1])
			}
		}
	}
	return result
}

func checkDupeRoute(route string, routes []string) bool {
	for _, c := range routes {
		if route == c {
			return true
		}
	}
	return false
}

func removeDestination(dest string, paths [][2]string) [][2]string {
	result := [][2]string{}
	for _, line := range paths {
		if (line[1] == dest) && dest != "start" && dest != "end" {
			// println("Removing ", dest)
			// println("line[0] ", line[0])
			// println("line[1] ", line[1])
			continue
		}
		result = append(result, line)
	}
	return result
}

func main() {

	// mainPart1()
	mainPart2()
}

func doRoute(paths [][2]string) string {

	visitCount := make(map[string]int)
	route := "start"
	current := "start"
	for current != "end" {
		// println("We're at ", current, " next options are:")
		nextOptions := nextCaves(current, paths, visitCount)
		// for _, path := range nextOptions {
		// 	println(path)
		// }
		if len(nextOptions) > 0 {
			choice := nextOptions[rand.Intn(len(nextOptions))]
			// println("We chose ", choice)
			route += "," + choice
			visitCount[choice]++
			if !unicode.IsUpper(rune(choice[0])) {
				// We can only visit this destination once, remove it from future paths.
				paths = removeDestination(choice, paths)
			}
			current = choice
			if choice == "end" {
				break
			}
		} else {
			// println("Dead end", route)
			break
		}
	}
	return route
}

func doRouteBetter(route string, start string, paths [][2]string) []string {

	visitCount := make(map[string]int)
	routes := []string{}
	route += "," + start
	if !unicode.IsUpper(rune(start[0])) {
		// We can only visit this destination once, remove it from future paths.
		paths = removeDestination(start, paths)
	}
	if start != "end" {
		// println("We're at ", current, " next options are:")
		nextOptions := nextCaves(start, paths, visitCount)
		if len(nextOptions) > 0 {
			for _, choice := range nextOptions {
				routes = append(routes, doRouteBetter(route, choice, paths)...)
			}
		}
	} else {
		routes = append(routes, route)
	}
	return routes
}

func doRouteBetterPart2(route string, start string, paths [][2]string, special string) []string {

	visitCount := make(map[string]int)
	routes := []string{}
	route += "," + start
	if !unicode.IsUpper(rune(start[0])) {
		if start == "start" || start == "end" {
			paths = removeDestination(start, paths)
		} else {
			if special == "" {
				special = start
			} else if start != special {
				paths = removeDestination(start, paths)
			} else {
				special = "that's enough"
				// paths = removeDestination(start, paths)
			}
		}
	}
	if start != "end" {
		// println("We're at ", current, " next options are:")
		nextOptions := nextCaves(start, paths, visitCount)
		if len(nextOptions) > 0 {
			for _, choice := range nextOptions {
				routes = append(routes, doRouteBetterPart2(route, choice, paths, special)...)
			}
		}
	} else {
		routes = append(routes, route)
	}
	return routes
}

func mainPart1() {
	a := tj.FileToSlice("input")
	rand.Seed(time.Now().UnixNano())

	caveSet := make(map[string]bool)
	paths := [][2]string{}

	// routesThroughCave := []string{}

	for _, line := range a {
		split := strings.Split(line, "-")
		caveSet[split[0]] = true
		caveSet[split[1]] = true
		paths = append(paths, [2]string{split[0], split[1]})
		paths = append(paths, [2]string{split[1], split[0]})
	}

	routesThroughCave := doRouteBetter("", "start", paths)

	count := 0
	for _, route := range routesThroughCave {
		if route[len(route)-3:] == "end" {
			println(route)
			count++
		}
	}
	println("Unique routes: ", count)

}

func mainPart2() {
	a := tj.FileToSlice("input")
	rand.Seed(time.Now().UnixNano())

	caveSet := make(map[string]bool)
	paths := [][2]string{}

	// routesThroughCave := []string{}

	for _, line := range a {
		split := strings.Split(line, "-")
		caveSet[split[0]] = true
		caveSet[split[1]] = true
		paths = append(paths, [2]string{split[0], split[1]})
		paths = append(paths, [2]string{split[1], split[0]})
	}

	lowerCaseCaves := []string{}
	for key, _ := range caveSet {
		if !unicode.IsUpper(rune(key[0])) {
			if key == "start" || key == "end" {
				continue
			}
			lowerCaseCaves = append(lowerCaseCaves, key)
			println("lower case cave: ", key)
		}
	}
	totalRoutes := []string{}
	for _, lowerCaseCave := range lowerCaseCaves {
		for _, newRoute := range doRouteBetterPart2("", "start", paths, lowerCaseCave) {
			if !checkDupeRoute(newRoute, totalRoutes) {
				totalRoutes = append(totalRoutes, newRoute)
			}
		}
	}

	count := 0
	for _, route := range totalRoutes {
		if route[len(route)-3:] == "end" {
			println(route)
			count++
		}
	}
	println("Unique routes: ", count)

}
