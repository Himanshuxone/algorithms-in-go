package main

import "fmt"

func main() {
	s := []int{4, 3, 9, 5, 1, 2, 8, 10}
	var treeNode *Node //initialize empty node struct to start recursion
	fmt.Printf("Create bst of provided array : %d \n", s)
	for i := 0; i < len(s); i++ {
		treeNode = insert(treeNode, s[i]) // the recursive function is initialize
		// before moving to the next iteration the value will traverse the whole tree recursively until found an emtpy node
	}
	intraverse(treeNode)
	fmt.Println()
}

// create a structure named node which also contains pointer to the same node function
type Node struct {
	value      int   // vale of the node which can be called as key or node of binary search tree
	leftChild  *Node // pointer to the node struct
	rightChild *Node // pointer to the right child of teh node
}

// create a function whiclll traverse recursively printing the node value until ffind the leaf node
func intraverse(root *Node) {
	if root == nil {
		return
	}
	intraverse(root.leftChild)
	fmt.Printf("%d", root.value)
	intraverse(root.rightChild)
}

// if we are going to insert a new node into the tree we will pass the node to the function and check for left or right child everytime until null is passed to the root
// where comes the base condition fo the recursion which checks for if the value pass to the insert node is null then we will insert a new node and return its address
// always remember we will traverse the tree for checking node exists and left value exists until we pass null to the function as node
// function contains root argument to check if it is ull or contains left or right and a value which will be inserted as a node if root passed int the function goes null
func insert(root *Node, v int) *Node {
	// check recursively until we reaches to the end of the tree where there is not left or right child and root is empty
	if root == nil {
		root = &Node{v, nil, nil} //insert the value when we finally comes to empty root passed by the left or right condition below
		return root
	}
	// if above condition is not satisfied then the root passed by the function contains value that means we should move down one more node to check and pass null
	// untill there is no left and right child so we keep moving until we pass null to the insert function called recursively in the below conditions

	// check if the root is greater than value or not
	if v < root.value {
		root.leftChild = insert(root.leftChild, v) // pass the current root left node if it exists else null will be passed
		// since there is no value in the root.left which will
		// leads to base condition and root is created which is inserted as left child to the parent root
	} else {
		root.rightChild = insert(root.rightChild, v)
	}
	return root // this value will be returned at the end in golang
}
