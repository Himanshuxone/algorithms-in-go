package main

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// Input: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 8
	// Output: 6
	// Explanation: The LCA (Lowest Common ancestor) of nodes 2 and 8 isx 6.
	var root *TreeNode = &TreeNode{Val: 6}
	root.Left = AddNode(root.Left, 2)
	root.Right = AddNode(root.Right, 8)
	root.Left.Left = AddNode(root.Left.Left, 1)
	root.Left.Right = AddNode(root.Left.Right, 4)
	root.Left.Left.Right = AddNode(root.Left.Left.Right, 0)
	root.Left.Left.Left = AddNode(root.Left.Left.Left, 0)
	root.Left.Right.Left = AddNode(root.Left.Right.Left, 7)
	root.Left.Right.Right = AddNode(root.Left.Right.Right, 9)
	root.Right.Left = AddNode(root.Right.Left, 3)
	root.Right.Right = AddNode(root.Right.Right, 5)
	levelOrderGraph(root)
}

func AddNode(node *TreeNode, val int) *TreeNode {
	if val == 0 {
		node = nil
	}
	node = &TreeNode{Val: val}
	return node
}

// Using BFS for directed graph/tree
func levelOrderGraph(root *TreeNode) [][]int {
	res := make([][]int, 0)
	next := make([]*TreeNode, 0)
	visited := make(map[*TreeNode]int, 0)
	level := 0
	frontier := make([]*TreeNode, 0)
	if root != nil {
		res = [][]int{{root.Val}}
		frontier = append(frontier, root)
		visited[root] = level
	} else {
		return [][]int{}
	}
	for len(frontier) != 0 {
		level += 1
		resp := make([]int, 0)
		for _, value := range frontier {
			if value.Left != nil {
				next = append(next, value.Left)
			}
			if value.Right != nil {
				next = append(next, value.Right)
			}
		}
		frontier = next
		for i := 0; i < len(next); i++ {
			if _, ok := visited[next[i]]; !ok {
				visited[next[i]] = level
				resp = append(resp, next[i].Val)
			}
		}
		if len(resp) != 0 {
			res = append(res, resp)
		}
		next = make([]*TreeNode, 0)
	}
	fmt.Println(visited)
	return res
}
