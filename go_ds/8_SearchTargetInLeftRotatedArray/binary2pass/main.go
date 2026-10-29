package main

import "fmt"

func search(nums []int, target int) int {
	// converge and find piviot
	l := 0
	r := len(nums) - 1

	for l < r {
		m := l + (r-l)/2

		if nums[m] > nums[r] {
			l = m + 1
		} else {
			r = m
		}
	}
	piviot := r
	// reset l & r
	l = 0
	r = len(nums) - 1
	// find the block where elements can be

	if target >= nums[piviot] && target <= nums[r] {
		l = piviot

	} else {

		r = piviot - 1
	}
	// find the element
	// fmt.Println("L", l, "R", r)
	for l <= r {
		mid := l + (r-l)/2

		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
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
