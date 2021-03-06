// Package rbtree is a Red-black search binary tree implementation with support ordered statistic on the tree
package rbtree

import "fmt"

const (
	// Black RB tree node
	Black = iota

	// Red RB tree node
	Red
)

// RbTree represents red-black tree interface
type RbTree interface {
	// Len returns the number of nodes in the tree.
	Len() int64

	// Insert inserts new node into Red-Black tree. Creates Root if tree is empty
	Insert(n Comparable)

	// DeleteNode searches and deletes node with key value specified from Red-black tree
	// It returns true if node was successfully deleted otherwise false
	DeleteNode(c Comparable) bool

	// DeleteAllNodes searches and deletes all found nodes with key value specified from Red-black tree
	// It returns true if nodes was successfully deleted otherwise false
	DeleteAllNodes(c Comparable) bool

	// WalkInorder walks tree inorder (left, node, right)
	WalkInorder(action func(Node))

	// WalkPostorder walks tree postorder (left, right, node)
	WalkPostorder(action func(Node))

	// WalkPreorder walks tree preorder (node, left, right)
	WalkPreorder(action func(Node))

	// Ascend calls the iterator for every value in the tree until iterator returns false.
	Ascend(iterator NodeIterator)

	// AscendRange calls the iterator for every value in the tree within the range
	// [from, to], until iterator returns false.
	AscendRange(from, to Comparable, iterator NodeIterator)

	// Descend calls the iterator for every value in the tree until iterator returns false.
	Descend(iterator NodeIterator)

	// DescendRange calls the iterator for every value in the tree within the range
	// [from, to], until iterator returns false.
	DescendRange(from, to Comparable, iterator NodeIterator)

	// Search searches value specified within search tree
	Search(value Comparable) (Node, bool)

	// Minimum gets tree's min element
	Minimum() Node

	// Maximum gets tree's max element
	Maximum() Node

	// OrderStatisticSelect gets i element from subtree
	OrderStatisticSelect(i int64) (Node, bool)
}

type rbTree struct {
	root *node
	tnil *node
}

// Node represent red-black tree node interface
type Node interface {
	// Subtree size including node itself
	Size() int64

	// Key gets node's key
	Key() Comparable

	// Successor gets node's successor
	Successor() Node

	// Predecessor gets node's predecessor
	Predecessor() Node
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
	fmt.Stringer
	LessThan(y interface{}) bool
	EqualTo(y interface{}) bool
}

// Int is the int type key that can be stored as Node key
type Int int

// String is the string type key that can be stored as Node key
type String string

func (n *node) Size() int64 {
	return n.size
}

func (n *node) String() string {
	return n.key.String()
}

func (n *node) Key() Comparable {
	return n.key
}

func (n *node) isNil() bool {
	return n == nil || n.key == nil
}

// LessThan define Comparable interface member for Int
func (x Int) LessThan(y interface{}) bool {
	return x < y.(Int)
}

// EqualTo define Comparable interface member for Int
func (x Int) EqualTo(y interface{}) bool {
	return x == y
}

func (x Int) String() string {
	return fmt.Sprintf("%d", x)
}

// LessThan define Comparable interface member for String
func (x *String) LessThan(y interface{}) bool {
	return *x < *(y.(*String))
}

// EqualTo define Comparable interface member for String
func (x *String) EqualTo(y interface{}) bool {
	return *x == *(y.(*String))
}

func (x *String) String() string {
	return string(*x)
}

// GetInt gets int key value from comparable
func GetInt(c Comparable) int {
	return int(c.(Int))
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
	if tree.root == nil || tree.root == tree.tnil {
		return 0
	}

	return tree.root.size
}
