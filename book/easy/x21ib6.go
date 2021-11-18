package easy

func singleNumber(nums []int) int {
	rs := 0
	for _, num := range nums {
		rs ^= num
	}
	return rs
}
