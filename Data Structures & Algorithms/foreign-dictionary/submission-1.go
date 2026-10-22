
func foreignDictionary(words []string) string {
	if len(words) == 1 {
		return words[0]
	}

	adj := make(map[byte]map[byte]bool)
	indegree := make(map[byte]int)

	// 1. 初始化：记录所有出现的字母，入度设为 0
	for _, word := range words {
		for i := 0; i < len(word); i++ {
			if _, ok := indegree[word[i]]; !ok {
				indegree[word[i]] = 0
			}
		}
	}

	// 2. 建图
	for i := 1; i < len(words); i++ {
		a, b := words[i-1], words[i]
		la, lb := len(a), len(b)
		k := 0
		for k < la && k < lb && a[k] == b[k] {
			k++
		}
		if k < la && k == lb {
			return ""
		} else if k == la {
			continue
		}

		u, v := a[k], b[k]
		if _, ok := adj[u]; !ok {
			adj[u] = make(map[byte]bool)
		}
		if adj[u][v] {
			continue
		}
		adj[u][v] = true
		indegree[v]++
	}

	// 3. 拓扑排序 (BFS)
	queue := []byte{}
	for char, degree := range indegree {
		if degree == 0 {
			queue = append(queue, char)
		}
	}

	// 为了保证结果的确定性（虽然题目说任意顺序），可以对初始 queue 排序
	// 但在 BFS 过程中，顺序由依赖关系决定
	
	res := make([]byte, 0)
	for len(queue) > 0 {
		// 弹出元素（为了满足某些测试用例的字母表顺序，这里可以选最小的，但通常不需要）
		u := queue[0]
		queue = queue[1:]
		res = append(res, u)

		for v := range adj[u] {
			indegree[v]--
			if indegree[v] == 0 {
				queue = append(queue, v)
			}
		}
	}

	// 4. 检查是否有环（结果长度是否等于总字母数）
	if len(res) < len(indegree) {
		return ""
	}
	return string(res)
}