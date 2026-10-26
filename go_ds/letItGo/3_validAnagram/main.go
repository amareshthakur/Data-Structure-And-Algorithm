package main

import (
	"sort"
)

func main() {
	s := "racecar"
	t := "carrace"
	validAnagram(s, t)
}

// HashTable Using Array
// Time Complexity -> O(n+m)
// Space Complexity -> O(1)
func validAnagram(s1, s2 string) bool {
	// fmt.Println(s1, s2)
	l1 := len(s1)
	l2 := len(s2)

	if l1 != l2 {
		return false
	}
	var count [26]int

	for i := 0; i < l1; i++ {
		count[s1[i]-'a']++
		count[s2[i]-'a']--
	}
	for _, num := range count {
		if num != 0 {
			return false
		}
	}

	return true
}

// Sorting
// Time Complexity : O(nlogn + mlogm)
// Space Coplexity : O(1) or O(n+m) based on sorting algorithm
// n -> len of string s
// m -> len of string t
func validAnagram_sorting(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	sRunes, tRunes := []rune(s), []rune(t)

	sort.Slice(sRunes, func(i, j int) bool { return sRunes[i] < sRunes[j] })
	sort.Slice(tRunes, func(i, j int) bool { return tRunes[i] < tRunes[j] })

	for i := range sRunes {
		if sRunes[i] != tRunes[i] {
			return false
		}
	}

	return true
}

// HashMAp
// Time Complexity -> O(n+m)
// Space Complexity -> O(1)
func isAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false

	}
	countS, countT := make(map[rune]int), make(map[rune]int)

	for i, ch := range s {
		countS[ch]++
		countT[rune(t[i])]++
	}
	for k, v := range countS {
		if countT[k] != v {
			return false
		}
	}
	return true
}
