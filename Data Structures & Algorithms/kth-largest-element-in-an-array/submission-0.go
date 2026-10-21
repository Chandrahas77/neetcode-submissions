type MinHeap []int
func (h MinHeap) Len() int{
	return len(h)
}
func (h MinHeap) Less(i,j int) bool{
	return h[i] < h[j]
}
func (h MinHeap) Swap(i,j int){
	h[i],h[j] = h[j],h[i]
}
func (h *MinHeap) Push(x any){
	*h = append(*h,x.(int))
}
func (h *MinHeap) Pop() any{
	old := *h
	n := len(old)
	curr := old[n-1]
	old = old[:n-1]
	*h = old
	return curr
}
func findKthLargest(nums []int, k int) int {
	h := (*MinHeap)(&nums)
	heap.Init(h)
	
	for h.Len() > k{
		heap.Pop(h)
	}
	return (*h)[0]
}
