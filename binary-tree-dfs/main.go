package main

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// Input: root = [3,9,20,null,null,15,7]
	// Output: 3
	// Input: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 8
	// Output: 6
	// Explanation: The LCA (Lowest Common ancestor) of nodes 2 and 8 isx 6.
	var root *TreeNode
	resp := make(map[int][]int, 0)
	root = &TreeNode{Val: 6}
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
	traverse(root, resp)
	TreeTraversalDFS(root)
}

func AddNode(node *TreeNode, val int) *TreeNode {
	if val == 0 {
		node = nil
	}
	node = &TreeNode{Val: val}
	return node
}

// using hash maps and recursion
func traverse(node *TreeNode, resp map[int][]int) {
	if node == nil {
		return
	}
	data := make([]int, 0)
	if node.Left != nil {
		data = append(data, node.Left.Val)
	}
	if node.Right != nil {
		data = append(data, node.Right.Val)
	}
	resp[node.Val] = data
	traverse(node.Left, resp)
	traverse(node.Right, resp)
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

func TreeTraversalDFS(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	parent := make(map[*TreeNode]*TreeNode)
	DFS(root.Left, parent)
	DFS(root.Right, parent)
	fmt.Println(parent)
	return root
}

func DFS(node *TreeNode, visited map[*TreeNode]*TreeNode) {
	if node == nil {
		return
	}

	if node.Left != nil {
		visited[node.Left] = node
		DFS(node.Left, visited)
	}
	if node.Right != nil {
		visited[node.Right] = node
		DFS(node.Right, visited)
	}
}
