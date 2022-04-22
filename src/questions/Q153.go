func findMin(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	curr := nums[0]
	for i := 1; i < len(nums); i++ {
		if curr > nums[i] {
			return nums[i]
		}
	}
	return curr
}