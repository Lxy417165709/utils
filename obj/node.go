package obj

type Intable interface {
	Int() int
}

type DoubleListNode struct {
	Val  interface{}
	Pre  *DoubleListNode
	Next *DoubleListNode
}

type Pair struct {
	Key   interface{}
	Value interface{}
}

type TreeNode struct {
	Val  interface{}
	Next []*TreeNode
}

