func topKFrequent(nums []int, k int) []int {
	counts := make(map[int]int)
	for _,v := range nums{
		counts[v]++
	}
	//buckets is number indexed
	buckets := make([][]int,len(nums)+1)
	for num,c := range counts{
		buckets[c] = append(buckets[c],num)
	}

	result := make([]int,0,k)
	for i := len(buckets) - 1; i > 0 && len(result) < k; i --{
		result = append(result, buckets[i]...)
	}
	return result
}
