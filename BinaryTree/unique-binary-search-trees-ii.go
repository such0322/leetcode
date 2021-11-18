package BinaryTree

func GenerateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}
	return helper(1, n)

}
func helper(start, end int) []*TreeNode {
	if start > end {
		return []*TreeNode{nil}
	}

	var alltrees []*TreeNode
	for i := start; i <= end; i++ {
		lefttrees := helper(start, i-1)
		righttrees := helper(i+1, end)
		for _, left := range lefttrees {
			for _, right := range righttrees {
				curtree := &TreeNode{i,nil,nil}
				curtree.Left= left
				curtree.Right = right
				alltrees =append(alltrees, curtree)
			}
		}
	}
	return alltrees
}
