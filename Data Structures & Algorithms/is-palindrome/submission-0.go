// import "unicode"
func isPalindrome(s string) bool {
	runes := []rune(s)
	left, right := 0, len(runes)-1
	for left < right{
		if !isAlphaNumeric(runes[left]){
			left++
		} else if !isAlphaNumeric(runes[right]){
			right--
		} else{
			if unicode.ToLower(runes[left]) != unicode.ToLower(runes[right]){
				return false
			}
			left++
			right--
		}
	}
	return true
}

func isAlphaNumeric(r rune) bool{
	return unicode.IsLetter(r) || unicode.IsDigit(r)
}
