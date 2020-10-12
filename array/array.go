package array

func IsExist(nums []int, targetNum int) bool {
	for _, num := range nums {
		if num == targetNum {
			return true
		}
	}
	return false
}

func IsExistSortedNums(nums []int,targetNum int) bool{
	left,right := 0,len(nums)-1
	for left<=right{
		mid := left+(right-left)/2
		if nums[mid]==targetNum{
			return true
		}
		if nums[mid]>targetNum{
			right = mid - 1
		}else{
			left = mid + 1
		}
	}
	return false
}
