package _1_array

/*
题目描述 / Problem Description
给定一个按非递减顺序排列的整数数组 nums，返回每个元素平方后仍按非递减顺序排列的新数组。
Given an integer array nums sorted in nondecreasing order,
return a new array containing each value's square, also sorted in nondecreasing order.

解题思路 / Solution Approach
文件提供排序法和双指针法。双指针比较数组两端的平方值，将较大值从结果数组末尾向前写入，
可在线性时间内完成。
The file provides sorting and two-pointer solutions. The two-pointer method compares
squared values at both ends and writes the larger one from the end of the result array.

关键逻辑：为什么这样做 / Why This Works
有序区间中的数都夹在两端值之间，所以最大绝对值一定在某一端，最大平方也在某一端。
每次比较两端就选出了剩余元素中的最大平方，应放入最大的空位 ans[k]，而不是结果开头。
移动被选中的端点并令 k-- 后，同样的理由继续成立；相遇时最后一个元素也要处理，因此条件是 i<=j。
Every value lies between the sorted interval's endpoints, so the maximum absolute value,
and hence maximum square, occurs at an endpoint. Comparing both selects the largest remaining square,
which belongs at the largest empty result index k. Remove that endpoint and decrement k to repeat.
Use i<=j to process the final remaining element too.

时间与空间复杂度 / Time and Space Complexity
n = len(nums)。sortedSquares：平方 O(n)，排序 O(n log n)，总时间 O(n log n)；
Go 当前整数排序的调用栈占 O(log n) 辅助空间，返回切片复用输入。
sortedSquares_TwoPointers：时间 O(n)，除结果外辅助空间 O(1)，新建结果占 O(n)。
n = len(nums). sortedSquares takes O(n log n) time, including squaring and sorting;
the current Go integer sort uses O(log n) stack space and the returned slice aliases the input.
sortedSquares_TwoPointers takes O(n) time, O(1) auxiliary space excluding its O(n) output.
*/

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
