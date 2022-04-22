//https://leetcode.com/problems/remove-duplicates-from-sorted-array/
package questions

func removeDuplicates(nums []int) int {
	var indexTracker int
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return 1
	}
	for _, element := range nums {
		if element != nums[indexTracker] {
			indexTracker++
			nums[indexTracker] = element
		}
	}
	return indexTracker + 1
}
