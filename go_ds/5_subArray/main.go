package main

import "fmt"

// func maxSubArray_brute(nums []int) (int[],int) {
func maxSubArray_brute(nums []int) int {
	// max := 0
	res := nums[0]
	n := len(nums)
	resArr := make([]int, n)
	for i := 0; i < n; i++ {
		sum := 0
		fmt.Println("Loop no ", i)
		for j := i; j < n; j++ {
			sum += nums[j]
			if sum > res {
				res = sum
				resArr = nums[i : j+1]
				fmt.Println("Arr", resArr, " sum -> ", sum)
			}
		}

	}
	return res
}

func main() {
	arr := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}

	maxSum := maxSubArray_brute(arr)
	fmt.Println(maxSum)
}
