func longestConsecutive(nums []int) int {
	numSet := make(map[int]struct{},len(nums))
	for _,num := range nums{
		numSet[num] = struct{}{}
	}
	longest := 0

	for num := range numSet{
		if _,exists := numSet[num-1]; !exists{
			curNum := num
			curStreak := 1
			for{
				if _,exists := numSet[curNum+1]; exists{
					curNum++
					curStreak++
				} else{
					break
				}
			}
			if curStreak > longest{
			longest = curStreak
			}	
		}
	}
	return longest
}
