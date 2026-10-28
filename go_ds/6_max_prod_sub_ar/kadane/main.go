package main

import "fmt"

func maxProduct(nums []int) int {
	res := nums[0]
	curMin, curMax := 1, 1
	for i := 0; i < len(nums); i++ {
		temp := curMax * nums[i]
		// fmt.Println("i=", i)
		// fmt.Println("temp -> ", temp, "\tnums[", i, "]", nums[i], "\tcurrMax", curMax)
		// fmt.Println("nums[i]*curMax", nums[i]*curMax, "\tnums[i]*curMin", nums[i]*curMin)
		curMax = max(nums[i]*curMax, max(nums[i]*curMin, nums[i]))
		// fmt.Println("temp", temp, "\tnums[i]*curMin", nums[i]*curMin)
		curMin = min(temp, min(nums[i]*curMin, nums[i]))
		// fmt.Println("max(res, curMax)", res, "\t", curMax)
		res = max(res, curMax)
	}
	return res
}

func main() {
	Array := []int{2, 3, -2, 4}
	// Array := []int{-2, 0, -1, 2, 4}
	fmt.Println(maxProduct(Array))
}
