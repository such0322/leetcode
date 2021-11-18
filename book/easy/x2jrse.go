package easy

func TwoSum(nums []int, target int) []int {
	im := make(map[int]int)
	for i, num := range nums {
		if _, ok := im[num]; !ok {
			im[target-num] = i
		} else {
			return []int{im[num], i}
		}
	}
	return nil
}
