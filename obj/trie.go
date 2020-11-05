package obj

import "fmt"

type Trie struct {
	root       *TreeNode
	countOfSon int
}

func NewTrie(countOfSon int) *Trie {
	return &Trie{
		root:       &TreeNode{},
		countOfSon: countOfSon,
	}
}

func (t *Trie) Insert(objs []Intable) {
	curRoot := t.root
	for _, obj := range objs {
		if curRoot.Next == nil {
			curRoot.Next = make([]*TreeNode, t.countOfSon)
		}
		if len(curRoot.Next) <= obj.Int() {
			panic(fmt.Sprintf("Count of curRoot's son too small,current is %d, expected %d", len(curRoot.Next), obj.Int()))
		}
		if curRoot.Next[obj.Int()] == nil {
			curRoot.Next[obj.Int()] = &TreeNode{
				Val:  false,
				Next: make([]*TreeNode, t.countOfSon),
			}
		}
		curRoot = curRoot.Next[obj.Int()]
	}
	curRoot.Val = true
}

func (t *Trie) IsExist(objs []Intable) bool {
	curRoot := t.root
	for _, obj := range objs {
		if len(curRoot.Next) <= obj.Int() {
			panic(fmt.Sprintf("Count of son too small,current is %d, expected %d", t.countOfSon, obj.Int()))
		}
		if curRoot.Next[obj.Int()] == nil {
			return false
		}
		curRoot = curRoot.Next[obj.Int()]
	}

	return curRoot.Val == true
}


