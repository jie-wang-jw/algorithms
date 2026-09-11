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

关键逻辑：为什么这样做 / Why This Works
固定 a 后，若 a+b+c<0，当前 c 已是候选区间最大值，保留 b 再选更小的 c 也不可能成功，所以可以排除当前 b，
移动 left。和过大时，当前 b 已是最小值，保留 c 也无法成功，所以移动 right。找到答案后，相同 b 或 c 只会再次得到同一值组合，
才跳过重复值。i<left<right 保证不同下标，去重限制的是答案值组合，不是禁止答案含相同数字。
For fixed a, if a+b+c<0, c is already the largest available partner,
so this b cannot work with any remaining c; advance left. If the sum is too large, b is already the smallest partner,
so discard this c by decrementing right. After a match, repeated b or c values only reproduce the same triplet.
i<left<right ensures distinct indices; deduplication does not forbid equal values within a triplet.

时间与空间复杂度 / Time and Space Complexity
n = len(nums)，r 为结果三元组数量。时间 O(n²)：排序 O(n log n)，固定一个数后每次双指针扫描 O(n)。
辅助空间 O(log n)，计入 Go 排序调用栈；结果占 O(r)，总额外空间 O(log n+r)。排序会修改输入。
For n values and r output triplets, time is O(n²): O(n log n) sorting plus linear scans for each fixed value.
Auxiliary space O(log n) includes the Go sorting stack; output takes O(r), for O(log n+r) total extra space. Sorting modifies the input.

补充解法：固定一个数 + 哈希 / Alternative: Fix One Value and Hash
threeSumHash 排序后固定 a，在其后扫描 c，用集合查找此前出现过的 b=-a-c。
先查询再插入，保证 b 与 c 来自不同下标。三个值天然有序，用 [3]int 集合去掉重复答案。
同样可以在 a>0 时提前结束，并跳过重复的 a。
平均时间 O(n²)，辅助空间 O(n+r)，r 是不同答案数，输出另占 O(r)；修改输入顺序。
相比双指针，这版能复用“两数之和”的思路，但需要额外集合。
After sorting, fix a and scan c, looking for an earlier b=-a-c before inserting c.
Query-before-insert ensures distinct indices; sorted triples form canonical deduplication keys.
It can also stop early once a>0 and skip duplicate values of a.
Expected time O(n²), auxiliary space O(n+r) plus O(r) output; input is sorted in place.
This reuses Two Sum, but needs more storage than two pointers.
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

// 1. 排序 + 双指针：推荐
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

// 2. 固定一个数 + 哈希：复用两数之和
func threeSumHash(nums []int) [][]int {
	// Sorting is not needed to find triplets, but it makes each answer come out in ascending order.
	// 排序不是找答案的必要条件，但能让每个答案本身保持升序，方便当作去重的 key。
	sort.Ints(nums)

	result := [][]int{}

	// recorded stores answers that have already been appended, keyed by their sorted values.
	// recorded 以有序三元组为 key，保存已经加入结果的答案。
	recorded := make(map[[3]int]bool)

	for i, a := range nums {
		// Later values are all at least as large as a, so a positive a can never reach 0.
		// 后面的数字都不小于 a，因此 a 大于 0 时总和不可能再等于 0。
		if a > 0 {
			break
		}

		// A repeated a would only regenerate the triplets found for the previous a.
		// 重复的 a 只会重新生成上一个 a 已经找到的三元组。
		if i > 0 && a == nums[i-1] {
			continue
		}

		// seen holds the values between i and the current c, which are the candidates for b.
		// seen 保存位于 i 和当前 c 之间的数值，它们是 b 的候选。
		seen := make(map[int]bool)

		for _, c := range nums[i+1:] {
			// a + b + c = 0, so the partner we need is b = -a-c.
			// 要满足 a+b+c=0，需要的另一个数就是 b = -a-c。
			b := -a - c

			if seen[b] {
				// b sits at an index between i and c, so the key is already ascending.
				// b 所在的下标位于 i 和 c 之间，因此 key 本身就是升序的。
				key := [3]int{a, b, c}
				if !recorded[key] {
					recorded[key] = true
					result = append(result, []int{a, b, c})
				}
			}

			// Insert after lookup so one position cannot supply both b and c.
			// 查询之后再插入，避免同一位置同时充当 b 和 c。
			seen[c] = true
		}
	}

	return result
}
