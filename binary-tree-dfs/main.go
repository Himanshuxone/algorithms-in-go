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
	// Explanation: The LCA (Lowest Common ancestor) of nodes 2 and 8 is 6.
	var root *TreeNode
	sl := []int{3, 9, 20, 0, 0, 15, 7}
	for _, value := range sl {
		root = createBinaryTree(root, value)
	}
	resp := make(map[int][]int, 0)
	traverse(root, resp)
}

func createBinaryTree(node *TreeNode, val int) *TreeNode {
	if node == nil {
		if val != 0 {
			// insert nill which means no child
			node = &TreeNode{Val: val}
		}
		return node
	}

	if val > node.Val {
		node.Right = createBinaryTree(node.Right, val)
	} else if val < node.Val {
		node.Left = createBinaryTree(node.Left, val)
	}
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

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	parent := make(map[*TreeNode]*TreeNode)
	DFS(root.Left, p, q, parent)
	DFS(root.Right, p, q, parent)
	fmt.Println(parent)
	return root
}

func DFS(node, p, q *TreeNode, visited map[*TreeNode]*TreeNode) {
	if node == nil {
		return
	}

	if node.Left != nil {
		visited[node.Left] = node
		DFS(node.Left, p, q, visited)
	}
	if node.Right != nil {
		visited[node.Right] = node
		DFS(node.Right, p, q, visited)
	}
}
