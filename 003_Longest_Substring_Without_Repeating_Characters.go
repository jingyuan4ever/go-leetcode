package go_leetcode

func lengthOfLongestSubstring(s string) int {
	max := 0
	currentLength := 0
	tbl := map[int32]int{}
	for i, c := range s{
		index, exists := tbl[c]
		currentLength++
		if exists{
			if currentLength > i - index{
				currentLength = i - index
			}
		}
		if currentLength > max{
			max = currentLength
		}
		tbl[c] = i
	}
	return max
}