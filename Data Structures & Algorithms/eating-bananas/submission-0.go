func minEatingSpeed(piles []int, h int) int {
	maxPile := 0
	for _,p := range piles{
		if p > maxPile{
			maxPile = p
		}
	}
	ans := maxPile
	left := 1
	right := maxPile
	for left <= right{
		mid := left + (right - left)/2
		if canFinish(piles,h,mid){
			ans = mid
			right = mid - 1
		}else{
			left = mid + 1
		}
	}
	return ans
}

func canFinish(piles []int, h,k int) bool{
	hoursSpent := 0
	for _,p := range piles{
		hoursSpent += (p + k - 1)/k
		if hoursSpent > h{
			return false
		}
	}
	return hoursSpent <= h
}