package _3_hash

/*
题目描述 / Problem Description
给定整数数组 nums 和整数 target，找出所有和为 target 的不重复四元组。
每个四元组必须使用四个不同下标的元素。
Given an integer array nums and an integer target,
return all unique quadruplets whose sum equals target, using four distinct indices.
*/

import "sort"

// 1. Sort + two nested loops + two pointers + pruning: fix two values, then search the rest from both ends. Recommended.
// 1. 排序 + 双层循环 + 双指针 + 剪枝：固定前两个数，再在剩余区间用左右指针寻找另外两个数。推荐。
// Time: O(n³), Space: O(log n+r) including the sort stack and r quadruplets.
// 时间复杂度：O(n³)，空间复杂度：O(log n+r)，含排序栈与结果。
//
// 排序后固定 i、j，再用双指针找剩余两数；和偏小排除 left，偏大排除 right，理由同有序两数和。
// Sort, fix i and j, then solve sorted two-sum: a small sum excludes left, a large sum excludes right.
// 仅在前缀和非负且超过 target 时剪枝，后续更大的数无法把和降回目标；负目标下不能只判断“超过 target”。
// Prune only when the fixed sum is nonnegative and exceeds target; larger remaining values cannot reduce it. Exceeding a negative target alone is insufficient.
// i 去重覆盖全局，j 只在同一个 i 内去重（j>i+1），否则会误删 [2,2,2,2] 这类答案。
// Deduplicate i globally and j only within its i group (j>i+1), preserving valid repeated-value answers such as [2,2,2,2].
// 命中后两端跳过重复值；i<j<left<right 保证不同下标。排序修改输入，四数和须在 int 范围内。
// Skip repeated endpoints after a match; ordered indices are distinct. Sorting mutates input and four-value sums must fit in int.
func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)

	res := [][]int{}

	for i := 0; i < len(nums)-3; i++ {
		if nums[i] > target && nums[i] >= 0 {
			break
		}

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j := i + 1; j < len(nums)-2; j++ {
			if nums[i]+nums[j] > target && nums[i]+nums[j] >= 0 {
				break
			}

			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			left, right := j+1, len(nums)-1

			for left < right {
				sum := nums[i] + nums[j] + nums[left] + nums[right]

				if sum == target {
					res = append(res, []int{nums[i], nums[j], nums[left], nums[right]})

					b, c := nums[left], nums[right]

					for left < right && nums[left] == b {
						left++
					}

					for left < right && nums[right] == c {
						right--
					}
				} else if sum < target {
					left++
				} else {
					right--
				}
			}
		}
	}

	return res
}

// 2. Fix two values + hash: reuse Two Sum by looking up the remaining pair before inserting.
// 2. 固定两个数 + 哈希：复用两数之和，先查剩余两数再插入。
// Time: O(n³) expected, Space: O(n+r) auxiliary plus O(r) output.
// 时间复杂度：期望 O(n³)，空间复杂度：辅助 O(n+r)，结果另占 O(r)。
//
// 固定 i、j 后，seen 仅保存 (j,k) 内的值；先查 need 再存 nums[k]，保证四个下标各不相同。
// After fixing i and j, seen holds values strictly between j and k; query before insertion to keep all four indices distinct.
// 排序使四元组按值递增，可用 [4]int 键去重；中间加减用 int64，避免题目范围在 32 位 int 上溢出。
// Sorting makes tuples canonical for [4]int deduplication; int64 intermediates protect the problem's bounds on 32-bit systems.
func fourSumHash(nums []int, target int) [][]int {
	sort.Ints(nums)

	result := [][]int{}

	recorded := make(map[[4]int]bool)

	for i := 0; i < len(nums)-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j := i + 1; j < len(nums)-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			seen := make(map[int64]bool)

			for k := j + 1; k < len(nums); k++ {
				need := int64(target) - int64(nums[i]) - int64(nums[j]) - int64(nums[k])

				if seen[need] {
					key := [4]int{nums[i], nums[j], int(need), nums[k]}
					if !recorded[key] {
						recorded[key] = true
						result = append(result, []int{key[0], key[1], key[2], key[3]})
					}
				}

				seen[int64(nums[k])] = true
			}
		}
	}

	return result
}
