func evalRPN(tokens []string) int {
	stack := []int{}
	for _,token := range tokens{
		switch token{
			case "+":
			right := stack[len(stack)-1]
			left := stack[len(stack)-2]
			stack = stack[:len(stack) - 2]
			stack = append(stack,right+left)
			case "-":
			right := stack[len(stack)-1]
			left := stack[len(stack)-2]
			stack = stack[:len(stack) - 2]
			stack = append(stack,left-right)

			case "*":
			right := stack[len(stack)-1]
			left := stack[len(stack)-2]
			stack = stack[:len(stack) - 2]
			stack = append(stack, left*right)
			case "/":
			right := stack[len(stack)-1]
			left := stack[len(stack)-2]
			stack = stack[:len(stack) - 2]
			stack = append(stack, left/right)
			default:
				val,_ := strconv.Atoi(token)
				stack = append(stack, val)
		}
	}
	return stack[0]
}
