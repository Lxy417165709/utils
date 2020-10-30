package array

type Comparable interface {
	Greater(c Comparable) bool
}

func IntsToInterfaces(ints []int) []interface{} {
	interaces := make([]interface{}, 0)
	for _, i := range ints {
		interaces = append(interaces, i)
	}
	return interaces
}

// ---------------- 查找相关 ----------------
func IsExist(nums []Comparable, targetNum Comparable) bool {
	for _, num := range nums {
		if num == targetNum {
			return true
		}
	}
	return false
}

func IsExistInSortedArray(nums []Comparable, targetNum Comparable) bool {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == targetNum {
			return true
		}
		if nums[mid].Greater(targetNum) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return false
}

func GetFirstGreaterIndex(nums []Comparable, targetNum Comparable) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid].Greater(targetNum) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}

func GetFirstGreaterOrEqualIndex(nums []Comparable, targetNum Comparable) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if isGreaterOrEqual(nums[mid], targetNum) {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}

// ---------------- 排序相关 ----------------
func QuickSort(nums []Comparable) {
	// 要加这个限制，否则 right 可能等于 -1
	if len(nums) == 0 {
		return
	}

	refNumIndex := 0
	left, right := 0, len(nums)-1
	for left <= right {
		for left <= right && isSmallerOrEqual(nums[left], nums[refNumIndex]) {
			left++
		}
		for left <= right && isGreaterOrEqual(nums[right], nums[refNumIndex]) {
			right--
		}
		if left <= right {
			nums[right], nums[left] = nums[left], nums[right]
		}
	}
	nums[right], nums[refNumIndex] = nums[refNumIndex], nums[right]
	QuickSort(nums[:right])
	QuickSort(nums[right+1:])
}

func isEqual(firstNum Comparable, secondNum Comparable) bool {
	return !firstNum.Greater(secondNum) && !secondNum.Greater(firstNum)
}

func isGreaterOrEqual(firstNum Comparable, secondNum Comparable) bool {
	return firstNum.Greater(secondNum) || isEqual(firstNum, secondNum)
}

func isSmallerOrEqual(firstNum Comparable, secondNum Comparable) bool {
	return !firstNum.Greater(secondNum)
}
