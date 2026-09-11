package _1_array

/*
题目描述 / Problem Description
给定一个按升序排列的整数数组 nums 和目标值 target，返回 target 的下标；如果目标值不存在，则返回 -1。
Given an integer array nums sorted in ascending order and a target value, return the target's index or -1 if it is absent.

解题思路 / Solution Approach
使用二分查找，每次比较中间元素并排除一半搜索区间。
文件分别演示左闭右闭区间、左闭右开区间两种迭代写法，以及与左闭右闭完全等价的递归写法。
Use binary search, comparing the middle element and discarding half of the search range each time.
The file demonstrates the closed and half-open interval conventions, plus the recursive form of the closed-interval version.

关键逻辑：为什么这样做 / Why This Works
为什么能排除一半？数组有序，若 nums[mid]>target，则 mid 及其右边都不可能是答案；小于时同理排除左半段。
边界更新必须保持区间定义：[left,right] 排除 mid 后用 right=mid-1；[left,right) 的右边本来不包含，所以用 right=mid。
每轮都排除已检查的 mid，区间严格缩小；只有候选区间为空才返回 -1。
Sorted order makes every value at or right of mid too large when nums[mid]>target;
the symmetric rule applies when it is too small. Preserve the interval convention:
excluding mid requires right=mid-1 for [left,right], but right=mid for [left,right).
Each unsuccessful round removes mid and strictly shrinks the candidates.
Return -1 only when no candidates remain.

时间与空间复杂度 / Time and Space Complexity
n = len(nums)。search1 和 search2 最坏时间均为 O(log n)，每轮搜索区间减半；辅助空间 O(1)，迭代实现只保存边界和中点。
n = len(nums). Both search1 and search2 take O(log n) worst-case time because the interval halves each round, and O(1) auxiliary space for boundaries and midpoint.

补充解法：递归写法 / Additional Approach: Recursive Form
searchRecursive 使用与 search2 完全相同的左闭右闭不变量，不是新的查找原理。
每次递归排除 mid，并只在可能含目标的那一半继续；left>right 表示区间为空，返回 -1。
时间 O(log n)，递归深度也是 O(log n)，因此调用栈占 O(log n) 辅助空间，不如迭代版的 O(1)。
searchRecursive reuses search2's closed-interval invariant rather than a different search principle.
Each call excludes mid and recurses into the only half that can still contain the target; left>right means the range is empty, so return -1.
Time O(log n) with recursion depth O(log n), so the call stack costs O(log n) auxiliary space instead of the iterative O(1).
*/

// 1. Closed interval [left, right]: both boundaries may contain the target.
// 1. 左闭右闭区间 [left, right]：左右边界都可能是答案。
// Time: O(log n), Space: O(1).
// 时间复杂度：O(log n)，空间复杂度：O(1)。
func search2(nums []int, target int) int {
	// The initial search range covers every valid index.
	// 初始搜索范围包含数组的所有有效下标。
	left := 0
	right := len(nums) - 1

	// Use <= because left == right still leaves one candidate to check.
	// 使用 <=，因为 left == right 时仍有一个候选位置需要检查。
	for left <= right {
		// This form avoids the overflow risk of (left + right) / 2.
		// 这种写法避免 (left + right) / 2 可能产生的整数溢出。
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			// mid is too large, so discard mid and everything to its right.
			// mid 太大，因此丢弃 mid 以及它右侧的区间。
			right = mid - 1
		} else {
			// mid is too small, so discard mid and everything to its left.
			// mid 太小，因此丢弃 mid 以及它左侧的区间。
			left = mid + 1
		}
	}

	return -1
}

// 2. Half-open interval [left, right): left is included, right is excluded.
// 2. 左闭右开区间 [left, right)：left 包含在内，right 不包含在内。
// Time: O(log n), Space: O(1).
// 时间复杂度：O(log n)，空间复杂度：O(1)。
func search1(nums []int, target int) int {
	left := 0
	// right may equal len(nums) because it is outside the search interval.
	// right 可以等于 len(nums)，因为它本身不属于搜索区间。
	right := len(nums)

	// Stop when [left, right) becomes empty, which happens at left == right.
	// 当 left == right 时区间为空，所以循环条件是 left < right。
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			// mid is excluded, so the new half-open interval is [left, mid).
			// 排除 mid 后，新区间是 [left, mid)，因此 right = mid。
			right = mid
		} else {
			// mid is too small, so the new interval starts at mid + 1.
			// mid 太小，新区间从 mid + 1 开始。
			left = mid + 1
		}
	}

	return -1
}

// 3. Recursive form of the closed interval: same invariant, expressed with recursion.
// 3. 左闭右闭区间的递归写法：不变量完全相同，只是改用递归表达。
// Time: O(log n), Space: O(log n) for the recursion stack.
// 时间复杂度：O(log n)，空间复杂度：O(log n)，由递归调用栈产生。
func searchRecursive(nums []int, target int) int {
	// searchRange looks for target inside the closed interval [left, right].
	// searchRange 在左闭右闭区间 [left, right] 中查找 target。
	var searchRange func(left, right int) int

	searchRange = func(left, right int) int {
		// An empty interval means every candidate position has been excluded.
		// 区间为空，说明所有候选位置都已经被排除。
		if left > right {
			return -1
		}

		// This form avoids the overflow risk of (left + right) / 2.
		// 这种写法避免 (left + right) / 2 可能产生的整数溢出。
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}

		if nums[mid] > target {
			// mid is too large, so recurse into the closed interval [left, mid-1].
			// mid 太大，因此只在左闭右闭区间 [left, mid-1] 中继续递归。
			return searchRange(left, mid-1)
		}

		// mid is too small, so recurse into the closed interval [mid+1, right].
		// mid 太小，因此只在左闭右闭区间 [mid+1, right] 中继续递归。
		return searchRange(mid+1, right)
	}

	// The initial closed interval covers every valid index of nums.
	// 初始的左闭右闭区间覆盖 nums 的所有有效下标。
	return searchRange(0, len(nums)-1)
}
