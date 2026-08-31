package _1_array

func removeElementBruteForce(nums []int, val int) int {
	// size is the length of the current valid part of nums.
	// size 表示 nums 当前仍然有效的区间长度。
	size := len(nums)
	for i := 0; i < size; i++ {
		if nums[i] == val {
			// Shift every following element one position to the left.
			// 把后面的元素全部向左移动一位，覆盖需要删除的元素。
			for j := i + 1; j < size; j++ {
				nums[j-1] = nums[j]
			}
			// Recheck index i because a new value has just been moved here.
			// i 位置刚被新元素覆盖，所以 i--，下一轮还要检查这个位置。
			i--
			// The valid range is now one element shorter.
			// 有效区间长度减少 1。
			size--
		}
	}

	return size
}

func removeElementFastSlow(nums []int, val int) int {
	// slow points to the next write position in the kept result.
	// slow 指向保留结果中的下一个写入位置。
	slow := 0

	// fast scans every element; slow moves only when an element is kept.
	// fast 扫描每个元素；只有保留元素时 slow 才向前移动。
	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] != val {
			// Copy the kept value to the front valid part of the slice.
			// 把需要保留的值写到切片前面的有效区域。
			nums[slow] = nums[fast]
			slow++
		}
	}

	return slow
}

func removeElementTwoPointers(nums []int, val int) int {
	// left searches for a value to remove; right searches for a value to keep.
	// left 从左找待删除值；right 从右找可以保留的值。
	left := 0
	right := len(nums) - 1

	for left <= right {
		// Values not equal to val are already in the correct valid area.
		// 不等于 val 的值已经位于有效区域，left 可以继续右移。
		for left <= right && nums[left] != val {
			left++
		}

		// Values equal to val at the end can be discarded directly.
		// 右侧等于 val 的值可以直接排除，right 向左移动。
		for left <= right && nums[right] == val {
			right--
		}
		// Replace the left-side val with a right-side value that should be kept.
		// 找到后，用右侧的非 val 覆盖左侧的 val。
		if left < right {
			nums[left] = nums[right]
			left++
			right--
		}
	}

	// [0, left) is the valid result, so left is also its length.
	// [0, left) 是最终有效区间，因此 left 也是新长度。
	return left
}
