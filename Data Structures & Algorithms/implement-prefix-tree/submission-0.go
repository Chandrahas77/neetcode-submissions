type Node struct{
	link [26]*Node
	flag bool
}
type PrefixTree struct {
	Root *Node
}

func Constructor() PrefixTree {
	return PrefixTree{
		Root : &Node{},
	}
}

func (this *PrefixTree) Insert(word string) {
	curr := this.Root
	for i := 0; i < len(word); i++{
		idx := word[i] - 'a'
		if curr.link[idx] == nil{
			curr.link[idx] = &Node{}
		}
		curr = curr.link[idx]
	}
	curr.flag = true
}

func (this *PrefixTree) Search(word string) bool {
	curr := this.Root
	for i := 0; i < len(word); i++{
		idx := word[i] - 'a'
		if curr.link[idx] == nil{
			return false
		}
		curr = curr.link[idx]
	}
	return curr.flag
}

func (this *PrefixTree) StartsWith(prefix string) bool {
	curr := this.Root
	for i :=0; i < len(prefix); i++{
		idx := prefix[i] - 'a'
		//path not exists; return false
		if curr.link[idx] == nil{
			return false
		}
		curr = curr.link[idx]
	}
	return true
}
