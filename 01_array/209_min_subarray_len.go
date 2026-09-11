package _1_array

/*
题目描述 / Problem Description
给定一个正整数 target 和一个由正整数组成的数组 nums，找出总和大于或等于 target 的最短连续子数组，
并返回其长度；如果不存在则返回 0。
Given a positive integer target and an array nums of positive integers,
return the minimum length of a contiguous subarray whose sum is at least target; return 0 if none exists.

解题思路 / Solution Approach
使用滑动窗口。右指针不断扩大窗口并累加元素；当窗口和达到 target 时，持续移动左指针缩小窗口，同时更新最短长度。
Use a sliding window. Expand the right boundary and add values; whenever the sum reaches target,
repeatedly shrink the left boundary while updating the minimum length.

关键逻辑：为什么这样做 / Why This Works
为什么不会漏掉最短窗口？数组元素都是正数，所以右端扩张只会增大和，左端收缩只会减小和。固定右端 j 时，
每次先记录合法窗口再收缩，直到不合法，便检查了当前仍可能改进答案的最短窗口。
已经丢弃的左端曾经对应一个合法窗口，以后再扩张只会更长，不会改进最短长度；若允许负数，这个理由就不成立。
Positive values make expansion increase the sum and shrinking decrease it. For each right endpoint,
record valid windows before shrinking until invalid. A discarded left endpoint already yielded a valid, shorter window,
so extending it later cannot improve the minimum. This reasoning does not hold with negative values.

时间与空间复杂度 / Time and Space Complexity
n = len(nums)。时间 O(n)：左右边界各最多前进 n 次，嵌套循环不是 O(n²)。辅助空间 O(1)，只维护窗口和与下标。
n = len(nums). Time O(n): each boundary advances at most n times, so the nested loops are not quadratic. Auxiliary space O(1) for the sum and indices.

补充解法 / Additional Approaches
minSubArrayLenBruteForce：固定左端点，向右累加，第一次达标即为该起点的最短答案。
正数保证继续右扩只会更长；时间 O(n²)，辅助空间 O(1)。
Fix each start and extend until the first qualifying sum: later ends cannot be shorter.
Time O(n²), auxiliary space O(1).
minSubArrayLenBinarySearch：prefix[j]-prefix[i]>=target 等价于 prefix[j]>=prefix[i]+target。
正数让前缀和递增，因此二分找第一个达标的 j，而不是任意一个 j。
时间 O(n log n)，辅助空间 O(n)；本题仍优先选择原有 O(n) 滑动窗口。
Positive values make prefix sums increasing. Lower-bound search finds the earliest qualifying end.
Time O(n log n), auxiliary space O(n); prefer the existing O(n) sliding window here.
*/

// 1. Sliding window: expand the right boundary, then shrink the left one while the sum qualifies.
// 1. 滑动窗口：右边界不断扩张，窗口和达标后再收缩左边界。
// Time: O(n), Space: O(1).
// 时间复杂度：O(n)，空间复杂度：O(1)。
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

// 2. Brute force: fix each start and extend right until the sum first qualifies.
// 2. 暴力解法：固定左端点，向右累加，第一次达标就停止。
// Time: O(n²), Space: O(1).
// 时间复杂度：O(n²)，空间复杂度：O(1)。
func minSubArrayLenBruteForce(target int, nums []int) int {
	// best uses an impossible length as its sentinel, exactly like result above.
	// best 同样使用不可能出现的长度作为哨兵初值。
	best := len(nums) + 1

	for left := range nums {
		sum := 0
		for right := left; right < len(nums); right++ {
			sum += nums[right]
			if sum >= target {
				// All values are positive, so extending further only makes this start longer.
				// 元素都是正数，继续右扩只会更长，因此当前长度就是该起点的最优解。
				best = min(best, right-left+1)
				break
			}
		}
	}

	if best > len(nums) {
		// The sentinel was never replaced, so no valid subarray exists.
		// 哨兵值从未被替换，说明不存在满足条件的子数组。
		return 0
	}

	return best
}

// 3. Prefix sums plus binary search: find the earliest end whose prefix sum is large enough.
// 3. 前缀和加二分查找：为每个起点二分出第一个达标的终点。
// Time: O(n log n), Space: O(n) for the prefix array.
// 时间复杂度：O(n log n)，空间复杂度：O(n)，由前缀和数组产生。
func minSubArrayLenBinarySearch(target int, nums []int) int {
	// prefix[i] is the sum of the first i values, so prefix[j]-prefix[i] is the sum of nums[i:j].
	// prefix[i] 表示前 i 个数之和，因此 prefix[j]-prefix[i] 就是 nums[i:j] 的和。
	prefix := make([]int, len(nums)+1)
	for i, value := range nums {
		prefix[i+1] = prefix[i] + value
	}

	best := len(nums) + 1

	for i := range nums {
		// Positive values make prefix strictly increasing, so a lower-bound search is valid.
		// 元素都是正数，前缀和严格递增，因此可以用二分查找第一个达标位置。
		left, right := i+1, len(prefix)

		// Find the first qualifying end in the half-open range [left, right).
		// 在左闭右开区间 [left, right) 中寻找第一个达标终点。
		for left < right {
			mid := left + (right-left)/2
			if prefix[mid]-prefix[i] >= target {
				// mid still qualifies, so it stays a candidate and becomes the new right bound.
				// mid 仍然达标，它本身还是候选答案，所以 right = mid。
				right = mid
			} else {
				// mid is too small, so the earliest qualifying end is after it.
				// mid 太小，第一个达标终点只能在它之后。
				left = mid + 1
			}
		}

		// left == len(prefix) means no end qualifies for this start.
		// left == len(prefix) 说明该起点没有任何达标终点。
		if left < len(prefix) {
			best = min(best, left-i)
		}
	}

	if best > len(nums) {
		return 0
	}

	return best
}
