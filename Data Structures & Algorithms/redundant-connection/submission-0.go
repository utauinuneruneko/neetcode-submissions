type UF struct {
	cnt int
	p []int
}

func Constructor(cnt int) UF {
	p := make([]int, cnt)
	for i := range p {
		p[i] = i
	}
	return UF{
		cnt: cnt,
		p: p,
	}
}

func(this UF) find(x int) int {
	if this.p[x] == x {
		return x
	}
	this.p[x] = this.find(this.p[x])
	return this.p[x]
}

func(this UF) union(p, q int) {
	pp, pq := this.find(p), this.find(q)
	if pp == pq {
		return
	}
	this.p[pp] = pq
}

func(this UF) isConnected(p, q int) bool {
	return this.find(p) == this.find(q)
}


func findRedundantConnection(edges [][]int) []int {
    n := len(edges)
	uf := Constructor(n+1)
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		if uf.isConnected(u, v) {
			return []int{u, v}
		}
		uf.union(u, v)
	}
	return []int{}
}
