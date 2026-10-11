func containsNearbyDuplicate(nums []int, k int) bool {
	seen := make(map[int]int,len(nums))

	for i := range nums{
		lastSeen,exists := seen[nums[i]]
		if exists && i - lastSeen <= k{
			return true
		}else{
			seen[nums[i]] = i
		}
	}
	return false
}
