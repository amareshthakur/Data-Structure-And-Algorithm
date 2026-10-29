package main

import "fmt"

func search(nums []int, target int) int {
	l, r := 0, len(nums)-1

	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			return mid
		}
		if nums[l] <= nums[mid] {
			if target > nums[mid] || target < nums[l] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		} else {
			if target < nums[mid] || target > nums[r] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
	}
	return -1
}
func main() {
	ar_1 := []int{4, 5, 6, 7, 0, 1, 2}

	t2 := 2
	res := search(ar_1, t2)
	fmt.Println("ar2", res)

}
