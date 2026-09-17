package _5_stack_queue

import "container/heap"

/*
239. 滑动窗口最大值 / Sliding Window Maximum

题目描述 / Problem Description
给定整数数组 nums 和窗口大小 k，窗口从数组最左侧每次向右移动一位，返回每个窗口中的最大值。
Given an integer array nums and a window size k, move the window one position at a
time from left to right and return the maximum value in every window.
*/

// 1. Monotonic decreasing deque (recommended): drop expired fronts and weaker backs so the front is the window max.
// 1. 单调递减队列：推荐；删除过期队首和不大于当前值的队尾，队首始终是窗口最大值。
// Time: O(n), Space: O(k) for the deque plus O(n-k+1) output.
// 时间复杂度：O(n)，空间复杂度：队列 O(k)，输出 O(n-k+1)。
//
// deque 存下标：从队首到队尾下标递增、值严格递减。下标<=i-k 已离开窗口，先从队首移除。
// deque stores increasing indices with strictly decreasing values; remove indices <=i-k because they have expired.
// 新值不小于队尾值时，旧队尾既更小又更早过期，今后不可能成为最大值，可永久丢弃。
// A new value at least as large dominates the older tail and expires later, so that tail can never be needed again.
// 窗口满后队首就是最大值；每项最多入队、出队各一次。要求 1<=k<=len(nums)。
// Once the window is full, its maximum is at the front; each index enters and leaves once. Require 1<=k<=len(nums).
func maxSlidingWindow(nums []int, k int) []int {
	deque := make([]int, 0, k)

	result := make([]int, 0, len(nums)-k+1)

	for i := range nums {
		for len(deque) > 0 && deque[0] <= i-k {
			deque = deque[1:]
		}

		for len(deque) > 0 &&
			nums[deque[len(deque)-1]] <= nums[i] {
			deque = deque[:len(deque)-1]
		}

		deque = append(deque, i)

		if i >= k-1 {
			result = append(result, nums[deque[0]])
		}
	}

	return result
}

// 2. Brute force per window: rescan each [left, left+k-1] without reusing the previous window.
// 2. 逐窗口暴力扫描：每个窗口重新扫描 k 个元素，不复用上一窗口的信息。
// Time: O((n-k+1)k), Space: O(1) auxiliary plus O(n-k+1) output.
// 时间复杂度：O((n-k+1)k)，空间复杂度：辅助 O(1)，输出 O(n-k+1)。
//
// 逐个扫描长度 k 的窗口 [left,left+k)，用首元素初始化最大值以支持全负数；要求 1<=k<=len(nums)。
// Scan each length-k window and initialize from its first value to handle negatives; require 1<=k<=len(nums).
func maxSlidingWindowBruteForce(nums []int, k int) []int {
	result := make([]int, 0, len(nums)-k+1)

	for left := 0; left+k <= len(nums); left++ {
		best := nums[left]

		for j := left + 1; j < left+k; j++ {
			best = max(best, nums[j])
		}

		result = append(result, best)
	}

	return result
}

// windowMaxHeap is a max-heap of [value, index] pairs ordered by value, not a fully sorted slice.
// windowMaxHeap 是按数值维护父子大小关系的大顶堆，保存 [值, 下标]，不是完全有序的切片。
type windowMaxHeap [][2]int

// Report the heap size so container/heap can locate the last slot.
// 返回堆中的元素数量，供 container/heap 定位末尾位置。
// Time: O(1), Space: O(1).
// 时间复杂度：O(1)，空间复杂度：O(1)。
func (h windowMaxHeap) Len() int {
	return len(h)
}

// Compare values only so the larger one sits closer to the root, forming a max-heap.
// 只比较数值，让较大值更接近堆顶，从而形成大顶堆。
// Time: O(1), Space: O(1).
// 时间复杂度：O(1)，空间复杂度：O(1)。
func (h windowMaxHeap) Less(i, j int) bool {
	return h[i][0] > h[j][0]
}

// Exchange two heap slots while container/heap sifts.
// 在上浮和下沉过程中交换堆中的两个元素。
// Time: O(1), Space: O(1).
// 时间复杂度：O(1)，空间复杂度：O(1)。
func (h windowMaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Append the new pair to the backing slice; heap.Push then sifts it up with Less and Swap.
// 把新元素追加到底层切片末尾；heap.Push 随后用 Less 和 Swap 上浮。
// Time: amortized O(1), worst O(n) for append; backing storage O(n). heap.Push adds O(log n) sift work.
// 追加均摊 O(1)，扩容时 O(n)；底层存储 O(n)。外层 heap.Push 另有 O(log n) 上浮操作。
//
// 这里只向底层切片追加；外层 heap.Push 随后上浮修复堆序。追加均摊 O(1)，扩容时单次 O(n)。
// This hook only appends; heap.Push then sifts upward to restore heap order. Append is amortized O(1), or O(n) when reallocating.
func (h *windowMaxHeap) Push(x any) {
	*h = append(*h, x.([2]int))
}

// Remove the last backing-slice element after container/heap has already moved the root there.
// 删除底层切片末尾元素；container/heap 调用前已把原堆顶移到末尾并修好剩余堆。
// Time: O(1), Space: O(1).
// 时间复杂度：O(1)，空间复杂度：O(1)。
func (h *windowMaxHeap) Pop() any {
	old := *h
	last := len(old) - 1

	value := old[last]

	*h = old[:last]

	return value
}

// 3. Max-heap with lazy deletion: insert each index, then pop expired roots only; auxiliary space is O(n), not O(k).
// 3. 惰性删除的大顶堆：每轮插入当前下标，只弹出过期堆顶；辅助空间是 O(n) 而不是 O(k)。
// Time: O(n log n), Space: O(n) for the heap because lazy deletion can retain expired entries.
// 时间复杂度：O(n log n)，空间复杂度：O(n)，由堆产生，因为惰性删除让堆规模由 n 而不是 k 决定。
//
// 堆存 [值,下标]，按值降序；只要堆顶下标过期就弹出，直到堆顶属于当前窗口，此时它才是窗口最大值。
// Store [value,index] in a max-heap; discard expired tops until the maximum belongs to the current window.
// 堆内其他过期项暂留，可能积累到 O(n)，不能声称空间 O(k)；要求 1<=k<=len(nums)。
// Expired non-top entries remain and may occupy O(n) space, not O(k); require valid k.
func maxSlidingWindowHeap(nums []int, k int) []int {
	h := &windowMaxHeap{}

	result := make([]int, 0, len(nums)-k+1)

	for i, value := range nums {
		heap.Push(h, [2]int{value, i})

		for (*h)[0][1] <= i-k {
			heap.Pop(h)
		}

		if i >= k-1 {
			result = append(result, (*h)[0][0])
		}
	}

	return result
}

// 4. Block prefix and suffix maxima: a length-k window spans at most two blocks, so max(suffix[left], prefix[right]) is the answer.
// 4. 分块预处理前后缀最大值：长度 k 的窗口最多跨两块，因此 max(suffix[left], prefix[right]) 就是窗口最大值。
// Time: O(n), Space: O(n) for the prefix and suffix arrays.
// 时间复杂度：O(n)，空间复杂度：O(n)，由前后缀数组产生。
//
// 按 k 个元素分块；prefix[i] 是本块开头到 i 的最大值，suffix[i] 是 i 到本块末尾的最大值。
// Partition into length-k blocks; prefix[i] and suffix[i] hold maxima from the block start and to its end.
// 任意长度 k 的窗口恰跨至多两块，所以最大值为 max(suffix[left],prefix[right])；要求有效 k。
// A length-k window spans at most two blocks, so its maximum is max(suffix[left],prefix[right]); require valid k.
func maxSlidingWindowBlocks(nums []int, k int) []int {
	prefix, suffix := make([]int, len(nums)), make([]int, len(nums))

	for i, value := range nums {
		if i%k == 0 {
			prefix[i] = value
		} else {
			prefix[i] = max(prefix[i-1], value)
		}
	}

	for i := len(nums) - 1; i >= 0; i-- {
		if i == len(nums)-1 || (i+1)%k == 0 {
			suffix[i] = nums[i]
		} else {
			suffix[i] = max(suffix[i+1], nums[i])
		}
	}

	result := make([]int, 0, len(nums)-k+1)

	for left := 0; left+k <= len(nums); left++ {
		right := left + k - 1

		result = append(result, max(suffix[left], prefix[right]))
	}

	return result
}
