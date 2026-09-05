package _5_stack_queue

/*
题目描述 / Problem Description
给定整数数组 nums 和窗口大小 k，窗口从数组最左侧每次向右移动一位，返回每个窗口中的最大值。
Given an integer array nums and a window size k, move the window one position at a time from left to right and return the maximum value in every window.

解题思路 / Solution Approach
使用保存下标的单调递减队列。每轮删除队首过期下标，再删除队尾所有不大于当前值的下标并加入当前下标；队首始终对应当前窗口最大值。
Use a decreasing monotonic deque of indices. Remove expired indices from the front, remove values no greater than the current value from the back, then append the current index; the front is always the maximum.

时间与空间复杂度 / Time and Space Complexity
n = len(nums)，k 为窗口长度。时间 O(n)，每个下标最多入队、出队一次；两个内层循环的总工作量是线性的。辅助队列空间 O(k)，返回结果 O(n-k+1)，包含结果总空间 O(n)。切片 append 的分配与复制按均摊计算。
For n values and window size k, time is O(n): each index enters and leaves the deque at most once, so the inner loops have linear aggregate work. Auxiliary deque space O(k), output O(n-k+1), total O(n). Slice allocation and copying are amortized.
*/

/*滑动窗口 + 单调队列
Sliding window + monotonic deque

左边 = 队头 front
右边 = 队尾 back
新下标从右边加入：
过期下标从左边删除：

三个固定步骤
第一步：删除已经离开窗口的队头
第二步：删除队尾所有不大于当前值的元素
第三步：加入当前下标

窗口数量约为 n
每个窗口检查 k 个数字

每轮的三个步骤 / Three Steps per Iteration
第一步：删除过期下标 / Remove Expired Indices
当前窗口的有效下标范围是：
The current window’s valid index range is:
[i-k+1, i]
如果队首下标小于 i-k+1，说明它已经离开窗口，需要删除。
If the front index is smaller than i-k+1, it has left the window and must be removed.
等价条件：
Equivalent condition:
deque[0] <= i-k
第二步：维护单调递减 / Maintain Decreasing Order
加入 nums[i] 前，从队尾删除所有小于或等于 nums[i] 的元素。
Before adding nums[i], remove every value from the back that is less than or equal to nums[i].
原因是这些元素以后不可能成为窗口最大值：
Those elements can never become a future window maximum because:
- nums[i] 大于或等于它们；
- nums[i] is greater than or equal to them;
- nums[i] 比它们更晚离开窗口。
- nums[i] leaves the window later.
因此可以永久删除这些较弱的候选元素。
Therefore, these weaker candidates can be discarded permanently.
第三步：加入当前下标 / Add the Current Index
把 i 加入队尾。
Append i to the back.
当第一个完整窗口形成后，将队首对应的数字加入结果。
Once a complete window has formed, append the value represented by the deque front to the result.


I use a monotonic deque to store indices whose corresponding values are in decreasing order.
Before adding a new index, I remove expired indices from the front and remove smaller values from the back.
Therefore,the front of the deque always represents the maximum value in the current window.
*/

func maxSlidingWindow(nums []int, k int) []int {
	// deque stores indices instead of values.
	// deque 保存数组下标，而不是直接保存数字。
	//
	// The corresponding values remain in decreasing order.
	// 这些下标对应的数字保持单调递减。
	deque := make([]int, 0, k)

	// There are len(nums)-k+1 complete windows.
	// 一共有 len(nums)-k+1 个完整窗口。
	result := make([]int, 0, len(nums)-k+1)

	// i is the right boundary of the current window.
	// i 是当前窗口的右边界。
	for i := range nums {
		// Step 1: remove indices that have left the window.
		// 第一步：删除已经离开当前窗口的下标。
		//
		// The valid window starts at i-k+1, so indices <= i-k are expired.
		// 当前窗口从 i-k+1 开始，因此下标 <= i-k 的元素已经过期。
		for len(deque) > 0 && deque[0] <= i-k {
			deque = deque[1:]
		}

		// Step 2: remove weaker candidates from the back.
		// 第二步：从队尾删除不可能成为最大值的候选元素。
		//
		// nums[i] is at least as large and will leave the window later.
		// nums[i] 不小于这些元素，而且会更晚离开窗口。
		for len(deque) > 0 &&
			nums[deque[len(deque)-1]] <= nums[i] {
			deque = deque[:len(deque)-1]
		}

		// Step 3: append the current index.
		// 第三步：将当前下标加入队尾。
		deque = append(deque, i)

		// The first complete window is formed when i reaches k-1.
		// 当 i 到达 k-1 时，第一个完整窗口形成。
		if i >= k-1 {
			// The front always represents the current maximum.
			// 队首始终对应当前窗口的最大值。
			result = append(result, nums[deque[0]])
		}
	}

	return result
}
