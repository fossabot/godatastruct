package rbtree

// This file contains all RB tree iteration methods implementations

// WalkInorder walks tree inorder (left, node, right)
func (tree *rbTree) WalkInorder(action func(Node)) {
	tree.root.walkInorder(func(n *node) { action(n) })
}

func (n *node) walkInorder(action func(*node)) {
	if n != nil && n.key != nil {
		n.left.walkInorder(action)
		action(n)
		n.right.walkInorder(action)
	}
}

// WalkPostorder walks tree postorder (left, right, node)
func (tree *rbTree) WalkPostorder(action func(Node)) {
	tree.root.walkPostorder(func(n *node) { action(n) })
}

func (n *node) walkPostorder(action func(*node)) {
	if n != nil && n.key != nil {
		n.left.walkPostorder(action)
		n.right.walkPostorder(action)
		action(n)
	}
}

// WalkPreorder walks tree preorder (node, left, right)
func (tree *rbTree) WalkPreorder(action func(Node)) {
	tree.root.walkPreorder(func(n *node) { action(n) })
}

func (n *node) walkPreorder(action func(*node)) {
	if n != nil && n.key != nil {
		action(n)
		n.left.walkPreorder(action)
		n.right.walkPreorder(action)
	}
}

// Ascend calls the iterator for every value in the tree until iterator returns false.
func (tree *rbTree) Ascend(iterator NodeIterator) {
	max := tree.Maximum()
	if max == nil {
		return
	}

	min := tree.minimum(tree.root)
	tree.ascend(min, max.Key(), iterator)
}

// AscendRange calls the iterator for every value in the tree within the range
// [from, to], until iterator returns false.
func (tree *rbTree) AscendRange(from, to Comparable, iterator NodeIterator) {
	if tree.root == nil || tree.root.key == nil || to == nil {
		return
	}
	curr, ok := tree.root.search(from)
	if ok {
		tree.ascend(curr, to, iterator)
	}
}

func (tree *rbTree) ascend(n *node, to Comparable, iterator NodeIterator) {
	curr := n
	ok := true
	for ok && curr != nil && curr.key != nil && (curr.key.LessThan(to) || curr.key.EqualTo(to)) {
		ok = iterator(curr)
		if ok {
			curr = tree.successor(curr)
		}
	}
}

// Descend calls the iterator for every value in the tree until iterator returns false.
func (tree *rbTree) Descend(iterator NodeIterator) {
	min := tree.Minimum()
	if min == nil {
		return
	}
	max := tree.maximum(tree.root)
	tree.descend(max, min.Key(), iterator)
}

// DescendRange calls the iterator for every value in the tree within the range
// [from, to], until iterator returns false.
func (tree *rbTree) DescendRange(from, to Comparable, iterator NodeIterator) {
	if tree.root == nil || to == nil {
		return
	}
	curr, ok := tree.root.search(from)
	if ok {
		tree.descend(curr, to, iterator)
	}
}

func (tree *rbTree) descend(n *node, to Comparable, iterator NodeIterator) {
	curr := n
	ok := true
	for ok && curr != nil && curr.key != nil && !curr.key.LessThan(to) {
		ok = iterator(curr)
		if ok {
			curr = tree.predecessor(curr)
		}
	}
}
