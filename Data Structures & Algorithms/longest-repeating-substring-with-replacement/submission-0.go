func characterReplacement(s string, k int) int {
	count := make(map[byte]int)
	res := 0
	maxFreq := 0
	l := 0

	for r := range len(s){
		count[s[r]]++
		if count[s[r]] > maxFreq{
			maxFreq = count[s[r]]
		}

		//validate the window
		for (r - l + 1) - maxFreq > k{
			count[s[l]]--
			l++
		}
		//check current window size with the prev
		if (r - l + 1) > res{
			res = r - l + 1
		}
	}
	return res
}
