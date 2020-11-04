package obj

import (
	"fmt"
	"testing"
)

func TestDoubleList(t *testing.T) {
	doubleList := NewDoubleList()

	nodes := make([]*DoubleListNode,0)
	for i:=1;i<=5;i++{
		nodes = append(nodes,&DoubleListNode{
			Val: i,
		})
		doubleList.InsertTail(nodes[i-1])
	}

	fmt.Println(doubleList.RangeAll())
	doubleList.Remove(nodes[2])
	fmt.Println(doubleList.RangeAll())

}

func TestLRU(t *testing.T) {
	lru := NewLRU(3)
	keyValue := map[string]int{
		"haha":1,
		"lala":2,
		"java":3,
		"www":4,
	}
	for key,value := range keyValue{
		lru.Insert(key,value)
		fmt.Printf("%+v\n",lru.GetMap())
	}
}

