func containsDuplicate(nums []int) bool {
	duplicates := map[int]struct{}{}
	sort.Ints(nums)
	for _, num := range nums {
		if _, exist := duplicates[num]; exist {
			return true
		}
		duplicates[num] = struct{}{}
	}
	return false
}