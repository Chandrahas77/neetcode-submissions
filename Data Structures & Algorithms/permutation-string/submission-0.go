func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2){
		return false
	}
	k := len(s1)
	freq1 := [26]int{}
	freq2 := [26]int{}
	for i := range len(s1){
		char1 := s1[i]
		char2 := s2[i]
		freq1[char1 - 'a']++
		freq2[char2 - 'a']++
	}
	if freq1 == freq2{
		return true
	}

	for i := k; i < len(s2); i++{
		freq2[s2[i]- 'a']++
		freq2[s2[i-k]- 'a']--

		if freq1 == freq2{
			return true
		}
	}
	return false
}
