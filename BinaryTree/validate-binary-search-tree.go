package BinaryTree

import "math"

func isValidBST(root *TreeNode) bool {
	return helper2(root, math.MinInt64, math.MaxInt64)
}

//子节点的值 左子节点>min <val, 右子节点>val <max
func helper2(node *TreeNode, left, right int) bool {
	if node == nil {
		return true
	}
	//如果这里有条件正确了，那就不是二叉树了，比如左子节点应该比right小   右子节点应该比left大
	if node.Val <= left || node.Val >= right {
		return false
	}
	return helper2(node.Left,left,node.Val) && helper2(node.Right,node.Val,right)
}
