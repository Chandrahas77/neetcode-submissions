func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	res := make([]int,n)
	//stores the indices of temperatures
	stack := make([]int,0,n)

	for i := 0; i < n; i++{
		currTemp := temperatures[i]
		for len(stack) > 0 && currTemp > temperatures[stack[len(stack) - 1]]{
			prevIdx := stack[len(stack) - 1]
			//pop the element
			stack = stack[:len(stack) - 1]
			res[prevIdx] = i - prevIdx
		}
		stack = append(stack,i)
	}
	return res
}
