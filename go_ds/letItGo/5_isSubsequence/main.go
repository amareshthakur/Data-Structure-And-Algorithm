package main

import "fmt"

func main() {
	// fmt.Println("hello")
	res := isSubsequence_recursion("heo", "hello")
	fmt.Println("Recursion", res)
}

// Best
//
func isSubsequence(s string, t string) bool {
	ls, lt := len(s), len(t)
	if ls == 0 {
		return true
	}

	var i, j int
	for i < ls && j < lt {
		if s[i] != t[j] {
			j++
			continue
		}

		if s[i] == t[j] {
			if i == ls-1 {
				return true
			}
			i++
			j++
		}
	}
	return false
}

// Recursion
// Time complexity -> O(n*m)
// /Space Complexity -> O(n)
func isSubsequence_recursion(s string, t string) bool {
	var rec func(i, j int) bool
	rec = func(i, j int) bool {
		if i == len(s) {
			return true
		}
		if j == len(t) {
			return false
		}

		if s[i] == t[j] {
			return rec(i+1, j+1)
		}
		return rec(i, j+1)
	}
	return rec(0, 0)
}

// Dynamic Programming Top Down
// Time complexity ->O(n*m)
// Space Complexity ->O(n*m)
func isSubsequence_dp_TD(s, t string) bool {
	n, m := len(s), len(t)
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, m)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	var rec func(i, j int) bool
	rec = func(i, j int) bool {
		if i == n {
			return true
		}
		if j == m {
			return false
		}
		if memo[i][j] != -1 {
			return memo[i][j] == 1
		}
		if s[i] == t[j] {
			if rec(i+1, j+1) {
				memo[i][j] = 1
			} else {
				memo[i][j] = 0
			}
		} else {
			if rec(i, j+1) {
				memo[i][j] = 1
			} else {
				memo[i][j] = 0
			}
		}
		return memo[i][j] == 1
	}
	return rec(0, 0)
}

// Dynamic Programming (Bottom-Up)
func isSubsequence_bu(s string, t string) bool {
	n, m := len(s), len(t)
	dp := make([][]bool, n+1)
	for i := range dp {
		dp[i] = make([]bool, m+1)
	}

	for j := 0; j <= m; j++ {
		dp[n][j] = true
	}

	for i := n - 1; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			if s[i] == t[j] {
				dp[i][j] = dp[i+1][j+1]
			} else {
				dp[i][j] = dp[i][j+1]
			}
		}
	}

	return dp[0][0]
}
