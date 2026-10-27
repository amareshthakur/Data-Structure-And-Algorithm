package main

import "fmt"

func main() {
	// nums_1 := []int{1, 2, 3, 1}
	// k1 := 3
	// nums_2 := []int{1, 0, 1, 1}
	// k2 := 1
	nums_3 := []int{1, 2, 3, 1, 2, 3}
	k3 := 4

	// res1 := cotainDuplicate_brute(nums_1, k1)
	// fmt.Println("1 -> ", res1)
	// res2 := cotainDuplicate_brute(nums_2, k2)
	// fmt.Println("res2", res2)
	res3 := cotainDuplicate_brute(nums_3, k3)
	fmt.Println("res2", res3)
	// rs2 := containDuplicate_Hmap(nums_3, k3)
	// fmt.Println("rs2", rs2)
}

// Brute force
// TC -> O(n * min(n,k))
// SC -> O(1)

func cotainDuplicate_brute(nums []int, K int) bool {
	for L := 0; L < len(nums); L++ {
		for R := L + 1; R < min(len(nums), L+K+1); R++ {
			if nums[L] == nums[R] {
				return true
			}
		}
	}
	return false
}

// HAshMAp
// TC -> O(n)
// SC -> O(n)
func containDuplicate_Hmap(nums []int, K int) bool {
	mp := make(map[int]int)

	for i, num := range nums {
		if j, ok := mp[num]; ok && i-j <= K {
			return true
		}
		mp[num] = i
	}
	return false
}

// HashSet
// {1, 2, 3, 1, 2, 3}
func containDuplicate_Hset(nums []int, k int) bool {
	window := make(map[int]bool)
	L := 0
	for R := 0; R < len(nums); R++ {
		if R-L > k {
			delete(window, nums[L])
			L++
		}
		if window[nums[R]] {
			return true
		}
		window[nums[R]] = true
	}
	return false
}
