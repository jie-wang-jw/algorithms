package _3_hash

/*
题目描述 / Problem Description
给定整数数组 nums 和整数 target，找出所有和为 target 的不重复四元组。
每个四元组必须使用四个不同下标的元素。
Given an integer array nums and an integer target,
return all unique quadruplets whose sum equals target, using four distinct indices.

解题思路 / Solution Approach
先排序，使用两层循环固定前两个数，再在剩余区间使用左右指针寻找另外两个数。
各层都跳过重复值，并在两层循环上做剪枝。当前实现用 int 求和，没有显式转换为 int64；
整数范围必须足以容纳四数之和。
Sort first, fix two values with nested loops, and use two pointers for the remaining pair.
Skip duplicates at every level and prune both loops. This implementation sums with int,
without an explicit int64 conversion; its range must accommodate the four-value sum.

关键逻辑：为什么这样做 / Why This Works
固定 i、j 后，剩下的是有序区间的两数和：和太小时，即使用最大的右值也不足，
当前左值可排除；和太大时，即使用最小的左值也超出，当前右值可排除。
j 的去重只针对同一个 i：j=i+1 是该组第一次选择，不能跳过，否则会漏掉 [2,2,2,2] 这类合法答案。
i<j<left<right 保证四个下标互不相同。
剪枝要同时判断 nums[i]>target 和 nums[i]>=0：只有当前缀和非负时，后面更大的数才无法把总和拉回 target；
target 为负时单看 nums[i]>target 会剪掉正确答案，例如 target=-10、nums=[-9,-3,-2,4] 中的 nums[i]=-9。
After fixing i and j, solve a sorted two-sum problem: a sum too small eliminates
the current left value even with its largest partner; a sum too large eliminates
the right value even with its smallest partner. Deduplicate j only within
the current i group. j=i+1 is its first choice and must be allowed,
including equal values such as [2,2,2,2]. i<j<left<right guarantees distinct indices.
Pruning must test both nums[i]>target and nums[i]>=0: only a nonnegative prefix guarantees that
the larger later values cannot pull the sum back down. With a negative target, testing nums[i]>target
alone would discard valid answers, as with nums[i]=-9 for target=-10 and nums=[-9,-3,-2,4].

时间与空间复杂度 / Time and Space Complexity
n = len(nums)，r 为结果四元组数量。时间 O(n³)：固定两项的组合数为 O(n²)，
每组双指针扫描 O(n)。排序 O(n log n) 不改变总阶。辅助空间 O(log n) 包括 Go 排序栈，结果 O(r)，总额外空间 O(log n+r)。
For n values and r output quadruplets, time is O(n³): O(n²) fixed pairs each require an O(n) scan.
Sorting adds O(n log n). Auxiliary space is O(log n) for the Go sorting stack; output O(r), total extra space O(log n+r).

补充解法：固定两个数 + 哈希 / Alternative: Fix Two Values and Hash
fourSumHash 排序后固定 a、b，在剩余后缀中用“先查后存”找到 c、d。
递增下标保证四个位置不同；用有序 [4]int 去重，而不是仅给输入去重（重复值可能参与合法答案）。
平均时间 O(n³)，辅助空间 O(n+r)，输出 O(r)，r 为答案数量；修改输入顺序。
中间加减使用 int64，避免四数和在 32 位 int 上溢出；最后返回原范围的 int 值。
Sort, fix a and b, then query before insertion for c and d in the suffix.
Increasing indices prevent reuse; canonical tuple keys deduplicate answers without removing necessary repeated values.
Expected time O(n³), auxiliary space O(n+r), output O(r); sorts the input.
Use int64 for intermediate arithmetic to avoid 32-bit int overflow.
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

// 1. 排序 + 双层循环 + 双指针 + 剪枝：推荐
func fourSum(nums []int, target int) [][]int {
	// Sorting is required for directional pointer movement and deduplication.
	// 排序是双指针定向移动和去重的前提。
	sort.Ints(nums)

	res := [][]int{}

	// i fixes the first number; three positions must remain after it.
	// i 固定第一个数，后面必须至少再留三个位置。
	for i := 0; i < len(nums)-3; i++ {
		/*
			Pruning: nums[i] > target together with nums[i] >= 0 means the other three values are also
			at least nums[i] >= 0, so every remaining sum stays above target.
			剪枝：nums[i] > target 且 nums[i] >= 0 时，另外三个数都不小于 nums[i] >= 0，
			因此后面所有组合的总和都会大于 target。

			target 可能是负数，所以不能只判断 nums[i] > target：例如 target=-10、nums[i]=-9 时，
			后面仍然可能出现更小的负数把总和拉回 target。
			The nums[i] >= 0 test is required because target can be negative.
		*/
		if nums[i] > target && nums[i] >= 0 {
			break
		}

		// Skip duplicate first numbers.
		// 跳过重复的第一个数。
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		// j fixes the second number; two positions must remain for left and right.
		// j 固定第二个数，后面必须给 left 和 right 各留一个位置。
		for j := i + 1; j < len(nums)-2; j++ {
			// The same pruning one level down: nums[i]+nums[j] >= 0 forces nums[j] >= 0,
			// so the remaining two values cannot decrease the sum.
			// 同样的剪枝下移一层：nums[i]+nums[j] >= 0 说明 nums[j] >= 0，
			// 剩下两个数都不小于它，无法把总和拉低。
			if nums[i]+nums[j] > target && nums[i]+nums[j] >= 0 {
				break
			}

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

// 2. 固定两个数 + 哈希：复用两数之和
func fourSumHash(nums []int, target int) [][]int {
	// Sorting keeps each answer ascending, which makes it usable as a deduplication key.
	// 排序让每个答案本身保持升序，可以直接当作去重的 key。
	sort.Ints(nums)

	result := [][]int{}

	// recorded stores answers that have already been appended, keyed by their sorted values.
	// recorded 以有序四元组为 key，保存已经加入结果的答案。
	recorded := make(map[[4]int]bool)

	for i := 0; i < len(nums)-3; i++ {
		// A repeated first number would only regenerate the same quadruplets.
		// 重复的第一个数只会重新生成相同的四元组。
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j := i + 1; j < len(nums)-2; j++ {
			// j = i+1 is this group's first choice, so deduplicate j only for j > i+1.
			// j = i+1 是本组的第一次选择，因此只在 j > i+1 时给 j 去重。
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			// seen holds the values between j and the current k, which are candidates for c.
			// seen 保存位于 j 和当前 k 之间的数值，它们是 c 的候选。
			seen := make(map[int64]bool)

			for k := j + 1; k < len(nums); k++ {
				// Compute in int64 so the four-value sum cannot overflow a 32-bit int.
				// 用 int64 计算，避免四数之和在 32 位 int 上溢出。
				need := int64(target) - int64(nums[i]) - int64(nums[j]) - int64(nums[k])

				if seen[need] {
					// need was found strictly between j and k, so the four values are already ascending.
					// need 来自 j 和 k 之间的位置，因此这四个数天然升序。
					// It equals an existing element, so converting it back to int cannot overflow.
					// 它等于数组中已有的元素，所以转回 int 不会溢出。
					key := [4]int{nums[i], nums[j], int(need), nums[k]}
					if !recorded[key] {
						recorded[key] = true
						result = append(result, []int{key[0], key[1], key[2], key[3]})
					}
				}

				// Insert after lookup so one position cannot supply both c and d.
				// 查询之后再插入，避免同一位置同时充当 c 和 d。
				seen[int64(nums[k])] = true
			}
		}
	}

	return result
}
