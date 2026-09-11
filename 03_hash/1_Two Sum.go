package _3_hash

import "sort"

/*
题目描述 / Problem Description
给定整数数组 nums 和整数 target，找出两个和为 target 的不同元素下标并返回。
题目保证恰好存在一个答案。
Given an integer array nums and an integer target, return the indices of two distinct elements whose sum equals target.
Exactly one answer is guaranteed.

解题思路 / Solution Approach
遍历数组并用哈希表保存已经见过的数字及其下标。对当前数字 num，
检查 target-num 是否已经出现；如果出现即可返回两个下标。
Scan the array while storing previously seen values and indices in a hash map. For each num,
check whether target-num has already appeared and return the two indices when found.

关键逻辑：为什么这样做 / Why This Works
查找时 seen 只保存下标小于 i 的元素，所以命中 j 后必有 j!=i，且 nums[j]=target-num，
两个数之和必为 target。例如 [3,3]、target=6：第一个 3 先存入，第二个 3 才找到它；不会把一个 3 使用两次。
At lookup time, seen contains only indices smaller than i.
A match j is therefore distinct from i and satisfies nums[j]=target-num.
With [3,3] and target=6, store the first 3, then match it when reading the second; no element is reused.

时间与空间复杂度 / Time and Space Complexity
n = len(nums)。在哈希表查询和写入平均 O(1) 的假设下，总时间 O(n)。辅助空间 O(n)，
最坏保存几乎全部数字及下标；返回两个下标只占 O(1)。
n = len(nums). Assuming average O(1) hash lookup and insertion, time is O(n).
Auxiliary space O(n) stores seen values and indices; the two-index output takes O(1).

补充解法 / Additional Approaches
twoSumBruteForce 枚举 i<j，保证不重复使用同一位置；时间 O(n²)，辅助空间 O(1)。
Enumerate i<j to use distinct positions: O(n²) time and O(1) auxiliary space.
twoSumSorted 将值与原下标一起复制后排序，再用双指针。
和偏小就增大左值，偏大就减小右值；返回保存的原下标，不能返回排序后的下标。
两个下标按数值排序的位置返回，因此不保证升序，题目也不要求顺序。
时间 O(n log n)，辅助空间 O(n)，不修改输入；哈希版平均 O(n) 更快。
Sort copied value-index pairs and use two pointers, increasing the left value for a small sum
and decreasing the right value for a large sum. Return original indices.
They come out in sorted-value order, so the pair is not guaranteed to be ascending; the problem allows any order.
Time O(n log n), auxiliary space O(n), no input mutation; hashing is expected O(n).
*/

/*
Use a hash map to store numbers I have already seen and their indices.
使用哈希表保存已经见过的数字及其下标。

For each number, I compute its complement, which is target minus the current number.
对每个数字计算它需要的另一个数：target - 当前数字。

If the complement is already in the map, I return the two indices.
如果需要的数已经在哈希表中，就返回两个下标。

Otherwise, I store the current number and continue.
否则保存当前数字并继续扫描。
*/

// 1. 哈希表一次遍历：推荐
func twoSum(nums []int, target int) []int {
	// Key: a seen number; value: its index.
	// key 是已经见过的数字，value 是该数字的下标。
	seen := map[int]int{}

	for i, num := range nums {
		// num + need = target.
		// 当前数字 num 需要搭配 need 才能得到 target。
		need := target - num
		// ok tells us whether need already exists in the map.
		// ok 表示 need 是否已经存在于哈希表中。
		if j, ok := seen[need]; ok {
			return []int{j, i}
		}
		// Store after checking so the same element cannot be used twice.
		// 先查找、后存入，避免同一个元素被使用两次。
		seen[num] = i
	}

	// No valid pair was found.
	// 没有找到满足条件的两个数。
	return nil
}

// 2. 暴力双循环：枚举所有下标对
func twoSumBruteForce(nums []int, target int) []int {
	// i selects the first position of the pair.
	// i 选择数对中的第一个位置。
	for i := 0; i < len(nums); i++ {
		// Starting j at i+1 keeps the two positions distinct and avoids checking a pair twice.
		// j 从 i+1 开始，既保证两个下标不同，也避免同一对被检查两次。
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	// Every pair of distinct positions has been tried.
	// 所有不同下标的组合都尝试过了。
	return nil
}

// 3. 排序 + 双指针：必须先保存原下标
func twoSumSorted(nums []int, target int) []int {
	// Sorting nums itself would destroy the indices the problem asks for, so copy value and index together.
	// 直接排序 nums 会丢失题目要求返回的下标，所以把数值和原下标一起复制出来。
	pairs := make([][2]int, len(nums))
	for i, v := range nums {
		pairs[i] = [2]int{v, i}
	}

	// Order by value only; each original index travels with its value.
	// 只按数值排序，每个原下标跟着自己的数值一起移动。
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][0] < pairs[j][0]
	})

	// left points at the smallest remaining value and right at the largest.
	// left 指向剩余区间的最小值，right 指向最大值。
	left, right := 0, len(pairs)-1

	for left < right {
		sum := pairs[left][0] + pairs[right][0]

		if sum == target {
			// Return the stored original indices, never the positions after sorting.
			// 返回保存下来的原下标，不能返回排序后的位置。
			return []int{pairs[left][1], pairs[right][1]}
		} else if sum < target {
			// The largest partner is already in use, so this left value can be discarded.
			// 当前左值已经配上了最大的右值仍然偏小，因此可以放弃这个左值。
			left++
		} else {
			// The smallest partner is already in use, so this right value can be discarded.
			// 当前右值已经配上了最小的左值仍然偏大，因此可以放弃这个右值。
			right--
		}
	}

	// The two pointers met without finding a pair.
	// 双指针相遇，仍然没有找到答案。
	return nil
}
