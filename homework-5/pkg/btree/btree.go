package btree

// Doc is a struct for representing page with URL and Title
type Doc struct {
	ID    int
	URL   string
	Title string
}

// Node is a struct representing Binary Tree Structure
type Node struct {
	LeftBranch  *Node
	RightBranch *Node
	*Doc
}

// Find returns binary tree element
func (b *Node) Find(id int) (Node, error) {
	if b.ID == id {
		return *b, nil
	}

	if id < b.ID && b.LeftBranch != nil {
		return b.LeftBranch.Find(id)
	}

	if id > b.ID && b.RightBranch != nil {
		return b.RightBranch.Find(id)
	}
	return Node{}, nil
}

// Insert a node into tree, in place
func (b *Node) Insert(node *Node) {
	if b.Doc == nil {
		*b = *node
		return
	}

	if node.ID == b.ID {
		return
	}

	if node.ID < b.ID {
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
