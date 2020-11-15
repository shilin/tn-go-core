// package main

// import (
// 	"fmt"
// )

package binarytree

// Identifier implies ID method
type Identifier interface {
	Identity() int
	// setID(int)
}

// BinaryTree is a struct representing Binary Tree Structure
type BinaryTree struct {
	LeftBranch  *BinaryTree
	RightBranch *BinaryTree
	Identifier
}

// Doc1 for page
type Doc1 struct {
	ID    int
	URL   string
	Title string
}

// Identity returns unique identity
func (d *Doc1) Identity() int {
	return d.ID
}

// Find returns binary tree element
func (b *BinaryTree) Find(id int) (BinaryTree, error) {
	if b.Identity() == id {
		// fmt.Println("in Find and equals")
		return *b, nil
	}

	if id < b.Identity() && b.LeftBranch != nil {
		// fmt.Println("in Find and left not nil")
		return b.LeftBranch.Find(id)
	}

	if id > b.Identity() && b.RightBranch != nil {
		// fmt.Println("in Find and right not nil")
		return b.RightBranch.Find(id)
	}
	return BinaryTree{}, nil
}

// Insert a node into tree, in place
func (b *BinaryTree) Insert(node *BinaryTree) {
	if node.Identity() == b.Identity() {
		return
	}

	if node.Identity() < b.Identity() {
		// fmt.Println("in insert")
		if b.LeftBranch != nil {
			// fmt.Println("in insert and leftBranch not nil")
			// fmt.Printf("in insert and leftBranch Identity %d", b.LeftBranch.Identity())
			// fmt.Println("")
			b.LeftBranch.Insert(node)
			return
		}
		// fmt.Println("in insert and leftBranch is nil")
		b.LeftBranch = node
		// fmt.Println("in insert and leftBranch is nil")
	}

	if b.RightBranch != nil {
		b.RightBranch.Insert(node)
		return
	}
	b.RightBranch = node
}

// func main() {
// 	d10 := Doc1{ID: 10}
// 	d15 := Doc1{ID: 15}
// 	d25 := Doc1{ID: 25}
// 	d5 := Doc1{ID: 5}
// 	d7 := Doc1{ID: 7}
// 	d8 := Doc1{ID: 8}

// 	root := BinaryTree{Identifier: &d10}

// 	root.LeftBranch = &BinaryTree{Identifier: &d5}
// 	root.RightBranch = &BinaryTree{Identifier: &d15}
// 	root.Insert(&BinaryTree{Identifier: &d7})
// 	root.Insert(&BinaryTree{Identifier: &d25})
// 	root.Insert(&BinaryTree{Identifier: &d8})
// 	fmt.Println(root.LeftBranch.RightBranch.Identity())
// 	// fmt.Println(root.LeftBranch.Identity())
// 	fmt.Println(root.LeftBranch.RightBranch.RightBranch.Identity())
// 	// fmt.Println(root.RightBranch.RightBranch.Identity())
// 	el, err := root.Find(25)
// 	if err == nil {
// 		fmt.Println(el.Identity())
// 	}

// }
