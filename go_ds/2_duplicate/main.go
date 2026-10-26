package main

import (
	"fmt"
	"sort"
)

func main() {
	nums_1 := []int{1, 2, 3, 3} //true
	nums_2 := []int{1, 2, 3, 4} //false
	fmt.Println("brute", findDuplicate_brute(nums_1))
	fmt.Println("brute", findDuplicate_brute(nums_2))
	// fmt.Println(nums_2)
	arr := []int{7, -3, 7, 0, -3, 2, 9, -1}
	fmt.Println("better", findDuplicate_better(arr))
	fmt.Println("better", findDuplicate_better(nums_2))
	//HAshSet
	fmt.Println("hash set", hasDuplicate(arr))
	// hashSet Length
	fmt.Println("hash len", HashSetlen(arr))
}

// bruteforce

func findDuplicate_brute(numsAr []int) bool {
	for i := 0; i < len(numsAr); i++ {
		for j := i + 1; j < len(numsAr); j++ {
			if numsAr[i] == numsAr[j] {
				return true
			}
		}
	}
	return false
}

// Better
func findDuplicate_better(numsAr []int) bool {
	sort.Ints(numsAr)
	// fmt.Println(numsAr)
	for i := 1; i < len(numsAr); i++ {
		if numsAr[i-1] == numsAr[i] {
			return true
		}
	}
	return false
}

// HashSet
func hasDuplicate(numArr []int) bool {
	seen := make(map[int]bool)
	for _, num := range numArr {
		if seen[num] {
			return true
		}
		seen[num] = true
	}
	return false
}

// HashSet Length
// 7, -3, 7, 0, -3, 2, 9, -1
func HashSetlen(nums []int) bool {
	seen := make(map[int]struct{})
	for _, num := range nums {
		seen[num] = struct{}{}
	}
	fmt.Print("seen set", seen)
	return len(seen) < len(nums)
}
