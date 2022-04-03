package main

import (
	"fmt"
)

type Node struct {
	value int
	left  *Node
	right *Node
}

func main() {
	sl := []int{9, 4, 3, 6, 1, 2, 10, 5, 7, 8, 16, 12, 17, 14}
	var treeNode *Node
	for _, value := range sl {
		treeNode = insertNode(treeNode, value)
	}
	// fmt.Printf("%v\n", sl)
	// fmt.Printf("TreeNode :- %v\n", treeNode.left.right.left)
	intraverse(treeNode)
}

// create a function which will traverse recursively printing the node value until find the leaf node
func intraverse(root *Node) {
	if root == nil {
		return
	}
	intraverse(root.left)
	fmt.Printf("%d\n", root.value)
	intraverse(root.right)
}

func insertNode(root *Node, val int) *Node {
	if root == nil {
		root = &Node{val, nil, nil}
		return root
	}

	if val < root.value {
		root.left = insertNode(root.left, val)
	} else {
		root.right = insertNode(root.right, val)
	}
	return root
}
