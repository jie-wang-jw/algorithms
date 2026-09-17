package _1_array

/*
题目描述 / Problem Description
给定一个整数数组 nums 和一个整数 val，原地删除所有等于 val 的元素，并返回剩余元素的数量。
返回后，nums 的前 k 个位置应保存所有未删除元素。
Given an integer array nums and an integer val, remove every occurrence of val in place and
return the number of remaining elements. Afterward, the first k positions of nums must contain the retained elements.
*/

// 1. Brute-force shifting: overwrite each val by shifting the rest of the valid range left.
// 1. 暴力移位法：把有效区间内后续元素整体左移一位，覆盖掉待删除的 val。
// Time: O(n²), Space: O(1).
// 时间复杂度：O(n²)，空间复杂度：O(1)。
//
// size 是有效长度。删除 nums[i] 后后缀左移，新到 i 的元素还未检查，所以 i-- 抵消循环自增。
// size is the active length; after shifting left, i-- ensures the replacement at i is checked next.
func removeElementBruteForce(nums []int, val int) int {
	size := len(nums)
	for i := 0; i < size; i++ {
		if nums[i] == val {
			for j := i + 1; j < size; j++ {
				nums[j-1] = nums[j]
			}
			i--
			size--
		}
	}

	return size
}

// 2. Fast and slow pointers: fast reads every element, slow writes only the kept ones.
// 2. 快慢指针法：fast 读取每个元素，slow 只写入需要保留的元素。
// Time: O(n), Space: O(1).
// 时间复杂度：O(n)，空间复杂度：O(1)。
//
// nums[:slow] 始终是已保留的非 val 元素；fast 扫描原数组，只把合格元素写到 slow 后再递增 slow。
// nums[:slow] holds retained values; fast scans once and copies each non-val value to the next output position.
// slow<=fast 保证写入不覆盖未读元素，保留相对顺序；返回 slow，尾部内容不属于答案。
// slow<=fast protects unread values and preserves order; only the returned prefix belongs to the result.
func removeElementFastSlow(nums []int, val int) int {
	slow := 0

	for fast := range nums {
		if nums[fast] != val {
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
//
// left 找待删除项，right 找可保留项，用右端值填左端空位；[left,right] 之外都已处理。
// left finds a removal slot and right a keeper; copy the keeper into the slot and shrink the unresolved range.
// 相遇时也要检查该元素，所以使用 left<=right；返回有效长度 left，但不保证原相对顺序。
// left<=right processes the last candidate; left is the final length, and relative order is not preserved.
func removeElementTwoPointers(nums []int, val int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		for left <= right && nums[left] != val {
			left++
		}

		for left <= right && nums[right] == val {
			right--
		}
		if left < right {
			nums[left] = nums[right]
			left++
			right--
		}
	}

	return left
}
