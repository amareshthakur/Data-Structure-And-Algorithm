//PS P:\DSA\go_ds\5_subArray>\recursion\main.go
package main

import "fmt"

func maxSubArray_recursion(nums []int) int {
	var dfs func(i int, flag bool) int
	n := len(nums)
	// {-2, 1, -3, 4, -1, 2, 1, -5, 4}
	dfs = func(i int, flag bool) int {
		if i == n-1 {
			if flag {
				return max(0, nums[i])
			}
			return nums[i]
		}
		if flag {
			return max(0, nums[i]+dfs(i+1, true))
		}
		return max(dfs(i+1, false), nums[i]+dfs(i+1, true))
	}

	return dfs(0, false)
}
func main() {
	arr := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}

	maxSum := maxSubArray_recursion(arr)
	fmt.Println(maxSum)
}
