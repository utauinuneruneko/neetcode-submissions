type State struct {
	pos int
	cost int
	step int
}

type PQ []State
func(h PQ) Len() int {return len(h)}
func(h PQ) Less(i, j int) bool {return h[i].cost < h[j].cost}
func(h PQ) Swap(i, j int) {h[i], h[j] = h[j], h[i]}
func(h *PQ) Push(x interface{}) {
	*h = append(*h, x.(State))
}
func(h *PQ) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = (*h)[:n-1]
	return x
}

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	edges := make([][]State, n)
	for i := range edges {
		edges[i] = make([]State, 0)
	}
	for _, flight := range flights {
		u, v, cost := flight[0], flight[1], flight[2]
		edges[u] = append(edges[u], State{
			pos: v,
			cost: cost,
		})
	}

    visited := make([]map[int]int, n)
	for i := range visited {
		visited[i] = make(map[int]int)
	}

	pq := &PQ{}
	heap.Init(pq)
	heap.Push(pq, State{
		pos: src,
		cost: 0,
		step: 0,
	})

	for pq.Len() > 0 {
		cur := heap.Pop(pq).(State)
		pos, cost, step := cur.pos, cur.cost, cur.step
		if step > k+1 {
			continue
		}
		if _, ok := visited[pos][step]; ok && cost > visited[pos][step] {
			continue
		} 
		visited[pos][step] = cost
		for _, neighbor := range edges[pos] {
			npos, ncost, nstep := neighbor.pos, cost+neighbor.cost, step+1
			if nstep > k+1 {
				continue
			}
			if _, ok := visited[npos][nstep]; ok && ncost >= visited[npos][nstep] {
				continue
			} else {
				visited[npos][nstep] = ncost
				heap.Push(pq, State{
					pos: npos,
					cost: ncost,
					step: nstep,
				})
			}
		}
	}

	res := math.MaxInt32
	for _, v := range visited[dst] {
		res = min(res, v)
	}
	if res == math.MaxInt32{
		return -1
	}
	return res
}
