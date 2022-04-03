package main

import (
	"fmt"
)

// Node will contains graph node structure
type Node struct {
	name        string
	color       string
	predecessor *Node
	d, f        int
}

var (
	graph = make(map[*Node][]*Node, 0)
	time  int
)

func main() {
	a := &Node{name: "A"}
	b := &Node{name: "B"}
	c := &Node{name: "C"}
	d := &Node{name: "D"}
	e := &Node{name: "E"}
	f := &Node{name: "F"}
	graph = map[*Node][]*Node{
		a: []*Node{},
		b: []*Node{},
		c: []*Node{},
		d: []*Node{},
		e: []*Node{d, f},
		f: []*Node{},
	}
	dfs(graph)
	fmt.Printf("Depth First Search: %+v", graph[e][0])
}

func dfs(graph map[*Node][]*Node) {
	for u := range graph {
		u.color = "white"
		u.predecessor = nil
	}

	time = 0
	// set time to zero initially
	for u := range graph {
		if u.color == "white" {
			dfsVisit(u)
		}
	}
}

//  Create a function to visit all the adjacent nodes of all the nodes in the graph
func dfsVisit(u *Node) {
	time = time + 1
	u.d = time
	u.color = "gray"
	for _, v := range graph[u] {
		if v.color == "white" {
			v.predecessor = u
			dfsVisit(v)
		}
	}
	u.color = "black"
	time = time + 1
	u.f = time
	fmt.Println(u.d, u.f, u.name)
}
