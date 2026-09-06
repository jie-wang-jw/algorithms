package _3_hash

/*
题目描述 / Problem Description
给定整数数组 nums，找出所有和为 0 的不重复三元组。每个三元组必须使用三个不同下标的元素。
Given an integer array nums, return all unique triplets whose sum is zero.
Each triplet must use elements at three distinct indices.

解题思路 / Solution Approach
先排序数组，再固定第一个数，并用左右指针寻找另外两个数。根据三数之和移动指针，同时跳过重复值以避免重复三元组。
Sort the array, fix the first value, and use two pointers to find the other two.
Move pointers according to the sum and skip duplicate values to avoid duplicate triplets.

时间与空间复杂度 / Time and Space Complexity
n = len(nums)，r 为结果三元组数量。时间 O(n²)：排序 O(n log n)，固定一个数后每次双指针扫描 O(n)。
辅助空间 O(log n)，计入 Go 排序调用栈；结果占 O(r)，总额外空间 O(log n+r)。排序会修改输入。
For n values and r output triplets, time is O(n²): O(n log n) sorting plus linear scans for each fixed value.
Auxiliary space O(log n) includes the Go sorting stack; output takes O(r), for O(log n+r) total extra space. Sorting modifies the input.
*/

import "sort"

/*
Sort + fix one number + two pointers.
排序 + 固定一个数 + 双指针。

Sort the array first.
先排序，使相同数字相邻，并让指针移动具有明确的增减方向。

Then I fix one number and use two pointers on the remaining range.
固定第一个数 a，再在右侧剩余区间使用 left 和 right 找另外两个数。

If the sum is too small, I move the left pointer to increase it.
如果总和太小，left 右移，让总和变大。

If the sum is too large, I move the right pointer to decrease it.
如果总和太大，right 左移，让总和变小。

I also skip duplicate values to avoid returning the same triplet multiple times.
跳过重复值，避免返回相同的三元组。
*/

func threeSum(nums []int) [][]int {
	// Sorting puts duplicates together and gives pointer movement a direction.
	// 排序后相同数字会相邻，也方便根据总和大小移动双指针。
	sort.Ints(nums)

	res := [][]int{}

	/*
		i fixes the first number a. Stop before the last two positions because left and right still need one position each.
		i 固定第一个数 a。i 最多到倒数第三个位置，因为后面还要给 left 和 right 各留一个位置。
	*/
	for i := 0; i < len(nums)-2; i++ {
		a := nums[i]

		// If a is positive, all later values are also positive, so the sum cannot be 0.
		// 如果 a 已经大于 0，后面的数字只会更大，总和不可能再等于 0。
		if a > 0 {
			break
		}

		// Skip a duplicate a because it would generate the same triplets again.
		// 当前 a 和前一个 a 相同时跳过，否则会重复生成相同三元组。
		if i > 0 && a == nums[i-1] {
			continue
		}

		// Search for b and c in the sorted range to the right of i.
		// 在 i 右侧的有序区间中寻找 b 和 c。
		left, right := i+1, len(nums)-1

		for left < right {
			b, c := nums[left], nums[right]
			sum := a + b + c

			if sum == 0 {
				// A valid triplet has been found.
				// 找到一个满足 a+b+c=0 的三元组。
				res = append(res, []int{a, b, c})

				// Skip every copy of b and c because this value combination is already recorded.
				// 当前 b、c 组合已经记录，跳过它们的所有重复值以避免重复答案。
				for left < right && nums[left] == b {
					left++
				}

				for left < right && nums[right] == c {
					right--
				}
			} else if sum > 0 {
				// The sum is too large; move right leftward to use a smaller c.
				// 总和太大，right 左移，换一个更小的 c。
				right--
			} else {
				// The sum is too small; move left rightward to use a larger b.
				// 总和太小，left 右移，换一个更大的 b。
				left++
			}
		}
	}

	return res
}
