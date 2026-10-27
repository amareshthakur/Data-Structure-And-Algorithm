package main

import "fmt"

func main() {
	nums_1 := []int{1, 2, 3, 1}
	k1 := 3
	res := containDuplicate_Hmap(nums_1, k1)
	fmt.Println(res)
}

// [0, 1, 2, 3]
// [1, 2, 3, 1]
// k=3
func containDuplicate_Hmap(nums []int, K int) bool {
	mp := make(map[int]int)
	for i, num := range nums {
		j, ok := mp[num]
		// fmt.Println(mp)
		if ok && i-j <= K {
			return true
		}
		mp[num] = i
	}
	return false
}
