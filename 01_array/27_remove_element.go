package _1_array

/*
题目描述 / Problem Description
给定一个整数数组 nums 和一个整数 val，原地删除所有等于 val 的元素，并返回剩余元素的数量。
返回后，nums 的前 k 个位置应保存所有未删除元素。
Given an integer array nums and an integer val, remove every occurrence of val in place and
return the number of remaining elements. Afterward, the first k positions of nums must contain the retained elements.

解题思路 / Solution Approach
本文件提供三种解法。题目允许改变元素顺序，只检查返回长度 k 和 nums[:k]，不要求真正缩短切片或清空尾部。
This file provides three solutions. Element order may change; only the returned length k and nums[:k] matter.
The slice need not be shortened or its trailing values cleared.

1. 暴力移位法 / Brute-force shifting — removeElementBruteForce
遇到 val 时，将有效区间内后续元素全部左移一位，并将有效长度减一。
When val is found, shift all later elements in the valid range one position left and reduce the valid length by one.
移位后必须重新检查当前位置，因为补过来的元素仍可能等于 val；遍历边界也必须使用更新后的有效长度。
Recheck the same position because its replacement may also equal val; the loop must use the updated valid length as its boundary.
时间 O(n²)，额外空间 O(1)，保留未删除元素的相对顺序。
Time O(n²), extra space O(1); preserves the relative order of retained elements.

2. 快慢指针法 / Fast and slow pointers — removeElementFastSlow
fast 扫描所有元素，slow 指向下一个写入位置。遇到非 val 元素时，将其写入 nums[slow]，然后 slow 加一。
fast scans every element, while slow marks the next write position. Copy each non-val element to nums[slow], then increment slow.
扫描结束后，nums[:slow] 保存所有有效元素，返回 slow。
After the scan, nums[:slow] contains all retained elements; return slow.
时间 O(n)，额外空间 O(1)，保留未删除元素的相对顺序。
Time O(n), extra space O(1); preserves the relative order of retained elements.

3. 左右双指针法 / Left and right pointers — removeElementTwoPointers
left 从左侧寻找 val，right 从右侧寻找非 val 元素；右侧的 val 可以直接跳过。
left searches from the front for val, while right searches from the back for a non-val element; trailing val elements can be skipped directly.
当 left < right 时，用 nums[right] 覆盖 nums[left]，然后两个指针向内移动。结束后返回 left，有效结果为 nums[:left]。
When left < right, overwrite nums[left] with nums[right] and move both pointers inward. Return left when finished; the valid result is nums[:left].
时间 O(n)，额外空间 O(1)，可能改变未删除元素的相对顺序；仅在需要填补左侧待删除位置时写入。
Time O(n), extra space O(1); may change the relative order of retained elements and writes only when filling a removal position on the left.

示例 / Example: nums = [3, 2, 2, 3, 4], val = 3
暴力法和快慢指针法：k = 3，有效前缀为 [2, 2, 4]。
Brute-force shifting and fast/slow pointers: k = 3, valid prefix [2, 2, 4].
左右双指针法：k = 3，有效前缀为 [4, 2, 2]。两种排列都符合题意。
Left/right pointers: k = 3, valid prefix [4, 2, 2]. Both orderings satisfy the problem.

关键逻辑：为什么这样做 / Why This Works
快慢指针：处理 fast 前，nums[:slow] 是已扫描部分中所有保留值，因此 slow 既是数量也是下一个写入位置。
始终 slow<=fast，写入只覆盖已处理位置或当前位置，不会破坏未来输入。
左右指针：nums[:left] 已全部有效，[left,right] 尚未处理。找到左侧 val 和右侧非 val 后，复制就填好了左侧空缺，
所以 left++；右侧值已搬入有效前缀，原位置应排除，所以 right--。指针交错时没有未处理元素，前缀长度 left 就是答案。这里是覆盖，不是交换。
Fast/slow: before processing fast, nums[:slow] contains every retained value already scanned.
Thus slow is both the count and next write index; slow<=fast prevents overwriting future input.
Left/right: nums[:left] is valid and [left,right] remains unprocessed. Copying the right survivor
fills the left removal slot, so advance left; its old right position must be excluded, so decrement right.
Once pointers cross, left is the valid prefix length. This is replacement, not a swap.

时间与空间复杂度 / Time and Space Complexity
n = len(nums)。暴力法最坏时间 O(n²)，多次移位的总成本可达 n+(n-1)+...+1；快慢指针和左右双指针均为 O(n)，
每个指针只单向移动。三种解法辅助空间均为 O(1)，修改输入并返回整数长度。
n = len(nums). Brute force takes O(n²) worst-case time due to repeated shifts;
both pointer solutions take O(n) because pointers move only forward or inward.
All three use O(1) auxiliary space and return an integer length.
*/

// 1. Brute-force shifting: overwrite each val by shifting the rest of the valid range left.
// 1. 暴力移位法：把有效区间内后续元素整体左移一位，覆盖掉待删除的 val。
// Time: O(n²), Space: O(1).
// 时间复杂度：O(n²)，空间复杂度：O(1)。
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

// 2. Fast and slow pointers: fast reads every element, slow writes only the kept ones.
// 2. 快慢指针法：fast 读取每个元素，slow 只写入需要保留的元素。
// Time: O(n), Space: O(1).
// 时间复杂度：O(n)，空间复杂度：O(1)。
func removeElementFastSlow(nums []int, val int) int {
	// slow points to the next write position in the kept result.
	// slow 指向保留结果中的下一个写入位置。
	slow := 0

	// fast scans every element; slow moves only when an element is kept.
	// fast 扫描每个元素；只有保留元素时 slow 才向前移动。
	for fast := range nums {
		if nums[fast] != val {
			// Copy the kept value to the front valid part of the slice.
			// 把需要保留的值写到切片前面的有效区域。
			nums[slow] = nums[fast]
			slow++
		}
	}

	return slow
}

// 3. Opposing pointers: fill each removal slot on the left with a survivor from the right.
// 3. 相向双指针法：用右侧需要保留的值，填补左侧待删除的位置。
// Time: O(n), Space: O(1); the relative order of kept elements may change.
// 时间复杂度：O(n)，空间复杂度：O(1)；保留元素的相对顺序可能改变。
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
