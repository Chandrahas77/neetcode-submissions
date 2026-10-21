type Point struct{
	coords []int
	dist int
}
type MaxHeap []Point

func (h MaxHeap) Len() int{
	return len(h)
}
func (h MaxHeap) Less(i,j int) bool{
	return h[i].dist > h[j].dist
}
func (h MaxHeap) Swap(i,j int){
	h[i],h[j] = h[j],h[i]
}
func (h *MaxHeap) Push(x any){
	*h = append(*h,x.(Point))
}
func (h *MaxHeap) Pop() any{
	old := *h
	n := len(old)
	curr := old[n-1]
	*h = old[:n-1]
	return curr
}
func kClosest(points [][]int, k int) [][]int {
	h := &MaxHeap{}
	heap.Init(h)
	for _,pt := range points{
		dist := pt[0]*pt[0] + pt[1]*pt[1]
		p := Point{dist: dist,coords: pt}
		if h.Len() < k{
			heap.Push(h,p)
		} else if dist < (*h)[0].dist{
			heap.Pop(h)
			heap.Push(h,p)
		}
	}
	res := make([][]int,k)
	for i := range k{
		res[i] = heap.Pop(h).(Point).coords
	}
	return res
}
