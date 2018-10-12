// Package rbtree is a Red-black search binary tree implementation with support ordered statistic on the tree
package rbtree

const (
	// Black RB tree node
	Black = iota

	// Red RB tree node
	Red
)

// RbTree represents red-black tree structure
type RbTree struct {
	Root *Node
	tnil *Node
}

// Node represent red-black tree node
type Node struct {
	// Node key (data)
	Key *Comparable

	// Subtree size including node itself
	Size int64

	color  int
	parent *Node
	left   *Node
	right  *Node
}

// Comparable defines comparable type interface
type Comparable interface {
	LessThan(y interface{}) bool
	EqualTo(y interface{}) bool
}

// Int is the int type key that can be stored as Node key
type Int int

// String is the string type key that can be stored as Node key
type String string

// LessThan define Comparable interface member for Int
func (x Int) LessThan(y interface{}) bool {
	return x < y.(Int)
}

// EqualTo define Comparable interface member for Int
func (x Int) EqualTo(y interface{}) bool {
	return x == y.(Int)
}

// LessThan define Comparable interface member for String
func (x String) LessThan(y interface{}) bool {
	return x < y.(String)
}

// EqualTo define Comparable interface member for String
func (x String) EqualTo(y interface{}) bool {
	return x == y.(String)
}

// GetIntKey gets int key value from tree node
func (n *Node) GetIntKey() int {
	return int((*n.Key).(Int))
}

// GetStringKey gets string key value from tree node
func (n *Node) GetStringKey() string {
	return string((*n.Key).(String))
}

// NewIntKey creates new int key to be stores as tree node key
func NewIntKey(v int) *Comparable {
	var r Comparable
	r = Int(v)
	return &r
}

// NewStringKey creates new string key to be stores as tree node key
func NewStringKey(v string) *Comparable {
	var r Comparable
	r = String(v)
	return &r
}

// NewRbTree creates new Red-Black empty tree
func NewRbTree() *RbTree {
	tnil := Node{color: Black}
	return &RbTree{tnil: &tnil}
}

// NewNode creates new node
func NewNode(si Comparable) *Node {
	return &Node{Key: &si}
}

// WalkInorder walks subtree inorder (left, node, right)
func (n *Node) WalkInorder(action func(*Node)) {
	if n != nil && n.Key != nil {
		n.left.WalkInorder(action)
		action(n)
		n.right.WalkInorder(action)
	}
}

// WalkInorder walks tree inorder (left, node, right)
func (tree *RbTree) WalkInorder(action func(*Node)) {
	tree.Root.WalkInorder(action)
}

// WalkPreorder walks subtree preorder (node, left, right)
func (n *Node) WalkPreorder(action func(*Node)) {
	if n != nil && n.Key != nil {
		action(n)
		n.left.WalkPreorder(action)
		n.right.WalkPreorder(action)
	}
}

// WalkPreorder walks tree preorder (node, left, right)
func (tree *RbTree) WalkPreorder(action func(*Node)) {
	tree.Root.WalkPreorder(action)
}

// Len returns the number of nodes in the tree.
func (tree *RbTree) Len() int64 {
	if tree.Root == nil {
		return 0
	}

	return tree.Root.Size
}
