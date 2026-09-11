package _1_array

/*
题目描述 / Problem Description
给定一个按非递减顺序排列的整数数组 nums，返回每个元素平方后仍按非递减顺序排列的新数组。
Given an integer array nums sorted in nondecreasing order,
return a new array containing each value's square, also sorted in nondecreasing order.

解题思路 / Solution Approach
文件提供排序法、双指针法和分界归并法。双指针比较数组两端的平方值，将较大值从结果数组末尾向前写入，
可在线性时间内完成。
The file provides sorting, two-pointer, and sign-boundary merge solutions. The two-pointer method compares
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

补充解法：分界后归并 / Alternative: Merge around the Sign Boundary
sortedSquaresMerge 找到第一个非负数，负数部分向左走时平方递增，非负部分向右走时平方递增。
每次取两边较小平方写入结果，相当于归并两个有序序列；一边耗尽后使用另一边。
时间 O(n)，除 O(n) 输出外辅助空间 O(1)，不修改输入。
Locate the first nonnegative value. Negative squares increase leftward; nonnegative squares increase rightward.
Merge the smaller next square, consuming the remaining side when the other is exhausted.
Time O(n), O(1) auxiliary space beyond O(n) output; input is unchanged.
*/

import (
	"sort"
)

// 1. Sorting: square every value in place and then sort the result.
// 1. 排序法：先原地把每个数平方，再对结果排序。
// Time: O(n log n), Space: O(log n) for the sort stack; the input slice is modified and reused.
// 时间复杂度：O(n log n)，空间复杂度：O(log n)，由排序调用栈产生；输入切片被原地修改并直接复用。
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

// 3. Merge around the sign boundary: treat the two halves as two sorted square sequences.
// 3. 分界归并法：把负数段和非负数段看成两个已排序的平方序列，再归并。
// Time: O(n), Space: O(1) auxiliary beyond the O(n) result; the input is not modified.
// 时间复杂度：O(n)，除 O(n) 结果数组外辅助空间 O(1)；不修改输入。
func sortedSquaresMerge(nums []int) []int {
	// right stops at the first nonnegative value, which splits the two sequences.
	// right 停在第一个非负数上，这个位置把数组分成两段。
	right := 0
	for right < len(nums) && nums[right] < 0 {
		right++
	}

	// left walks backwards over the negatives, where squares grow as left decreases.
	// left 向左遍历负数段，下标越小平方越大，所以反向走才是递增顺序。
	left := right - 1
	result := make([]int, 0, len(nums))

	// Keep merging while either sequence still has an unused value.
	// 只要还有一段没用完，就继续归并。
	for left >= 0 || right < len(nums) {
		// Take from the negative side when the nonnegative side is exhausted,
		// or when its square is the smaller of the two candidates.
		// 非负数段已用完，或负数段的平方更小时，就从负数段取值。
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
