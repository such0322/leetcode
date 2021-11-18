package BinaryTree

func FindTilt(root *TreeNode) (ans int) {
	dfs(root, &ans)
	return
}

func dfs(node *TreeNode, ans *int) int {
	if node == nil {
		return 0
	}
	sumLeft := dfs(node.Left, ans)
	sumRight := dfs(node.Right, ans)
	//坡度上增加2个子集的差值绝对值
	*ans += abs(sumLeft - sumRight)
	//返回左右加自身
	return sumLeft + sumRight + node.Val
}

func abs(n int) int {
	if n < 0 {
		n = -n
	}
	return n
}
