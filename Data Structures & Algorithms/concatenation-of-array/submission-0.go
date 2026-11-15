func getConcatenation(nums []int) []int {
    n := len(nums)
	for i := range n{
		nums = append(nums,nums[i])
	}
	return nums
}
