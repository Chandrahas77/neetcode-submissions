func isValid(s string) bool {
	if len(s)%2 != 0{
		return false
	}
    stack := make([]byte,0,len(s))
	matching := map[byte]byte{
		')' : '(',
		'}' : '{',
		']' : '[',
	}
	for i := 0; i < len(s); i++{
		ch := s[i]
		if expectedOpen,isClose := matching[ch]; isClose{
			if len(stack) == 0 || stack[len(stack)-1] != expectedOpen{
				return false
			}
			stack = stack[:len(stack) - 1]
		}else{
			stack = append(stack,ch)
		}
	}
	return len(stack) == 0
}
