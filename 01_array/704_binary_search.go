package _1_array

/*
题目描述 / Problem Description
给定一个按升序排列的整数数组 nums 和目标值 target，返回 target 的下标；如果目标值不存在，则返回 -1。
Given an integer array nums sorted in ascending order and a target value, return the target's index or -1 if it is absent.
*/

// 1. Closed interval [left, right]: both boundaries may contain the target.
// 1. 左闭右闭区间 [left, right]：左右边界都可能是答案。
// Time: O(log n), Space: O(1).
// 时间复杂度：O(log n)，空间复杂度：O(1)。
//
// 候选区间是闭区间 [left,right]，left==right 仍有一个候选，所以循环条件是 <=。
// The candidate interval is inclusive [left,right]; equality still leaves one candidate, hence <=.
// 有序性保证可排除 mid 及其错误一侧：偏大取 right=mid-1，偏小取 left=mid+1；区间为空返回 -1。
// Sorted order excludes mid and the wrong side: use mid-1 or mid+1; return -1 when the interval is empty.
func search2(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return -1
}

// 2. Half-open interval [left, right): left is included, right is excluded.
// 2. 左闭右开区间 [left, right)：left 包含在内，right 不包含在内。
// Time: O(log n), Space: O(1).
// 时间复杂度：O(log n)，空间复杂度：O(1)。
//
// 候选区间是 [left,right)，right 不属于候选；left==right 时为空，所以循环条件是 <。
// The interval is half-open [left,right); right is excluded and equality means empty, hence <.
// mid 偏大时 right=mid 已排除 mid；偏小时必须 left=mid+1 才能排除 mid。前提是数组递增。
// Setting right=mid excludes an oversized mid; an undersized mid requires left=mid+1. The array must be sorted.
func search1(nums []int, target int) int {
	left := 0
	right := len(nums)

	for left < right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return -1
}

// 3. Recursive form of the closed interval: same invariant, expressed with recursion.
// 3. 左闭右闭区间的递归写法：不变量完全相同，只是改用递归表达。
// Time: O(log n), Space: O(log n) for the recursion stack.
// 时间复杂度：O(log n)，空间复杂度：O(log n)，由递归调用栈产生。
//
// 每层处理闭区间 [left,right]；left>right 才表示无候选。根据有序性排除 mid 后只递归可能的一半。
// Each call handles inclusive [left,right]; left>right is empty. Sorted order leaves only one half after excluding mid.
// 递归深度 O(log n)，因此辅助空间不是迭代版的 O(1)。
// Recursion depth is O(log n), unlike the iterative version's O(1) extra space.
func searchRecursive(nums []int, target int) int {
	var searchRange func(left, right int) int

	searchRange = func(left, right int) int {
		if left > right {
			return -1
		}

		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}

		if nums[mid] > target {
			return searchRange(left, mid-1)
		}

		return searchRange(mid+1, right)
	}

	return searchRange(0, len(nums)-1)
}
