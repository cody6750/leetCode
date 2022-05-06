func search(nums []int, target int) int {
	lower := 0
	upper := len(nums) - 1
	for lower <= upper {
		middle := (upper-lower)/2 + lower
		if nums[middle] == target {
			return middle
		}
		if nums[middle] > target {
			upper = middle - 1
			continue
		}
		if nums[middle] < target {
			lower = middle + 1
			continue
		}
	}
	return -1

}

//  [-1,0,3,5,9,12,13,15,18,29,52], 0
//  [                            ]
//

// 1. Find Middle
// 2. If middle = value, return
// 3. If middle > value, go left
// 4. If middle < value, go right
// Do this until there is no more values