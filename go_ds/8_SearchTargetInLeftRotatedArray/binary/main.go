package main

import "fmt"

func search(nums []int, target int) int {
	// converge leftP and rightP to minimum element's index i.e 1st element of array if array is not rotated
	left := 0
	right := len(nums) - 1

	for left < right {
		mid := left + (right-left)/2
		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	// Now left = right is index  to minimum element in array of left rotated sorted array
	// let's say it a piviot poit
	pivoit := left
	// fmt.Println(pivoit)

	var binarySearch func(left, right int) int
	binarySearch = func(left, right int) int {
		for left <= right {
			mid := left + (right-left)/2

			if nums[mid] == target {
				return mid
			} else if nums[mid] < target {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
		return -1
	}
	// search in left from piviot  of aray
	result := binarySearch(0, pivoit-1)
	if result != -1 {
		return result
	}
	// search in right  from piviot of aray
	return binarySearch(pivoit, len(nums)-1)

}

func main() {
	ar_1 := []int{4, 5, 6, 7, 0, 1, 2}
	// t1 := 4
	// res := search(ar_1, t1)
	// fmt.Println("ar1", res)

	t2 := 2
	res := search(ar_1, t2)
	fmt.Println("ar2", res)

	// ar3 := []int{1}
	// t3 := 0
	// res = search(ar3, t3)
	// fmt.Println("ar3", res)
}

// Input: nums = [4,5,6,7,0,1,2], target = 0
// Output: 4
// Input: nums = [4,5,6,7,0,1,2], target = 3
// Output: -1
// Input: nums = [1], target = 0
// Output: -1
