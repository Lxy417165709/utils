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
	fmt.Println("IsSorted",isSorted(nums))
	QuickSort(nums)
	fmt.Println("After sorting", nums)
	fmt.Println("IsSorted",isSorted(nums))
}

func TestGetKthSmall(t *testing.T) {
	nums := []Comparable{
		Int(1), Int(8), Int(2),
		Int(4), Int(8), Int(7),
		Int(6), Int(7), Int(89),
		Int(1), Int(98), Int(3),
		Int(18), Int(9), Int(3),
		Int(18), Int(-5), Int(123),
	}
	fmt.Printf("%dth small: %v\n", 1, GetKthSmall(nums, 1))
	fmt.Printf("%dth small: %v\n", 5, GetKthSmall(nums, 5))
	fmt.Printf("%dth small: %v\n", 6, GetKthSmall(nums, 6))
	fmt.Printf("%dth small: %v\n", 10, GetKthSmall(nums, 10))
	fmt.Printf("%dth small: %v\n", 15, GetKthSmall(nums, 15))
	fmt.Printf("%dth small: %v\n", 16, GetKthSmall(nums, 16))
	fmt.Printf("%dth small: %v\n", 20, GetKthSmall(nums, 20))
}

func TestPermutation(t *testing.T) {
	nums := []interface{}{
		1, 1, 1, 2,
	}
	var resultSet [][]interface{}
	Permutation(&resultSet, nums, nil)
	fmt.Println(len(resultSet))
	fmt.Println(resultSet)
}

func TestSpiralOrder(t *testing.T) {
	matrix := [][]interface{}{
		{1, 2, 3, 100},
		{8, 9, 4, 1000},
		{7, 6, 5, 10000},
	}
	fmt.Println(SpiralOrder(matrix))
}

type ListNode struct {
	Val  int
	Next *ListNode
}
