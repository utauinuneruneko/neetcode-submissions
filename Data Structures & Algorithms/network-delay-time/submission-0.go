type State struct {
	Pos int
	Cost int
}

type PQ []State
func (h PQ) Len() int {return len(h)}
func(h PQ) Less(i, j int) bool {return h[i].Cost < h[j].Cost}
func(h PQ) Swap(i, j int) {h[i], h[j] = h[j], h[i]}
func(h *PQ) Push(x interface{}) {
	*h = append(*h, x.(State))
}
func(h *PQ) Pop () interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = (*h)[:n-1]
	return x
}

func networkDelayTime(times [][]int, n int, k int) int {
	edges := make([][]State, n+1)
	for i := 1; i <= n; i++ {
		edges[i] = make([]State, 0)
	}
	for _, time := range times {
		u, v, cost := time[0], time[1], time[2]
		edges[u] = append(edges[u], State{Pos: v, Cost: cost})
	}
    pq := &PQ{}
	heap.Push(pq, State{k, 0})
	visited := make([]int, n+1)
	for i := range visited{
		visited[i] = math.MaxInt32
	}
	visited[0] = 0
	for pq.Len() > 0 {
		cur := heap.Pop(pq).(State)
		cost := cur.Cost
		pos := cur.Pos
		if cost < visited[pos] {
			visited[pos] = cost
		} else {
			continue
		}
		for _, next := range edges[pos] {
			if visited[next.Pos] < cost+next.Pos {
				continue
			} 
			heap.Push(pq, State{next.Pos, cost+next.Cost})
		}
	}
	res := 0
	for _, val := range visited {
		if val == math.MaxInt32 {
			return -1
		}
		res = max(res, val)
	}
	return res
}
