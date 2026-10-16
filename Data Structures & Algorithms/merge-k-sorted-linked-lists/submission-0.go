/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type PQ []*ListNode
func(h PQ) Len() int {return len(h)}
func(h PQ) Less(i, j int) bool {return h[i].Val < h[j].Val}
func(h PQ) Swap(i, j int) {h[i], h[j] = h[j], h[i]}
func(h *PQ) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
} 
func(h *PQ) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = (*h)[:n-1]
	return x
}


func mergeKLists(lists []*ListNode) *ListNode {
    pq := PQ{}
	for _, list := range lists {
		if list == nil {
			continue
		}
		heap.Push(&pq, list)
	}
	dummy := &ListNode{}
	p := dummy
	for pq.Len() > 0 {
		cur := heap.Pop(&pq).(*ListNode)
		next := cur.Next
		p.Next = cur
		p = p.Next
		if next != nil {
			heap.Push(&pq, next)
		}
	}
	p.Next = nil
	return dummy.Next
}
