package _1_array

/*
题目描述 / Problem Description
给定一个按非递减顺序排列的整数数组 nums，返回每个元素平方后仍按非递减顺序排列的新数组。
Given an integer array nums sorted in nondecreasing order,
return a new array containing each value's square, also sorted in nondecreasing order.
*/

import (
	"sort"
)

// 1. Sorting: square every value in place and then sort the result.
// 1. 排序法：先原地把每个数平方，再对结果排序。
// Time: O(n log n), Space: O(log n) for the sort stack; the input slice is modified and reused.
// 时间复杂度：O(n log n)，空间复杂度：O(log n)，由排序调用栈产生；输入切片被原地修改并直接复用。
//
// 先原地平方再排序，会修改 nums；要求平方值可由 int 表示。
// Square and sort in place, modifying nums; squared values must fit in int.
func sortedSquares(nums []int) []int {
	for i, val := range nums {
		nums[i] *= val
	}
	sort.Ints(nums)
	return nums
}

// 2. Two pointers: the largest square must come from one of the two ends.
// 2. 双指针法：最大平方值一定来自当前区间的最左端或最右端。
// Time: O(n), Space: O(1) auxiliary beyond the O(n) result.
// 时间复杂度：O(n)，除 O(n) 结果数组外辅助空间 O(1)。
//
// 有序数组剩余部分的最大绝对值一定在两端；比较两端平方，把较大者写到 ans[k]，从后往前填。
// The largest absolute value in the remaining sorted interval is at an end; place its square at ans[k] from right to left.
// 只移动被取走的一端，i<=j 保证最后一项不遗漏；输入不变，平方值须在 int 范围内。
// Advance only the chosen endpoint; i<=j includes the final item. Input is unchanged and squares must fit in int.
func sortedSquares_TwoPointers(nums []int) []int {
	n := len(nums)
	i, j, k := 0, n-1, n-1
	ans := make([]int, n)

	for i <= j {
		lm, rm := nums[i]*nums[i], nums[j]*nums[j]
		if lm > rm {
			ans[k] = lm
			i++
		} else {
			ans[k] = rm
			j--
		}
		k--
	}
	return ans
}

// 3. Merge around the sign boundary: treat the two halves as two sorted square sequences.
// 3. 分界归并法：把负数段和非负数段看成两个已排序的平方序列，再归并。
// Time: O(n), Space: O(1) auxiliary beyond the O(n) result; the input is not modified.
// 时间复杂度：O(n)，除 O(n) 结果数组外辅助空间 O(1)；不修改输入。
//
// right 从首个非负数向右，left 从最后一个负数向左；两路平方都递增，可像合并有序数组一样取较小者。
// Scan nonnegative values rightward and negative values leftward; both square sequences increase, so merge their minima.
// 一路耗尽就取另一路；短路判断先检查边界，避免读取越界下标。
// When one side is exhausted, take the other; short-circuit boundary checks prevent invalid reads.
func sortedSquaresMerge(nums []int) []int {
	right := 0
	for right < len(nums) && nums[right] < 0 {
		right++
	}

	left := right - 1
	result := make([]int, 0, len(nums))

	for left >= 0 || right < len(nums) {
		takeLeft := right == len(nums) ||
			(left >= 0 && nums[left]*nums[left] <= nums[right]*nums[right])

		if takeLeft {
			result = append(result, nums[left]*nums[left])
			left--
		} else {
			result = append(result, nums[right]*nums[right])
			right++
		}
	}

	return result
}
