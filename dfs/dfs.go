package main

import (
	"fmt"
)

func main() {

	graph := map[string][]string{
		"A": []string{"B"},
		"B": []string{"C"},
		"C": []string{"E"},
		"D": []string{"B"},
		"E": []string{"D", "F"},
	}

	fmt.Println("Graph for searching a node:", graph)

	var visitedNodes map[string]int
	visitedNodes = make(map[string]int)
	value := dfsVisit(graph, visitedNodes, "A", "F")
	fmt.Println(value)
}

//  Create a function to visit all the adjacent nodes of all the nodes in the graph
func dfsVisit(graph map[string][]string, visitedNodes map[string]int, node, soughtValue string) bool {

	if node == soughtValue {
		return true
	}

	for _, value := range graph[node] {
		if _, ok := visitedNodes[value]; !ok {
			visitedNodes[value] = 1
			fmt.Println("Nodes vistied:", visitedNodes)
			if dfsVisit(graph, visitedNodes, value, soughtValue) {
				return true
			}
		}
	}
	return false
}
