package bit

import (
	"fmt"
	"testing"
)

func TestGetCountOfOne(t *testing.T) {
	fmt.Println(GetCountOfOne(0))    // 0
	fmt.Println(GetCountOfOne(1))    // 1
	fmt.Println(GetCountOfOne(1127)) // 6
	fmt.Println(GetCountOfOne(-1))   // 64
}

func TestGetCountOfDiffBit(t *testing.T) {
	fmt.Println(GetCountOfDiffBit(1,2))			// 2
	fmt.Println(GetCountOfDiffBit(100,99))		// 3
	fmt.Println(GetCountOfDiffBit(100,-99))		// 62
	fmt.Println(GetCountOfDiffBit(-1,-2))		// 1
	fmt.Println(GetCountOfDiffBit(1,-1))		// 63
	fmt.Println(GetCountOfDiffBit(0,-1))		// 64
	fmt.Println(GetCountOfDiffBit(0,1))		    // 1
}