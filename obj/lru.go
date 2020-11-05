package obj

import "fmt"

type LRU struct {
	doubleList *DoubleList
	keyNode    map[interface{}]*DoubleListNode
	cap        int
}

func NewLRU(cap int) *LRU {
	return &LRU{
		doubleList: NewDoubleList(),
		keyNode:    make(map[interface{}]*DoubleListNode),
		cap:        cap,
	}
}

func (l *LRU) GetMap() map[interface{}]*DoubleListNode {
	return l.keyNode
}

func (l *LRU) Insert(key interface{}, value interface{}) {
	fmt.Printf("Before insert, LRU's length is %d,LRU's cap is %d\n",l.doubleList.Size(),l.cap)
	if l.keyNode[key] == nil {
		l.keyNode[key] = &DoubleListNode{}
		if l.doubleList.Size() == l.cap {
			removedNode := l.doubleList.RemoveTail()
			delete(l.keyNode,removedNode.Val.(*Pair).Key)
			fmt.Printf("As LRU full, key <%s> has been remove\n", removedNode.Val.(*Pair).Key)
		}
	}
	l.keyNode[key].Val = &Pair{
		Key:   key,
		Value: value,
	}
	l.doubleList.Remove(l.keyNode[key])
	l.doubleList.InsertHead(l.keyNode[key])
}
