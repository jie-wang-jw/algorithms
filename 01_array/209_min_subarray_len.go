package _1_array

/*
题目描述 / Problem Description
给定一个正整数 target 和一个由正整数组成的数组 nums，找出总和大于或等于 target 的最短连续子数组，并返回其长度；如果不存在则返回 0。
Given a positive integer target and an array nums of positive integers, return the minimum length of a contiguous subarray whose sum is at least target; return 0 if none exists.

解题思路 / Solution Approach
使用滑动窗口。右指针不断扩大窗口并累加元素；当窗口和达到 target 时，持续移动左指针缩小窗口，同时更新最短长度。
Use a sliding window. Expand the right boundary and add values; whenever the sum reaches target, repeatedly shrink the left boundary while updating the minimum length.

时间与空间复杂度 / Time and Space Complexity
n = len(nums)。时间 O(n)：左右边界各最多前进 n 次，嵌套循环不是 O(n²)。辅助空间 O(1)，只维护窗口和与下标。
n = len(nums). Time O(n): each boundary advances at most n times, so the nested loops are not quadratic. Auxiliary space O(1) for the sum and indices.
*/

func minSubArrayLen(target int, nums []int) int {
	// i is the left boundary of the current sliding window.
	// i 是当前滑动窗口的左边界。
	i := 0
	l := len(nums)  // Length of the input array. / 输入数组的长度。
	sum := 0        // Sum of the current window nums[i:j+1]. / 当前窗口 nums[i:j+1] 的元素和。
	result := l + 1 // Use an impossible length as the initial answer. / 用不可能出现的长度作为初始答案。

	// j expands the right boundary one element at a time.
	// j 每次向右移动一位，扩张窗口右边界。
	for j := range l {
		sum += nums[j]
		// Once the sum reaches target, shrink from the left as much as possible.
		// 当窗口和达到 target 后，尽可能从左侧收缩窗口。
		for sum >= target {
			// Both i and j are included, so the length is j - i + 1.
			// i 和 j 都包含在窗口内，因此长度是 j - i + 1。
			//result = min(result, j-i+1)
			subLength := j - i + 1
			if subLength < result {
				result = subLength
			}

			// Remove nums[i] before moving the left boundary forward.
			// 左边界右移前，先从窗口和中减去 nums[i]。
			sum -= nums[i]
			i++
		}
	}
	if result == l+1 {
		// The sentinel was never updated, so no valid subarray exists.
		// 初始哨兵值从未更新，说明不存在满足条件的子数组。
		return 0
	}
	return result
}
