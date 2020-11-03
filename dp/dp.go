package dp

// ----------------------- 最长公共子串 -----------------------
func LongestSubArray(arr1 []interface{}, arr2 []interface{}) []interface{} {
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
	lss := make([][][]interface{}, len(arr1)+1)
	for i := 0; i < len(lss); i++ {
		lss[i] = make([][]interface{}, len(arr2)+1)
	}
	var result []interface{}
	for i := 1; i <= len(arr1); i++ {
		for t := 1; t <= len(arr2); t++ {
			if arr1[i-1] == arr2[t-1] {
				//lss[i][t] = longestArray(
				//	append(lss[i-1][t-1], arr1[i-1]),
				//	lss[i-1][t],
				//	lss[i][t-1],
				//)	这种写法也可以
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
