package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(nums)
	// fmt.Println(nums)
	ans := getConcatenation(nums)
	fmt.Println("ans", ans)
	ans = getConcatenation1(nums)
	fmt.Println("ans", ans)
	fmt.Println("nums", nums)

}

//	n -> length
// ans[i] == nums [i]
// ans[i+n] == nums [i]
//

func getConcatenation(nums []int) []int {
	ans := make([]int, 0, 2*len(nums))
	for i := 0; i < 2; i++ {
		for _, num := range nums {
			ans = append(ans, num)
		}
	}
	return ans
}

func getConcatenation1(nums []int) []int {
	n := len(nums)
	ans := make([]int, 2*n)
	for i, num := range nums {
		ans[i] = num
		ans[i+n] = num
	}
	return ans
}
