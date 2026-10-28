package main

import "fmt"

func maxProduct(nums []int) int {
	// {2, 3, -2, 4}
	res := nums[0]
	n := len(nums)
	for i := 0; i < n; i++ {
		cur := nums[i]
		res = max(res, cur)
		for j := i + 1; j < n; j++ {
			cur *= nums[j]
			res = max(res, cur)
		}
	}
	return res
}

func main() {
	Array := []int{2, 3, -2, 4}
	fmt.Println(maxProduct(Array))
}
