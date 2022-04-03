package main

import "fmt"

func main() {
	// create a adjacency lists of graph
	maps := map[string]map[int]string{
		"a": {1: "z", 2: "s"},
		"z": {1: "a"},
		"s": {1: "a", 2: "x"},
		"x": {1: "d", 2: "s", 3: "c"},
		"d": {1: "c", 2: "x", 3: "f"},
		"c": {1: "x", 2: "d", 3: "f", 4: "v"},
		"v": {1: "c", 2: "f"},
		"f": {1: "d", 2: "c", 3: "v"},
	}
	// start visitng graph from node s
	visit(maps, "s")
}

func visit(maps map[string]map[int]string, start string) {
	// initialize an empty slice for start
	frontier := []string{start}
	parent := make(map[string]string)
	visited := []string{start}
	counter := 1
	for len(frontier) > 0 {
		next := []string{}
		for _, value := range frontier {
			for _, v := range maps[value] {
				if stringInSlice(v, visited) == false {
					visited = append(visited, v)
					parent[value] = v
					next = append(next, v)
				}
			}
		}
		frontier = next
		counter++
		// print parent of each node
		fmt.Println(parent)
	}
}

func stringInSlice(str string, visited []string) bool {
	for _, v := range visited {
		if str == v {
			return true
		}
	}
	return false
}
