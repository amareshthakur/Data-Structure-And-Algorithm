package main

func twoSum(numbers []int, target int) []int {
	ans := []int{0, 0}
	for i := 0; i < len(numbers)-1; i++ {
		num_required := target - numbers[i]
		req_index, presnce := binarySerach(numbers[1:len(numbers)], num_required)
		if presnce {
			ans = []int{i, req_index}
		}
	}
	return ans
}
func binarySerach(arr []int, target int) (int, bool) {
	l, r := 0, len(arr)-1
	for l <= r {
		mid := l + (r-l)/2
		if arr[mid] == target {
			return mid, true
		} else if arr[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}

	}
	return -1, false
}
