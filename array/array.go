package array


type Comparable interface {
	Greater(c Comparable) bool
}


func IntsToInterfaces(ints []int) []interface{} {
	interaces := make([]interface{},0)
	for _,i := range ints{
		interaces = append(interaces,i)
	}
	return interaces
}

func IsExist(nums []interface{}, targetNum interface{}) bool {
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
		if targetNum.Greater(nums[mid]) {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}
