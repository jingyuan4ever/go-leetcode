package go_leetcode

func twoSum(nums []int, target int) []int {
	m := map[int]int{}
	l := len(nums)
	for i:=0;i<l;i++{
		if v, exists := m[target - nums[i]]; exists{
			return []int{v, i}
		}
		m[nums[i]] = i
	}
	return []int{-1,-1}
}
