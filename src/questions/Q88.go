package questions

import "log"

func merge(nums1 []int, m int, nums2 []int, n int) {
	if m == 0 {
		log.Print("m is 0")
		nums1 = nums2
		log.Print(nums1)
		return
	}
	if n == 0 {
		log.Print("n is 0")
		log.Print(nums1)
		return
	}
	length := m + n - 1
	mi := m - 1
	ni := n - 1
	for length >= 0 {
		if mi > 0 && nums1[mi] >= nums2[ni] {
			nums1[length] = nums1[mi]
			mi--
		} else if ni > 0 && nums1[mi] < nums2[ni] {
			nums1[length] = nums2[ni]
			ni--
		}
		length--
	}
}
func search(nums []int, target int) int {

	for len(nums) != 0 {
		if nums[len(nums)/2] == target {
			return target
		}
		if nums[len(nums)/2] > target {
			nums = nums[0 : len(nums)/2]
		}
		if nums[len(nums)/2] < target {
			nums = nums[len(nums)/2:]
		}
	}
	return -1

}
