package main

import (
	"fmt"
	"leetcode/BinaryTree"
)

func main() {
	var root *BinaryTree.TreeNode
	root = &BinaryTree.TreeNode{}
	root.Val = 4

	root.Left = &BinaryTree.TreeNode{}
	root.Left.Val = 2
	root.Left.Left = &BinaryTree.TreeNode{Val: 3, Left: nil, Right: nil}
	root.Left.Right = &BinaryTree.TreeNode{Val: 5, Left: nil, Right: nil}

	root.Right = &BinaryTree.TreeNode{}
	root.Right.Val = 9
	root.Right.Left = &BinaryTree.TreeNode{Val: 0, Left: nil, Right: nil}
	root.Right.Right = &BinaryTree.TreeNode{Val: 7, Left: nil, Right: nil}
	aa := BinaryTree.FindTilt(root)
	fmt.Println(aa)
}
