type PQ []int
func(h PQ) Len() int {return len(h)}
func(h PQ) Less(i, j int) bool {return h[i] < h[j]}
func(h PQ) Swap(i, j int) {h[i], h[j] = h[j], h[i]}
func(h *PQ) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func(h *PQ) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = (*h)[:n-1]
	return x
}

type KthLargest struct {
    k int
	pq *PQ
}


func Constructor(k int, nums []int) KthLargest {
    pq := &PQ{}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	for i := 0; i < k && i < len(nums); i++ {
		heap.Push(pq, nums[i])
	}
	return KthLargest{
		k: k,
		pq: pq,
	}
}


func (this *KthLargest) Add(val int) int {
	heap.Push(this.pq, val)
	if this.k < this.pq.Len() {
		heap.Pop(this.pq)
	}
	return (*this.pq)[0]
}
