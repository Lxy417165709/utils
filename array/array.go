package array

func IsExist(nums []int, targetNum int) bool {
	for _, num := range nums {
		if num == targetNum {
			return true
		}
	}
	return false
}
