package questions

func twoSum(nums []int, target int) []int {
	differences := make(map[int]int)
	for i, num := range nums {
		if index, exist := differences[target-num]; exist {
			return []int{index, i}
		}
		differences[num] = i
	}
	return []int{0, 0}
}
