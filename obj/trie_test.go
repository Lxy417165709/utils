package obj

import (
	"fmt"
	"testing"
)

type Byte byte

func (b Byte) Int() int {
	return int(b)
}

func TestTrie(t *testing.T) {
	trie := NewTrie(128)
	insertCases := [][]Intable{
		{Byte('a'), Byte('b'), Byte('c'), Byte('d'), Byte('e')},
		{Byte('a'), Byte('b'), Byte('c'), Byte('e'), Byte('e')},
		{Byte('h'), Byte('e'), Byte('l'), Byte('l'), Byte('o')},
		{Byte('h'), Byte('i')},
	}
	searchCases := [][]Intable{
		{Byte('a'), Byte('b'), Byte('c'), Byte('d'), Byte('e')},
		{Byte('a'), Byte('b'), Byte('c'), Byte('e'), Byte('e')},
		{Byte('h'), Byte('e'), Byte('l'), Byte('l'), Byte('o')},
		{Byte('h'), Byte('i')},
		{Byte('h')},
		{Byte('a')},
		{Byte('a'), Byte('b'), Byte('c'), Byte('e')},
		{Byte('a'), Byte('b'), Byte('e'), Byte('e')},
		{},
	}

	for _,testCase := range insertCases{
		trie.Insert(testCase)
	}
	for _,testCase := range searchCases{
		fmt.Println(testCase,trie.IsExist(testCase))
	}

}
