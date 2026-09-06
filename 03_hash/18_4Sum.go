package _3_hash

/*
题目描述 / Problem Description
给定整数数组 nums 和整数 target，找出所有和为 target 的不重复四元组。
每个四元组必须使用四个不同下标的元素。
Given an integer array nums and an integer target,
return all unique quadruplets whose sum equals target, using four distinct indices.

解题思路 / Solution Approach
先排序，使用两层循环固定前两个数，再在剩余区间使用左右指针寻找另外两个数。
各层都跳过重复值。当前实现用 int 求和，没有显式转换为 int64；整数范围必须足以容纳四数之和。
Sort first, fix two values with nested loops, and use two pointers for the remaining pair.
Skip duplicates at every level. This implementation sums with int, without an explicit int64 conversion;
its range must accommodate the four-value sum.

关键逻辑：为什么这样做 / Why This Works
固定 i、j 后，剩下的是有序区间的两数和：和太小时，即使用最大的右值也不足，
当前左值可排除；和太大时，即使用最小的左值也超出，当前右值可排除。
j 的去重只针对同一个 i：j=i+1 是该组第一次选择，不能跳过，否则会漏掉 [2,2,2,2] 这类合法答案。
i<j<left<right 保证四个下标互不相同。
After fixing i and j, solve a sorted two-sum problem: a sum too small eliminates
the current left value even with its largest partner; a sum too large eliminates
the right value even with its smallest partner. Deduplicate j only within
the current i group. j=i+1 is its first choice and must be allowed,
including equal values such as [2,2,2,2]. i<j<left<right guarantees distinct indices.

时间与空间复杂度 / Time and Space Complexity
n = len(nums)，r 为结果四元组数量。时间 O(n³)：固定两项的组合数为 O(n²)，
每组双指针扫描 O(n)。排序 O(n log n) 不改变总阶。辅助空间 O(log n) 包括 Go 排序栈，结果 O(r)，总额外空间 O(log n+r)。
For n values and r output quadruplets, time is O(n³): O(n²) fixed pairs each require an O(n) scan.
Sorting adds O(n log n). Auxiliary space is O(log n) for the Go sorting stack; output O(r), total extra space O(log n+r).
*/

import "sort"

/*
a + b + c + d = target
3Sum: fix 1 number, then use left/right to find 2 numbers.
3Sum：固定 1 个数，再用 left/right 找另外 2 个数。

4Sum: fix 2 numbers, then use left/right to find 2 numbers.
4Sum：固定 2 个数，再用 left/right 找另外 2 个数。

I sort the array first.
先排序，让重复值相邻，并让双指针可以根据总和大小移动。
Then I fix the first two numbers with two loops.
使用两层循环固定前两个数。
For the remaining part of the array, I use two pointers to find the other two numbers.
在剩余有序区间中使用双指针寻找另外两个数。
If the sum is smaller than the target, I move the left pointer to increase the sum.
If the sum is larger than the target, I move the right pointer to decrease the sum.
I skip duplicate values for each position to avoid duplicate quadruplets.
*/

func fourSum(nums []int, target int) [][]int {
	// Sorting is required for directional pointer movement and deduplication.
	// 排序是双指针定向移动和去重的前提。
	sort.Ints(nums)

	res := [][]int{}

	// i fixes the first number; three positions must remain after it.
	// i 固定第一个数，后面必须至少再留三个位置。
	for i := 0; i < len(nums)-3; i++ {
		// Skip duplicate first numbers.
		// 跳过重复的第一个数。
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		// j fixes the second number; two positions must remain for left and right.
		// j 固定第二个数，后面必须给 left 和 right 各留一个位置。
		for j := i + 1; j < len(nums)-2; j++ {
			// Compare with the previous j only within the current i group.
			// 只在当前 i 的范围内给 j 去重，因此条件是 j > i+1。
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			// Search for the remaining two numbers in the range after j.
			// 在 j 后面的区间寻找剩余两个数。
			left, right := j+1, len(nums)-1

			for left < right {
				sum := nums[i] + nums[j] + nums[left] + nums[right]

				if sum == target {
					// Record the quadruplet before moving past duplicate values.
					// 先记录当前四元组，再移动指针跳过重复值。
					res = append(res, []int{nums[i], nums[j], nums[left], nums[right]})

					b, c := nums[left], nums[right]

					// Skip all copies of the current third and fourth values.
					// 跳过第三、第四个数的所有重复值，避免重复答案。
					for left < right && nums[left] == b {
						left++
					}

					for left < right && nums[right] == c {
						right--
					}
				} else if sum < target {
					// Increase the sum by moving left to a larger value.
					// 总和太小，left 右移以增大总和。
					left++
				} else {
					// Decrease the sum by moving right to a smaller value.
					// 总和太大，right 左移以减小总和。
					right--
				}
			}
		}
	}

	return res
}
