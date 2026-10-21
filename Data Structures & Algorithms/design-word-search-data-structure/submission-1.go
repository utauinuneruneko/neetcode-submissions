type WordDictionary struct {
    children []*WordDictionary
	isEnd bool
}

func Constructor() WordDictionary {
    return WordDictionary{
		children: make([]*WordDictionary, 26),
	}
}

func (this *WordDictionary) AddWord(word string)  {
	node := this
    for _, ch := range word {
        idx := ch - 'a'
        if node.children[idx] == nil {
            node.children[idx] = &WordDictionary{
				children: make([]*WordDictionary, 26),
			}
        }
        node = node.children[idx]
    }
    node.isEnd = true
}

func (this *WordDictionary) find(tree *WordDictionary, word string, index int) bool {
	if index == len(word) {
		return true
	}
	if tree == nil {
		return false
	}

	if word[index] == '.' {
		if len(word)-1 == index {
			for _, nxt := range tree.children {
				if nxt != nil && nxt.isEnd {
					return true
				}
			}
			return false
		} else {
			res := false
			for _, nxt := range tree.children {
				if nxt == nil {
					continue
				}
				res = res || this.find(nxt, word, index+1)
				if res {
					return true
				}
			}
			return false
		}
	} else {
		idx := word[index]-'a'
		if tree.children[idx] == nil {
			return false
		} else if !tree.children[idx].isEnd && index == len(word)-1 {
			return false
		} else if tree.children[idx].isEnd && index == len(word)-1 {
			return true
		}
		return this.find(tree.children[idx], word, index+1)
	}
}

func (this *WordDictionary) Search(word string) bool {
    if this.children == nil {
		return false
	}
	return this.find(this, word, 0)
}
