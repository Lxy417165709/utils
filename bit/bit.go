package bit

// num 可以是正数、负数，但不能溢出。
func GetCountOfOne(num int) int {
	count := 0
	for num != 0 {
		count++
		num = num & (num - 1)
	}
	return count
}

// num1、num2 可以是正数、负数，但不能溢出。
func GetCountOfDiffBit(num1, num2 int) int {
	return GetCountOfOne(num1 ^ num2)
}
