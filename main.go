package main

import (
	"fmt"
	"leetcode/BinaryTree"
)

func main() {
	var root *BinaryTree.TreeNode
	root = &BinaryTree.TreeNode{}
	root.Val = 3
	root.Left = &BinaryTree.TreeNode{Val: 9, Left: nil, Right: nil}
	root.Left.Left = &BinaryTree.TreeNode{Val: 6, Left: nil, Right: nil}

	root.Right = &BinaryTree.TreeNode{}
	root.Right.Val = 20
	root.Right.Left = &BinaryTree.TreeNode{Val: 15, Left: nil, Right: nil}
	root.Right.Right = &BinaryTree.TreeNode{Val: 7, Left: nil, Right: nil}
	aa := BinaryTree.LevelOrder(root)
	fmt.Println(aa)
}
