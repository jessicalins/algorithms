package main

import (
	"fmt"

	. "github.com/wwgberlin/algorithms/tree"
)

//Binary search trees satisfy the rule that every node on the left of node n
//has a value little than the value of node n, and consequentially every node on
//the right has a value larger than or equal to the value of n.

//Recursively find the right leaf to insert a new Node with given value
//into the binary search tree
//if node is nil it's impossible to perform insert,
//you may return without doing anything, or panic.
func Insert(node *Node, value int) {
	// worst case: O(log n) - height of the tree
	if node == nil {
		panic("Could not plant node!")
	}
	if value < node.Value {
		if node.Left != nil {
			Insert(node.Left, value)
		} else {
			node.Left = &Node{Value: value}
		}
	} else {
		if node.Right != nil {
			Insert(node.Right, value)
		} else {
			node.Right = &Node{Value: value}
		}
	}
}

//Recursively print the items in the tree from small to large
//our base recursion check is if the given node (root of the subtree) is nil
//if not nil, we recursively call PrintSorted with the left node, then print the current node
//then call PrintSorted with the right node.
func PrintSorted(node *Node) {
	if node != nil {
		PrintSorted(node.Left)
		fmt.Println(node.Value)
		PrintSorted(node.Right)
	}
}

func main() {
	root := Node{Value: 5}
	Insert(&root, 4)
	Insert(&root, 8)
	Insert(&root, 7)
	Insert(&root, 6)
	Insert(&root, 10)
	Insert(&root, 2)
	Insert(&root, 3)
	Insert(&root, 1)
	Insert(&root, 2)
	Insert(&root, 5)
	Insert(&root, 11)
	Insert(&root, 0)
	fmt.Println(PrintTree(&root))
	PrintSorted(&root)
	fmt.Print("\n")
}
