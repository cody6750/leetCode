func productExceptSelf(nums []int) []int {
	n := len(nums)
	leftArray := make([]int, n)
	rightArray := make([]int, n)
	result := make([]int, n)

	for i, _ := range nums {
		if i == 0 {
			leftArray[0] = 1
			rightArray[len(nums)-1] = 1
			continue
		}
		leftArray[i] = nums[i-1] * leftArray[i-1]
		rightArray[len(nums)-1-i] = nums[len(nums)-i] * rightArray[len(nums)-i]
	}

	for i, _ := range nums {
		result[i] = leftArray[i] * rightArray[i]
	}

	return result
}


