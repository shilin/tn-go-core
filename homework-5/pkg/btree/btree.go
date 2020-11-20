package btree

// Identifier implies Identity, URLstring, Titlestring methods
type Identifier interface {
	Identity() int
	URLstring() string
	Titlestring() string
}

// Node is a struct representing Binary Tree Structure
type Node struct {
	LeftBranch  *Node
	RightBranch *Node
	Identifier
}

// Find returns binary tree element
func (b *Node) Find(id int) (Node, error) {
	if b.Identity() == id {
		return *b, nil
	}

	if id < b.Identity() && b.LeftBranch != nil {
		return b.LeftBranch.Find(id)
	}

	if id > b.Identity() && b.RightBranch != nil {
		return b.RightBranch.Find(id)
	}
	return Node{}, nil
}

// Insert a node into tree, in place
func (b *Node) Insert(node *Node) {

	if b.Identifier == nil {
		b = node
		return
	}

	if node.Identity() == b.Identity() {
		return
	}

	if node.Identity() < b.Identity() {
		if b.LeftBranch != nil {
			b.LeftBranch.Insert(node)
			return
		}
		b.LeftBranch = node
	}

	if b.RightBranch != nil {
		b.RightBranch.Insert(node)
		return
	}
	b.RightBranch = node
}
