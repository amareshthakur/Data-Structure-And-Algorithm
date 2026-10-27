package main

import "fmt"

// TC -> O(n)
// SC -> O(1) extra and O(n) for array
func productExceptSelf(nums []int) []int {
	n := len(nums)
	res := make([]int, len(nums))
	prefix := 1
	fmt.Println("PREFIX")
	for i := 0; i < n; i++ {
		res[i] = prefix
		fmt.Println("res[", i, "] -> ", res[i])
		prefix *= nums[i]
		fmt.Println("prefix -> ", prefix)
	}
	fmt.Println("POSTFIX")
	postfix := 1
	for i := n - 1; i >= 0; i-- {
		res[i] *= postfix
		fmt.Println("res[", i, "] -> ", res[i])
		postfix *= nums[i]
		fmt.Println("postfix -> ", postfix)
	}

	return res
}

// Prefix Suffix Sum
// TC -> O(n)
// SC -> O(n)
func productExceptSelf_pref_suf(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	pref := make([]int, n)
	suff := make([]int, n)

	pref[0], suff[n-1] = 1, 1

	for i := 1; i < n; i++ {
		pref[i] = nums[i+1] * suff[i+1]
	}
	for i := n - 2; i >= 0; i-- {
		suff[i] = nums[i+1] * suff[i+1]

	}
	for i := 0; i < n; i++ {
		ans[i] = pref[i] * suff[i]
	}
	return ans
}

// TC -> O(n)
// SC -> O(1) extra space && O(n) space for oparray
func productExceptSelf_division(nums []int) []int {
	n := len(nums)
	prod := 1
	zeroCount := 0
	for _, num := range nums {
		if num != 0 {
			prod *= num
		} else {
			zeroCount++
		}
	}
	ans := make([]int, n)
	if zeroCount > 1 {
		return ans
	}
	for i, num := range nums {
		if zeroCount > 0 {
			if num == 0 {
				ans[i] = prod
			} else {
				ans[i] = 0
			}
		} else {
			ans[i] = prod / num
		}
	}
	return ans
}

// TC -> O(n^2)
// Sc -> O(1) extra space
// 	  -> O(n) space for o/p array
func productExceptSelf_brute(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)

	for i := 0; i < n; i++ {
		mul := 1
		for j := 0; j < n; j++ {
			if i != j {
				mul *= nums[j]
			}
		}
		ans[i] = mul
	}

	return ans
}
func main() {
	arr := []int{1, 2, 3, 4}
	// arr := []int{-1, 1, 0, -3, 3}
	// op := productExceptSelf_division(arr)
	op := productExceptSelf(arr)
	fmt.Println(op)
}
