func merge(intervals [][]int) [][]int {
    sort.Slice(intervals,func(i,j int) bool{
		return intervals[i][0] < intervals[j][0]
	})
	res := [][]int{intervals[0]}
	for _,interval :=range intervals[1:]{
		start,end := interval[0],interval[1]
		lastEnd := res[len(res)-1][1]
		if start <=lastEnd{
			res[len(res)-1][1] = max(lastEnd,end)
		}else{
			res = append(res,[]int{start,end})
		}
	}
	return res
}
