type MaxHeap []int
func (h MaxHeap) Len() int{
	return len(h)
}
func (h MaxHeap) Less(i,j int) bool{
	return h[i] > h[j]
}
func (h MaxHeap) Swap(i,j int){
	h[i],h[j] = h[j],h[i]
}
func (h *MaxHeap) Push(x any){
	*h = append(*h,x.(int))
}
func (h *MaxHeap) Pop() interface{}{
	old := *h
	n := len(old)
	x := old[n-1]
	// old[n-1] = nil
	*h = old[:n-1]
	return x
}
func lastStoneWeight(stones []int) int {
	pq := &MaxHeap{}
	heap.Init(pq)
	for _,sUp := range stones{
		heap.Push(pq,sUp)
	}
	for pq.Len() > 1{
		first := heap.Pop(pq).(int)
		second := heap.Pop(pq).(int)
		if first != second{
			heap.Push(pq,first - second)
		}
	}
	if pq.Len() == 0{
		return 0
	}
	return heap.Pop(pq).(int)
}
