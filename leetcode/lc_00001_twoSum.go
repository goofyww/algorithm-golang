package leetcode

/**
 * 两数之和
 */
func TwoSum(nums []int, target int) []int {
	mapList := make(map[int]int, len(nums))
	for i, num := range nums {
		x := target - num
		if _, ok := mapList[x]; ok {
			return []int{mapList[x], i}
		}
		mapList[num] = i
	}
	return []int{}
}
