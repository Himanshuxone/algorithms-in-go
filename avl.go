package main

import "fmt"

func main() {
	fmt.Println("Welcome to AVL trees")
	s := []int{12, 8, 18, 5, 11, 17, 4, 7, 2}
	fmt.Println(s)
	var treeNode *Node
	for i := 0; i < len(s); i++ {
		treeNode = insert(treeNode, s[i])
	}
	inorder(treeNode)
}

type Node struct {
	key         int
	h           int
	left, right *Node
}

func insert(root *Node, v int) *Node {
	if root == nil {
		root = &Node{v, 0, nil, nil}
		root.h = max(height(root.left), height(root.right)) + 1
		return root
	}
	if v < root.key {
		root.left = insert(root.left, v)
		root.h = max(height(root.left), height(root.right)) + 1
	}
	if v > root.key {
		root.right = insert(root.right, v)
		root.h = max(height(root.left), height(root.right)) + 1
	}
	root.h = max(height(root.left), height(root.right)) + 1
	return root
}

func inorder(root *Node) {
	if root == nil {
		return
	}
	inorder(root.left)
	fmt.Printf("{%d,%d}\n", root.key, root.h)
	inorder(root.right)
}

func height(root *Node) int {
	if root != nil {
		return root.h
	}
	return -1
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
