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

func GetKthSmall(nums []Comparable, k int) Comparable {
	if len(nums) == 0 {
		return nil
	}
	rightPositionIndex := Partition(nums)
	smallOrder := rightPositionIndex + 1
	if smallOrder == k {
		return nums[k-1]
	}
	if smallOrder > k {
		return GetKthSmall(nums[:rightPositionIndex], k)
	} else {
		return GetKthSmall(nums[rightPositionIndex+1:], k-smallOrder)
	}
}

// ---------------- 排序相关 ----------------
func QuickSort(nums []Comparable) {
	if len(nums) == 0 {
		return
	}
	rightPositionIndex := Partition(nums)
	QuickSort(nums[:rightPositionIndex])
	QuickSort(nums[rightPositionIndex+1:])
}

func Partition(nums []Comparable) int {
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
	return right
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

// ---------------- 排列组合 ----------------
func Permutation(resultSet *[][]interface{}, nums []interface{}, curNums []interface{}) {
	if len(nums) == 0 {
		// 如果 curNums 是空集，这里也加入结果集
		*resultSet = append(*resultSet, deepCopy(curNums))
		return
	}
	hasBeenHandled := make(map[interface{}]bool)
	for index, num := range nums {
		if hasBeenHandled[num] {
			continue
		}
		hasBeenHandled[num] = true
		nums[0], nums[index] = nums[index], nums[0]
		Permutation(resultSet, nums[1:], append(curNums, nums[0]))
		nums[0], nums[index] = nums[index], nums[0]
	}
}

// ---------------- 遍历 ----------------
// 螺旋矩阵
func SpiralOrder(matrix [][]interface{}) []interface{} {
	if len(matrix) == 0 {
		return nil
	}
	result := make([]interface{}, 0)
	left, right, up, down := 0, len(matrix[0])-1, 0, len(matrix)-1
	for {
		// 1. 向右走
		for x, y := up, left; y <= right; y++ {
			result = append(result, matrix[x][y])
		}
		up++
		if !(left <= right && up <= down) {
			break
		}

		// 2. 向下走
		for x, y := up, right; x <= down; x++ {
			result = append(result, matrix[x][y])
		}
		right--
		if !(left <= right && up <= down) {
			break
		}

		// 3. 向左走
		for x, y := down, right; y >= left; y-- {
			result = append(result, matrix[x][y])
		}
		down--
		if !(left <= right && up <= down) {
			break
		}

		// 4. 向上走
		for x, y := down, left; x >= up; x-- {
			result = append(result, matrix[x][y])
		}
		left++
		if !(left <= right && up <= down) {
			break
		}
	}
	return result
}

func deepCopy(oldSlice []interface{}) []interface{} {
	newSlice := make([]interface{}, len(oldSlice))
	copy(newSlice, oldSlice)
	return newSlice
}
