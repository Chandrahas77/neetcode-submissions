func lengthOfLongestSubstring(s string) int {
	charSet := make(map[byte]struct{})
	maxLen := 0
	left := 0
	for right := range len(s){
		char := s[right]
		for {
			if _,exists := charSet[char]; !exists{
				break
			}
			delete(charSet,s[left])
			left++
		}
		charSet[char] = struct{}{}
		windowLen := right - left + 1
		if windowLen > maxLen{
			maxLen = windowLen
		}
	}
	return maxLen
}
