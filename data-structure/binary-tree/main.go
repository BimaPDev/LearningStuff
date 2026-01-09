package main

import "fmt"

//node = components of the binary search tree

type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

// Insert
func (n *Node) Insert(k int) {
	if n.Key < k {
		// move right
		if n.Right == nil {
			n.Right = &Node{Key: k}
		} else {
			n.Right.Insert(k)
		}
	} else if n.Key > k {
		// move left
		if n.Left == nil {
			n.Left = &Node{Key: k}
		} else {
			n.Left.Insert(k)
		}
	}
}

// search
func (n *Node) Search(k int) bool {
	if n == nil {
		return false
	}
	if k < n.Key {
		return n.Left.Search(k)
	}
	if k > n.Key {
		return n.Right.Search(k)
	}
	return true
}

func main() {
	var tree *Node
	tree.Insert(100)
	tree.Insert(20)
	tree.Insert(50)
	tree.Insert(60)
	tree.Insert(30)

	fmt.Println("Search 30:", tree.Search(30))
	fmt.Println("Search 999:", tree.Search(999))
}
