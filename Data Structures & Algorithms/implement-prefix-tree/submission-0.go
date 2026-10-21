type Tree struct {
	val byte
	next []*Tree
}

type PrefixTree struct {
	words map[string]struct{}
	trees []*Tree
}

func Constructor() PrefixTree {
	trees := make([]*Tree, 26)
    return PrefixTree{
		words: make(map[string]struct{}),
		trees: trees,
	}
}

func (this *PrefixTree) insertTree(tree *Tree, word string, index int) {
	if index == len(word) {
		return
	}
	idx := int(word[index]-'a')
	if tree.next[idx] == nil {
		tree.next[idx] = &Tree{
			val: word[index],
			next: make([]*Tree, 26),
		}
	}
	this.insertTree(tree.next[idx], word, index+1)
}

func (this *PrefixTree) Insert(word string) {
	if _, ok := this.words[word]; ok {
		return
	}
	this.words[word] = struct{}{}
	idx := int(word[0]-'a')
	if this.trees[idx] == nil {
		this.trees[idx] = &Tree{
			val: word[0],
			next: make([]*Tree, 26),
		}
	}
	this.insertTree(this.trees[idx], word, 1)
}

func (this *PrefixTree) Search(word string) bool {
	if _, ok := this.words[word]; ok {
		return true
	}
	return false
}

func (this *PrefixTree) find (tree *Tree, str string, index int) bool {
	if index == len(str) {
		return true
	}
	if tree.next[int(str[index]-'a')] == nil {
		return false
	}
	return this.find(tree.next[int(str[index]-'a')], str, index+1)
}

func (this *PrefixTree) StartsWith(prefix string) bool {
	if this.trees[int(prefix[0]-'a')] == nil {
		return false
	}
	return this.find(this.trees[int(prefix[0]-'a')], prefix, 1)
}
