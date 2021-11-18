package BinaryTree

// LevelOrder 广度优先搜索BFS
func LevelOrder(root *TreeNode) [][]int {
	var ret [][]int
	if root == nil {
		return ret
	}
	var arr []*TreeNode
	arr = append(arr, root)
	for len(arr) > 0 {
		var arr2 []*TreeNode
		var r []int
		for j := 0; j < len(arr); j++ {
			r = append(r, arr[j].Val)
			if arr[j].Left != nil {
				arr2 = append(arr2, arr[j].Left)
			}
			if arr[j].Right != nil {
				arr2 = append(arr2, arr[j].Right)
			}
		}
		ret = append(ret, r)
		arr = arr2
	}
	return ret
}
