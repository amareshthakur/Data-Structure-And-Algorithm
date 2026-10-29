package main

import "fmt"

func maxProduct(nums []int) int {
	res := nums[0]
	curMin, curMax := 1, 1
	for i := 0; i < len(nums); i++ {
		temp := curMax * nums[i]
		curMax = max(temp, max(nums[i]*curMin, nums[i]))
		curMin = min(temp, min(nums[i]*curMin, nums[i]))
		res = max(res, curMax)
	}
	return res
}

func main() {
	Array := []int{2, 3, -2, 4}
	// Array := []int{-2, 0, -1, 2, 4}
	fmt.Println(maxProduct(Array))
}
