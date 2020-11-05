package obj

type DoubleList struct {
	dummyHead *DoubleListNode
	dummyTail *DoubleListNode
	length    int
}

func NewDoubleList() *DoubleList {
	dummyHead := &DoubleListNode{}
	dummyTail := &DoubleListNode{}
	dummyHead.Next = dummyTail
	dummyTail.Pre = dummyHead
	return &DoubleList{
		dummyHead: dummyHead,
		dummyTail: dummyTail,
	}
}

func (d *DoubleList) Remove(nodeInList *DoubleListNode) *DoubleListNode {
	if nodeInList == nil || nodeInList.Pre == nil || nodeInList.Next == nil {
		return nil
	}
	nodeInList.Pre.Next = nodeInList.Next
	nodeInList.Next.Pre = nodeInList.Pre
	nodeInList.Next = nil
	nodeInList.Pre = nil
	d.length--
	return nodeInList
}

func (d *DoubleList) InsertTail(node *DoubleListNode) {
	node.Pre = d.dummyTail.Pre
	node.Next = d.dummyTail
	d.dummyTail.Pre.Next = node
	d.dummyTail.Pre = node
	d.length++
}

func (d *DoubleList) InsertHead(node *DoubleListNode) {
	node.Pre = d.dummyHead
	node.Next = d.dummyHead.Next
	d.dummyHead.Next.Pre = node
	d.dummyHead.Next = node
	d.length++
}

func (d *DoubleList) RemoveTail() *DoubleListNode {
	if d.length == 0 {
		return nil
	}
	return d.Remove(d.dummyTail.Pre)

}

func (d *DoubleList) RemoveHead() *DoubleListNode {
	if d.length == 0 {
		return nil
	}
	return d.Remove(d.dummyHead.Next)
}

func (d *DoubleList) RangeAll() []interface{} {
	result := make([]interface{}, 0)
	curHead := d.dummyHead
	for curHead.Next != d.dummyTail {
		result = append(result, curHead.Next.Val)
		curHead = curHead.Next
	}
	return result
}

func (d *DoubleList) Size() int {
	return d.length
}
