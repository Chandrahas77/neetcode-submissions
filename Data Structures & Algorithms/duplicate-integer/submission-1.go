func hasDuplicate(nums []int) bool {
    set := make(map[int]struct{})
	for _,v := range nums{
		if _,exists := set[v]; !exists{
			set[v] = struct{}{}
		}else{
			return true
		}
	}
	return false
}
