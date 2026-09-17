package _1_array

/*
题目描述 / Problem Description
给定一个正整数 target 和一个由正整数组成的数组 nums，找出总和大于或等于 target 的最短连续子数组，
并返回其长度；如果不存在则返回 0。
Given a positive integer target and an array nums of positive integers,
return the minimum length of a contiguous subarray whose sum is at least target; return 0 if none exists.
*/

// 1. Sliding window: expand the right boundary, then shrink the left one while the sum qualifies.
// 1. 滑动窗口：右边界不断扩张，窗口和达标后再收缩左边界。
// Time: O(n), Space: O(1).
// 时间复杂度：O(n)，空间复杂度：O(1)。
//
// i、j 是窗口两端，sum 是 nums[i:j+1] 的和。元素均为正数：右扩使和增加，左缩使和减少。
// i and j bound the window; positive values make expansion increase its sum and shrinking decrease it.
// 每次达标先记录长度再左缩，直到不达标；已舍弃的左端再配更远的右端只会更长，不会漏掉更优解。
// Record each valid length before shrinking; extending a discarded start later cannot yield a shorter answer.
// result=l+1 表示尚未找到答案；左右指针各最多走 n 次，所以内外循环合计 O(n)。
// result=l+1 means no answer yet; each pointer advances at most n times, giving O(n) total work.
func minSubArrayLen(target int, nums []int) int {
	i := 0
	l := len(nums)
	sum := 0
	result := l + 1

	for j := range l {
		sum += nums[j]
		for sum >= target {
			subLength := j - i + 1
			if subLength < result {
				result = subLength
			}

			sum -= nums[i]
			i++
		}
	}
	if result == l+1 {
		return 0
	}
	return result
}

// 2. Brute force: fix each start and extend right until the sum first qualifies.
// 2. 暴力解法：固定左端点，向右累加，第一次达标就停止。
// Time: O(n²), Space: O(1).
// 时间复杂度：O(n²)，空间复杂度：O(1)。
//
// 固定 left 后逐步右扩；第一次达标就是这个起点的最短答案，继续右扩只会更长。
// For each left endpoint, the first qualifying end is shortest; later ends only increase the length.
// best=len(nums)+1 是无解哨兵，遍历后仍未更新就返回 0。
// best=len(nums)+1 is an impossible length; return 0 if no candidate replaces it.
func minSubArrayLenBruteForce(target int, nums []int) int {
	best := len(nums) + 1

	for left := range nums {
		sum := 0
		for right := left; right < len(nums); right++ {
			sum += nums[right]
			if sum >= target {
				best = min(best, right-left+1)
				break
			}
		}
	}

	if best > len(nums) {
		return 0
	}

	return best
}

// 3. Prefix sums plus binary search: find the earliest end whose prefix sum is large enough.
// 3. 前缀和加二分查找：为每个起点二分出第一个达标的终点。
// Time: O(n log n), Space: O(n) for the prefix array.
// 时间复杂度：O(n log n)，空间复杂度：O(n)，由前缀和数组产生。
//
// prefix[k] 是前 k 项之和，所以 nums[i:j] 达标等价于 prefix[j]>=prefix[i]+target。
// prefix[k] sums the first k values; nums[i:j] qualifies iff prefix[j]>=prefix[i]+target.
// 正数使 prefix 严格递增；在 [i+1,len(prefix)) 二分第一个达标位置：达标令 right=mid，否则 left=mid+1。
// Positive values make prefix increasing; lower-bound search retains mid when valid and excludes it otherwise.
// left==len(prefix) 表示此起点无解；否则长度为 left-i，取所有起点的最小值。
// left==len(prefix) means no qualifying end; otherwise minimize left-i over all starts.
func minSubArrayLenBinarySearch(target int, nums []int) int {
	prefix := make([]int, len(nums)+1)
	for i, value := range nums {
		prefix[i+1] = prefix[i] + value
	}

	best := len(nums) + 1

	for i := range nums {
		left, right := i+1, len(prefix)

		for left < right {
			mid := left + (right-left)/2
			if prefix[mid]-prefix[i] >= target {
				right = mid
			} else {
				left = mid + 1
			}
		}

		if left < len(prefix) {
			best = min(best, left-i)
		}
	}

	if best > len(nums) {
		return 0
	}

	return best
}
