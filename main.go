package main

import "fmt"

// TreeNode is a binary search tree node.
type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// BSTIterator yields a BST's values in ascending order, one at a time, with
// O(h) memory: the stack holds the left spine of the subtree still to come.
type BSTIterator struct {
	stack []*TreeNode
}

// NewBSTIterator positions the iterator before the smallest value.
func NewBSTIterator(root *TreeNode) *BSTIterator {
	it := &BSTIterator{}
	it.pushLeft(root)
	return it
}

func (it *BSTIterator) pushLeft(n *TreeNode) {
	for ; n != nil; n = n.Left {
		it.stack = append(it.stack, n)
	}
}

// HasNext reports whether another value remains.
func (it *BSTIterator) HasNext() bool { return len(it.stack) > 0 }

// Next returns the next value in ascending order. Amortized O(1).
func (it *BSTIterator) Next() int {
	n := it.stack[len(it.stack)-1]
	it.stack = it.stack[:len(it.stack)-1]
	it.pushLeft(n.Right)
	return n.Val
}

func tn(v int, l, r *TreeNode) *TreeNode { return &TreeNode{Val: v, Left: l, Right: r} }

func main() {
	root := tn(7, tn(3, nil, nil), tn(15, tn(9, nil, nil), tn(20, nil, nil)))
	it := NewBSTIterator(root)
	for it.HasNext() {
		fmt.Print(it.Next(), " ")
	}
	fmt.Println()
}
