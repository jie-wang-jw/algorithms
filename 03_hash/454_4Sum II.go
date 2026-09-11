package _3_hash

import "sort"

/*
题目描述 / Problem Description
给定四个整数数组 nums1、nums2、nums3 和 nums4，
统计满足 nums1[i] + nums2[j] + nums3[k] + nums4[l] = 0 的下标四元组数量。
Given four integer arrays nums1, nums2, nums3, and nums4,
count index tuples satisfying nums1[i] + nums2[j] + nums3[k] + nums4[l] = 0.

解题思路 / Solution Approach
将四个数组分成两组。用哈希表统计前两个数组所有两数和的出现次数，
再遍历后两个数组的两数和，累加其相反数在哈希表中的频次。
Split the arrays into two pairs. Count every sum from the first pair in a hash map,
then scan sums from the second pair and add the frequency of each opposite sum.

关键逻辑：为什么这样做 / Why This Works
对固定的 (c,d)，每一对满足 a+b=-(c+d) 的下标都能组成一个不同四元组，所以应加上频次，而不是只加 1。
若某个所需和由前两数组产生 3 次，又被后两数组的 2 对下标需要，扫描时就累加 3+3=6；值相同但下标不同也要计数。
For each fixed (c,d) index pair, every (a,b) index pair with sum -(c+d) creates a distinct tuple,
so add its frequency rather than 1. If a needed sum occurs for 3 first-pair combinations and 2
second-pair combinations, the scans add 3+3=6. Equal values at different indices still count separately.

时间与空间复杂度 / Time and Space Complexity
四个数组等长 n 时，平均时间 O(n²)，前两组配对各遍历 n² 次；辅助空间 O(n²)，
保存不同两数和及频次。若长度分别为 a,b,c,d，则时间 O(ab+cd)，空间 O(ab)。
For four arrays of length n, average time is O(n²) and
auxiliary space O(n²) for pair-sum frequencies. With lengths a,b,c,d, time is O(ab+cd) and space O(ab).

补充解法：两组和排序 / Alternative: Sort Pair Sums
fourSumCountSorted 保存 A+B 和 C+D 的全部位置组合（不要去重），排序后从两端找和为 0。
命中时若左和出现 x 次、右和出现 y 次，就贡献 x*y 组下标组合，而不是只加 1。
四个数组各长 n 时，时间 O(n² log(n+1))，辅助空间 O(n²)，不修改输入；原哈希法平均更快。
Store all A+B and C+D index-pair sums, keeping multiplicities, then sort and search from opposite ends.
A matching run of x left sums and y right sums contributes x*y tuples, not one.
For four length-n arrays: O(n² log(n+1)) time, O(n²) auxiliary space, no input mutation.
*/

/*
Split the four arrays into two pairs.
把四个数组分成两组处理。

First, I count all possible sums from nums1 and nums2 using a hash map.
先用哈希表统计 nums1 和 nums2 所有两数之和出现的次数。

Then, for each sum from nums3 and nums4, I look for its negative value in the map.
再遍历 nums3 和 nums4 的两数之和，并在哈希表中查找它的相反数。

The frequency in the map tells me how many valid tuples we can form.
哈希表中的频次表示当前组合能够配出多少个有效四元组。
*/

// 1. 分组哈希：前两数和存 map，推荐
func fourSumCount(A []int, B []int, C []int, D []int) int {
	// Key: a + b; value: how many index pairs produce that sum.
	// key 是 a+b，value 是产生这个和的下标组合数量。
	sumCount := map[int]int{}
	count := 0

	for _, a := range A {
		for _, b := range B {
			// Do not store only true: duplicate pairs must all be counted.
			// 不能只保存 true，因为不同下标产生的重复组合也必须计数。
			sumCount[a+b]++
		}
	}

	for _, c := range C {
		for _, d := range D {
			// We need (a+b) + (c+d) = 0, so a+b must equal -(c+d).
			// 要满足总和为 0，a+b 必须等于 -(c+d)。
			need := -(c + d)
			// A missing key has value 0 in Go, so it adds nothing automatically.
			// Go 中不存在的 map key 读取结果为 0，因此无需单独判断是否存在。
			count += sumCount[need]
		}
	}

	return count
}

// 2. 两组和排序 + 双指针：保留重复次数
func fourSumCountSorted(a, b, c, d []int) int {
	// Keep every index combination; deduplicating sums here would lose tuples.
	// 保留每一个下标组合，这里去重会丢掉答案。
	leftSums, rightSums := []int{}, []int{}

	for _, x := range a {
		for _, y := range b {
			leftSums = append(leftSums, x+y)
		}
	}
	for _, x := range c {
		for _, y := range d {
			rightSums = append(rightSums, x+y)
		}
	}

	// Sorting gives the two pointers a direction to move in.
	// 排序后双指针才有明确的移动方向。
	sort.Ints(leftSums)
	sort.Ints(rightSums)

	result := 0

	// i scans the left sums upward and j scans the right sums downward.
	// i 从小到大扫描左侧的和，j 从大到小扫描右侧的和。
	for i, j := 0, len(rightSums)-1; i < len(leftSums) && j >= 0; {
		sum := leftSums[i] + rightSums[j]

		if sum < 0 {
			// Even the largest remaining right sum is too small for this left sum.
			// 即使配上剩余最大的右侧和仍然偏小，因此放弃当前左侧和。
			i++
		} else if sum > 0 {
			// Even the smallest remaining left sum is too large for this right sum.
			// 即使配上剩余最小的左侧和仍然偏大，因此放弃当前右侧和。
			j--
		} else {
			// Equal sums appear in runs after sorting, so measure both runs at once.
			// 排序后相等的和是连续的一段，所以一次量出两侧的整段长度。
			x, y := leftSums[i], rightSums[j]
			startI, startJ := i, j

			for i < len(leftSums) && leftSums[i] == x {
				i++
			}
			for j >= 0 && rightSums[j] == y {
				j--
			}

			// Multiplicities count independent choices of the two index pairs.
			// 两组下标对可以独立选择，所以重复次数相乘。
			result += (i - startI) * (startJ - j)
		}
	}

	return result
}
