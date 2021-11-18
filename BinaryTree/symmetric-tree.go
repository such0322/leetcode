package BinaryTree

func isSymmetric(root *TreeNode) bool {
	return check(root.Left, root.Right)
}

func check(l, r *TreeNode) bool {
	if l == nil && r == nil {
		return true
	}
	if l == nil || r == nil {
		return false
	}
	return l.Val == r.Val && check(l.Left, r.Right) && check(l.Right, r.Left)
}
