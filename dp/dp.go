package dp

import "github.com/Lxy417165709/utils/array"

// ----------------------- 最长递增子串 -----------------------
func LongestIncrementSubArray(nums []array.Comparable) []array.Comparable {
	if len(nums) == 0 {
		return nil
	}

	lis := make([][]array.Comparable, len(nums))
	lis[0] = []array.Comparable{
		nums[0],
	}
	result := lis[0]

	for i := 1; i < len(nums); i++ {
		if nums[i].Greater(nums[i-1]) {
			lis[i] = append(lis[i-1], nums[i])
		} else {
			lis[i] = []array.Comparable{nums[i]}
		}
		if len(lis[i]) > len(result) {
			result = lis[i]
		}
	}
	return result
}

// ----------------------- 最长公共子串 -----------------------
func LongestSubArray(arr1 []interface{}, arr2 []interface{}) []interface{} {
	/*
		dp[i][t] -> 以 arr[:i] 为结尾, 及以 arr2[:t] 为结尾的最长公共子串

		则有如下关系式

		dp[i][t] = {
			if i == 0: 0
			if t == 0: 0
			if arr1[i-1] == arr2[t-1]: dp[i-1][t-1] + 1
			else: 0
		}
	*/

	lsa := make([][][]interface{}, len(arr1)+1)
	for i := 0; i < len(lsa); i++ {
		lsa[i] = make([][]interface{}, len(arr2)+1)
	}
	var result []interface{}
	for i := 1; i <= len(arr1); i++ {
		for t := 1; t <= len(arr2); t++ {
			if arr1[i-1] == arr2[t-1] {
				lsa[i][t] = append(lsa[i-1][t-1], arr1[i-1])
			}
			if len(lsa[i][t]) > len(result) {
				result = lsa[i][t]
			}
		}
	}
	return result
}

// ----------------------- 最长公共子序列 -----------------------
func LongestSubSequence(arr1 []interface{}, arr2 []interface{}) []interface{} {
	/*
		dp[i][t] -> 以 arr[:i] 为结尾, 及以 arr2[:t] 为结尾的最长公共子序列

		则有如下关系式

		dp[i][t] = {
			if i == 0: 0
			if t == 0: 0
			if arr1[i-1] == arr2[t-1]: dp[i-1][t-1] + 1
			else: max(dp[i-1][t], dp[i][t-1])
		}
	*/
	lss := make([][][]interface{}, len(arr1)+1)
	for i := 0; i < len(lss); i++ {
		lss[i] = make([][]interface{}, len(arr2)+1)
	}
	var result []interface{}
	for i := 1; i <= len(arr1); i++ {
		for t := 1; t <= len(arr2); t++ {
			if arr1[i-1] == arr2[t-1] {
				lss[i][t] = append(lss[i-1][t-1], arr1[i-1])
			} else {
				lss[i][t] = longestArray(
					lss[i-1][t],
					lss[i][t-1],
				)
			}
			if len(lss[i][t]) > len(result) {
				result = lss[i][t]
			}
		}
	}
	return result
}

func longestArray(arrs ...[]interface{}) []interface{} {
	var result []interface{}
	for _, arr := range arrs {
		if len(arr) > len(result) {
			result = arr
		}
	}
	return result
}
