func containsDuplicate(nums []int) bool {
	exist := map[int]struct{}{}
	for _, num := range nums {
		if _, exists := exist[num]; exists {
			return true
		} else {
			exist[num] = struct{}{}
		}
	}
	return false
}