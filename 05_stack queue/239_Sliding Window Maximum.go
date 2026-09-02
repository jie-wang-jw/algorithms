package _5_stack_queue

/*
题目描述 / Problem Description
给定整数数组 nums 和窗口大小 k，窗口从数组最左侧每次向右移动一位，返回每个窗口中的最大值。
Given an integer array nums and a window size k, move the window one position at a time from left to right and return the maximum value in every window.

解题思路 / Solution Approach
使用保存下标的单调递减队列。每轮删除队首过期下标，再删除队尾所有不大于当前值的下标并加入当前下标；队首始终对应当前窗口最大值。
Use a decreasing monotonic deque of indices. Remove expired indices from the front, remove values no greater than the current value from the back, then append the current index; the front is always the maximum.
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


I use a monotonic deque to store indices whose corresponding values are in decreasing order.
Before adding a new index, I remove expired indices from the front and remove smaller values from the back.
Therefore,the front of the deque always represents the maximum value in the current window.
*/

func maxSlidingWindow(nums []int, k int) []int {
	// deque stores indices, not values.
	// deque 保存数组下标，而不是直接保存数字。
	//
	// The corresponding values are kept in decreasing order:
	// nums[deque[0]] >= nums[deque[1]] >= ...
	// 这些下标对应的数字保持单调递减：
	// nums[deque[0]] >= nums[deque[1]] >= ...
	deque := make([]int, 0)

	// There are len(nums)-k+1 windows.
	// 一共有 len(nums)-k+1 个窗口。
	result := make([]int, 0, len(nums)-k+1)

	// i is the right boundary of the current window.
	// i 是当前窗口的右边界。
	for i := 0; i < len(nums); i++ {

		// Step 1: remove the front index if it has left the window.
		// 第一步：如果队头下标已经离开窗口，就将它删除。

		// Calculate the left boundary of the current window.
		// 计算当前窗口的左边界。
		windowLeft := i - k + 1
		// Remove the front index if it is outside the window.
		// 如果队头下标位于窗口左侧，说明它已经过期，需要删除。
		if len(deque) > 0 && deque[0] < windowLeft {
			deque = deque[1:]
		}

		// Step 2: remove all values from the back that are
		// less than or equal to the current value.
		// 第二步：删除队尾所有小于或等于当前数字的元素。
		//
		// They cannot become a future maximum because nums[i]
		// is larger and will stay in the window longer.
		// 它们以后不可能成为最大值，因为 nums[i] 更大，
		// 而且 nums[i] 会更晚离开窗口。
		for len(deque) > 0 &&
			//队尾对应的数字 <= 当前新进入窗口的数字
			nums[deque[len(deque)-1]] <= nums[i] {
			deque = deque[:len(deque)-1]
		}

		// Step 3: add the current index.
		// 第三步：将当前下标加入队尾。
		deque = append(deque, i)

		// A complete window is formed when i reaches k-1.
		// 当 i 到达 k-1 时，第一个完整窗口形成。
		if i >= k-1 {
			// The front always points to the maximum value.
			// 队头始终指向当前窗口中的最大值。
			result = append(result, nums[deque[0]])
		}
	}
	return result
}
