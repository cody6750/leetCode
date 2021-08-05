package questions

var (
	complements map[string]int
	finalList   []int
	testCases   []int
)

func TwoSum(nums []int, target int) []int {
	var complements = make(map[int]int)
	var finalList []int
	if len(nums) == 2 {
		finalList = append(finalList, 0)
		finalList = append(finalList, 1)
		return finalList
	}
	for index, num := range nums {
		complements[num] = index
	}
	for index, num := range nums {
		_, ok := complements[target-num]
		if ok == true && complements[target-num] != index {
			finalList = append(finalList, complements[target-num])
			finalList = append(finalList, index)
			return finalList
		}
	}
	return finalList
}
