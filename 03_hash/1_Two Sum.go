package _3_hash

/*
题目描述 / Problem Description
给定整数数组 nums 和整数 target，找出两个和为 target 的不同元素下标并返回。题目保证恰好存在一个答案。
Given an integer array nums and an integer target, return the indices of two distinct elements whose sum equals target. Exactly one answer is guaranteed.

解题思路 / Solution Approach
遍历数组并用哈希表保存已经见过的数字及其下标。对当前数字 num，检查 target-num 是否已经出现；如果出现即可返回两个下标。
Scan the array while storing previously seen values and indices in a hash map. For each num, check whether target-num has already appeared and return the two indices when found.

时间与空间复杂度 / Time and Space Complexity
n = len(nums)。在哈希表查询和写入平均 O(1) 的假设下，总时间 O(n)。辅助空间 O(n)，最坏保存几乎全部数字及下标；返回两个下标只占 O(1)。
n = len(nums). Assuming average O(1) hash lookup and insertion, time is O(n). Auxiliary space O(n) stores seen values and indices; the two-index output takes O(1).
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
