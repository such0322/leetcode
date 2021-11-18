package easy

func MoveZeroes(nums []int) {
	if nums == nil || len(nums) == 0 {
		return
	}

	idx := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0  {
			nums[i],nums[idx] = nums[idx],nums[i]
			idx++
		}
	}
}
