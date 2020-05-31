// Package rbtree is a Red-black search binary tree implementation with support ordered statistic on the tree
package rbtree

const (
	// Black RB tree node
	Black = iota

	// Red RB tree node
	Red
)

// RbTree represents red-black tree interface
type RbTree interface {
	Len() int64
	Insert(n Comparable)
	DeleteNode(c Comparable) bool
	WalkInorder(action func(Node))
	WalkPostorder(action func(Node))
	WalkPreorder(action func(Node))
	Ascend(iterator NodeIterator)
	AscendRange(from, to Comparable, iterator NodeIterator)
	Descend(iterator NodeIterator)
	DescendRange(from, to Comparable, iterator NodeIterator)
	Search(value Comparable) (Node, bool)
	Minimum() Node
	Maximum() Node
	OrderStatisticSelect(i int64) (Node, bool)
}

type rbTree struct {
	root *node
	tnil *node
}

// Node represent red-black tree node interface
type Node interface {
	Comparable

	// Subtree size including node itself
	Size() int64
}

// node represent red-black tree node implementation
type node struct {
	key Comparable

	// Subtree size including node itself
	size int64

	color  int
	parent *node
	left   *node
	right  *node
}

// NodeIterator allows callers of Ascend* to iterate in-order over portions of
// the tree.  When this function returns false, iteration will stop and the
// associated Ascend* function will immediately return.
type NodeIterator func(Node) bool

// Comparable defines comparable type interface
type Comparable interface {
	LessThan(y interface{}) bool
	EqualTo(y interface{}) bool
}

// Int is the int type key that can be stored as Node key
type Int int

// String is the string type key that can be stored as Node key
type String string

func (n *node) LessThan(y interface{}) bool {
	switch t := y.(type) {
	case *node:
		return n.key.LessThan(t.key)
	case Comparable:
		return n.key.LessThan(t)
	}

	return n.key == nil
}

func (n *node) EqualTo(y interface{}) bool {
	switch t := y.(type) {
	case *node:
		return n.key.EqualTo(t.key)
	case Comparable:
		return n.key.EqualTo(t)
	}

	return n.key != nil
}

func (n *node) Size() int64 {
	return n.size
}

// LessThan define Comparable interface member for Int
func (x Int) LessThan(y interface{}) bool {
	return x < y.(Int)
}

// EqualTo define Comparable interface member for Int
func (x Int) EqualTo(y interface{}) bool {
	return x == y
}

// LessThan define Comparable interface member for String
func (x *String) LessThan(y interface{}) bool {
	return *x < *(y.(*String))
}

// EqualTo define Comparable interface member for String
func (x *String) EqualTo(y interface{}) bool {
	return *x == *(y.(*String))
}

// GetInt gets int key value from comparable
func GetInt(c Comparable) int {
	return int(c.(Int))
}

// GetString gets string value from Comparable
func GetString(c Comparable) string {
	if c == nil {
		return ""
	}
	return string(*c.(*String))
}

// NewInt creates new Comparable that contains int key
func NewInt(v int) Comparable {
	r := Int(v)
	return r
}

// NewString creates new string Comparable
func NewString(v string) Comparable {
	s := String(v)
	return &s
}

// NewRbTree creates new Red-Black empty tree
func NewRbTree() RbTree {
	return newRbTree()
}

func newRbTree() *rbTree {
	tnil := node{color: Black}
	return &rbTree{tnil: &tnil}
}

// Len returns the number of nodes in the tree.
func (tree *rbTree) Len() int64 {
	if tree.root == nil {
		return 0
	}

	return tree.root.size
}
