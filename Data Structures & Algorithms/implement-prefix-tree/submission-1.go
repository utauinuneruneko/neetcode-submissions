type PrefixTree struct {
	children [26]*PrefixTree
    isEnd    bool
}

func Constructor() PrefixTree {
    return PrefixTree{}
}

func (this *PrefixTree) Insert(word string) {
	node := this
    for _, ch := range word {
        idx := ch - 'a'
        if node.children[idx] == nil {
            node.children[idx] = &PrefixTree{}
        }
        node = node.children[idx]
    }
    node.isEnd = true
}

func (this *PrefixTree) Search(word string) bool {
	node := this.searchPrefix(word)
    return node != nil && node.isEnd
}


func (this *PrefixTree) StartsWith(prefix string) bool {
    return this.searchPrefix(prefix) != nil
}

// 辅助函数：提取公共的搜索逻辑
func (this *PrefixTree) searchPrefix(prefix string) *PrefixTree {
    node := this
    for _, ch := range prefix {
        idx := ch - 'a'
        if node.children[idx] == nil {
            return nil
        }
        node = node.children[idx]
    }
    return node
}
