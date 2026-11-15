func getConcatenation(nums []int) []int {
    n := len(nums)
	res := make([]int,2*len(nums))
	for i := range nums{
		res[i],res[i+n] = nums[i],nums[i]
	}
	return res
}
