package main

import "fmt"

func findMin(nums []int) int {
	min := nums[0]
	n := len(nums)
	for i := n - 1; i >= 0; i-- {
		if nums[i] > min {
			min = nums[i+1]
			break
		}
	}
	return min
}

func main() {
	// arr := []int{3, 4, 5, 1, 2}
	// arr := []int{4, 5, 6, 7, 0, 1, 2}
	arr := []int{11, 13, 15, 17}
	res := findMin(arr)
	fmt.Println(res)

}
