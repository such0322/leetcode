package BinaryTree

import "fmt"

func RecoverTree(root *TreeNode) {
	var nums []int
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		nums = append(nums, node.Val)
		inorder(node.Right)
	}
	inorder(root)
	fmt.Println(nums)
	x, y := findtwosp(nums)
	fmt.Println(x, y)
	recover(root, 2, x, y)
	fmt.Println(root)
}

func findtwosp(nums []int) (x, y int) {
	index1, index2 := -1, -1
	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1] <= nums[i] {
			index2 = i + 1
			if index1 == -1 {
				index1 = i
			} else {
				break
			}

		}
	}
	return nums[index1], nums[index2]
}

func recover(root *TreeNode, count, x, y int) {
	if root == nil {
		return
	}
	if root.Val == x || root.Val == y {
		if root.Val == x {
			root.Val = y
		} else {
			root.Val = x
		}
		count--
		if count == 0 {
			return
		}
	}
	recover(root.Left, count, x, y)
	recover(root.Right, count, x, y)
}
