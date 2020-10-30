package array

import (
	"fmt"
	"testing"
)

type Int int

func (i Int) Greater(c Comparable) bool {
	j := c.(Int)
	return i > j
}

type String string

func (i String) Greater(c Comparable) bool {
	j := c.(String)
	return i > j
}

func Test(t *testing.T) {
	fmt.Println(IsExistInSortedArray([]Comparable{
		Int(1), Int(2), Int(3),
	}, Int(2)))
	fmt.Println(GetFirstGreaterIndex([]Comparable{
		String("123"), String("456"), String("789"),
	}, String("456")))
	fmt.Println(GetFirstGreaterOrEqualIndex([]Comparable{
		String("123"), String("456"), String("789"),
	}, String("456")))
}

func TestQuickSort(t *testing.T) {
	nums := []Comparable{
		Int(1), Int(8), Int(2),
		Int(4), Int(8), Int(7),
		Int(6), Int(7), Int(89),
		Int(1), Int(98), Int(3),
		Int(18), Int(9), Int(3),
		Int(18), Int(-5), Int(123),
	}
	fmt.Println("Before sorting", nums)
	QuickSort(nums)
	fmt.Println("After sorting", nums)
}
