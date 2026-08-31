package _1_array

func minSubArrayLen(target int, nums []int) int {
	// i is the left boundary of the current sliding window.
	// i 是当前滑动窗口的左边界。
	i := 0
	l := len(nums)  // Length of the input array. / 输入数组的长度。
	sum := 0        // Sum of the current window nums[i:j+1]. / 当前窗口 nums[i:j+1] 的元素和。
	result := l + 1 // Use an impossible length as the initial answer. / 用不可能出现的长度作为初始答案。

	// j expands the right boundary one element at a time.
	// j 每次向右移动一位，扩张窗口右边界。
	for j := 0; j < l; j++ {
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
