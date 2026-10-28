package main

import "fmt"

func main() {
	arr := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}

	maxSum := maxSubArray_kadanes(arr)
	fmt.Println(maxSum)
}

func maxSubArray_kadanes(numms []int) int {
	maxSubArray := numms[0]
	curSum := 0
	n := len(numms)
	for i := 1; i < n; i++ {
		if curSum < 0 {
			curSum = numms[i]
		} else {
			curSum += numms[i]
		}
		maxSubArray = max(maxSubArray, curSum)
	}
	return maxSubArray
}
