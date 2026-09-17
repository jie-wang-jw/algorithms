package _1_array

/*
题目描述 / Problem Description
给定一个整数数组和多组闭区间查询 [left, right]，对每组查询输出该区间内所有元素的总和。
Given an integer array and multiple inclusive range queries [left, right], output the sum of all elements in each queried range.
*/

import (
	"bufio"
	"fmt"
	"os"
)

// 1. Prefix sums: build prefix once, then answer each query with a single subtraction.
// 1. 前缀和：预处理出 prefix 数组，之后每次查询只做一次减法。
// Time: O(n+q) for n values and q queries, Space: O(n) for the nums and prefix arrays.
// 时间复杂度：O(n+q)（n 为元素数、q 为查询数），空间复杂度：O(n)，由 nums 与 prefix 数组产生。
//
// prefix[k] 是前 k 项之和，prefix[0]=0；闭区间 [left,right] 的和为 prefix[right+1]-prefix[left]。
// prefix[k] sums the first k values; inclusive [left,right] sums to prefix[right+1]-prefix[left].
// 逐对读取查询直到读入失败；查询须满足 0<=left<=right<n，累加和须在 int 范围内。
// Read query pairs until input ends or fails; require valid inclusive indices and sums representable as int.
func prefixSum() {
	in := bufio.NewReader(os.Stdin)

	var n int
	if _, err := fmt.Fscan(in, &n); err != nil {
		return
	}

	nums := make([]int, n)
	prefix := make([]int, n+1)

	for i := 0; i < n; i++ {
		if _, err := fmt.Fscan(in, &nums[i]); err != nil {
			return
		}
		prefix[i+1] = prefix[i] + nums[i]
	}

	var left, right int
	for {
		_, err := fmt.Fscan(in, &left, &right)
		if err != nil {
			break
		}

		sum := prefix[right+1] - prefix[left]
		fmt.Println(sum)
	}
}

// 2. Direct summation: add up every value in [left, right] without any preprocessing.
// 2. 逐次求和：不做任何预处理，直接累加闭区间 [left, right] 内的每个值。
// Time: O(right-left+1) per query, Space: O(1).
// 时间复杂度：每次查询 O(right-left+1)，空间复杂度：O(1)。
//
// 逐项累加闭区间 [left,right]，因此包括 right；调用方保证 0<=left<=right<len(nums)。
// Sum both endpoints of [left,right]; the caller supplies valid inclusive indices.
func rangeSumBruteForce(nums []int, left, right int) int {
	sum := 0

	for i := left; i <= right; i++ {
		sum += nums[i]
	}

	return sum
}
