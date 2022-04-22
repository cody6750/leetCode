package questions

func maxSubArray(nums []int) int {
	mx, cur := nums[0], nums[0]
	for _, n := range nums[1:] {
		cur = max(n, cur+n)
		mx = max(mx, cur)
	}
	return mx
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
