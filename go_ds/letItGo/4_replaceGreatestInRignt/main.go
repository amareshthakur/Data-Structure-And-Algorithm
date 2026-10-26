package main

import "fmt"

func main() {
	// fmt.Println("Hello")
	r := []int{17, 18, 5, 4, 6, 1}
	// [18,6,6,6,1,-1]
	fmt.Println(r)
	fmt.Println(replaceElements_brute(r))

}

// Brute force
// Time Complexity -> O(n^2)
// Space complexity -> O(1)
func replaceElements_brute(arr []int) []int {
	n := len(arr)
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		rightMAx := -1
		// [17 18 5 4 6 1]
		for j := i + 1; j < n; j++ {
			if arr[j] > rightMAx {
				rightMAx = arr[j]
			}
		}
		ans[i] = rightMAx
	}
	return ans
}

// Suffix
// Time Complexity -> O(n)
// Space complexity -> O(1)
func replaceElements_suffix(arr []int) []int {
	n := len(arr)
	ans := make([]int, n)
	rightMAx := -1
	for i := n - 1; i >= 0; i-- {
		ans[i] = rightMAx
		if arr[i] > rightMAx {
			rightMAx = arr[i]
		}
	}
	return ans
}
