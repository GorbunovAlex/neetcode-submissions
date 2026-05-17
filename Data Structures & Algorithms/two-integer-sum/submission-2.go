func twoSum(nums []int, target int) []int {
	valIdx := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		complement := target - nums[i]
		if _, exists := valIdx[complement]; exists {
			return []int{valIdx[complement], i}
		} else {
			valIdx[nums[i]] = i
		}
	}
	return nil
}
