package _1_array

import (
	"sort"
)

// Square every value and then sort the result.
// 排序法：先把每个数平方，再对结果排序。
func sortedSquares(nums []int) []int {
	for i, val := range nums {
		nums[i] *= val
	}
	sort.Ints(nums)
	return nums
}

// Two pointers: the largest square must come from one of the two ends.
// 双指针法：最大平方值一定来自当前区间的最左端或最右端。
func sortedSquares_TwoPointers(nums []int) []int {
	n := len(nums)
	// i scans from the left, j scans from the right, and k writes from the end.
	// i 从左扫描，j 从右扫描，k 从结果数组末尾向前写入。
	i, j, k := 0, n-1, n-1
	ans := make([]int, n)

	for i <= j {
		// Compare squares because a large negative value may have a larger square.
		// 比较平方值，因为绝对值较大的负数平方后也可能最大。
		lm, rm := nums[i]*nums[i], nums[j]*nums[j]
		if lm > rm {
			// Put the larger square at the current largest unfilled position.
			// 把较大的平方值放到当前尚未填充的最大位置 k。
			ans[k] = lm
			i++
		} else {
			ans[k] = rm
			j--
		}
		// One result position has been filled, so move k to the left.
		// 当前结果位置已经填好，k 向左移动一位。
		k--
	}
	return ans
}
