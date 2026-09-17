package _1_array

/*
卡码 44. 开发商购买土地 / Developer Purchasing Land

题目描述 / Problem Description
n×m 区块各有土地价值。只允许沿一条横向或纵向直线把区域分成两块，每块至少一个区块。
求两块土地总价值之差的最小值。
An n by m grid of land values must be split by one horizontal or one vertical cut
into two nonempty blocks. Return the minimum absolute difference of the two block sums.

示例 / Example
3×3 网格 [[1,2,3],[2,1,3],[1,2,3]] 沿第二、三列之间切开，两侧和都是 9，答案 0。
The 3 by 3 grid [[1,2,3],[2,1,3],[1,2,3]] can be split between columns 2 and 3 into two sums of 9, so the answer is 0.
*/

import (
	"bufio"
	"fmt"
	"os"
)

// 1. Row and column prefix sums
// 1. 行列前缀和
// Accumulate full rows/columns as one block; difference is |sum-2*cut|. Skip cuts that leave a side empty.
// 累加完整行/列作为一块，差值 |sum-2*cut|。跳过会让一侧为空的切割。
// Time: O(n*m), Space: O(n+m) for the row and column sums.
// 时间复杂度：O(n*m)，空间复杂度：O(n+m)，由行列和数组产生。
//
// 按题意输入矩形正数网格，且至少有一条合法分割线。先统计总和 sum、每行和 horizontal、每列和 vertical。
// For a positive rectangular grid with a legal cut, compute total, row sums, and column sums.
// cut 为分割线一侧的累计和，另一侧为 sum-cut，差值因此是 |sum-2*cut|。
// If one side sums to cut, the other is sum-cut, giving difference |sum-2*cut|.
// 只枚举前 n-1 行和前 m-1 列之后的切线，保证两侧非空；所有合法横切、竖切都被覆盖。
// Cut only after the first n-1 rows or m-1 columns, covering all cuts with two nonempty sides.
func landPurchasePrefix(grid [][]int) int {
	n := len(grid)
	if n == 0 {
		return 0
	}
	m := len(grid[0])

	sum := 0
	horizontal := make([]int, n)
	vertical := make([]int, m)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			sum += grid[i][j]
			horizontal[i] += grid[i][j]
			vertical[j] += grid[i][j]
		}
	}

	result := absInt(sum)
	horizontalCut := 0
	for i := 0; i < n-1; i++ {
		horizontalCut += horizontal[i]
		result = min(result, absInt(sum-2*horizontalCut))
	}

	verticalCut := 0
	for j := 0; j < m-1; j++ {
		verticalCut += vertical[j]
		result = min(result, absInt(sum-2*verticalCut))
	}
	return result
}

// 2. Accumulate while scanning
// 2. 遍历时累加
// Update at each row end and each column end without storing separate row/column sums.
// 扫到行尾或列尾时更新答案，不再单独保存行列和。
// Time: O(n*m), Space: O(1) besides the input grid.
// 时间复杂度：O(n*m)，空间复杂度：除输入外 O(1)。
//
// count 累加完整行或完整列；只有到行尾或列尾才对应一条直线分割，差值为 |sum-2*count|。
// count accumulates whole rows or columns; only their ends define straight cuts, with difference |sum-2*count|.
// 按列扫描前清零 count；最后一行或列之后不能切，否则另一块为空。沿用正数矩形及可分割的题目约束。
// Reset count before scanning columns; exclude cuts after the final row or column. Assume a positive, splittable rectangle.
func landPurchaseScan(grid [][]int) int {
	n := len(grid)
	if n == 0 {
		return 0
	}
	m := len(grid[0])

	sum := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			sum += grid[i][j]
		}
	}

	result := absInt(sum)
	count := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			count += grid[i][j]
			if j == m-1 && i < n-1 {
				result = min(result, absInt(sum-2*count))
			}
		}
	}

	count = 0
	for j := 0; j < m; j++ {
		for i := 0; i < n; i++ {
			count += grid[i][j]
			if i == n-1 && j < m-1 {
				result = min(result, absInt(sum-2*count))
			}
		}
	}
	return result
}

// ACM input: read n, m and the grid, then print the prefix-sum answer.
// ACM 输入：读入 n、m 和网格，再打印前缀和解。
// Time: O(n*m), Space: O(n*m) for the grid plus O(n+m) for the prefix arrays.
// 时间复杂度：O(n*m)，空间复杂度：网格 O(n*m)，另加行列和 O(n+m)。
//
// 按 ACM 格式读取 n、m 和 n*m 个值；输入不足即返回，完整读入后输出最小分割差值。
// Read n, m and n*m values; stop on incomplete input, otherwise print the minimum partition difference.
func landPurchase() {
	in := bufio.NewReader(os.Stdin)
	var n, m int
	if _, err := fmt.Fscan(in, &n, &m); err != nil {
		return
	}

	grid := make([][]int, n)
	for i := 0; i < n; i++ {
		grid[i] = make([]int, m)
		for j := 0; j < m; j++ {
			if _, err := fmt.Fscan(in, &grid[i][j]); err != nil {
				return
			}
		}
	}
	fmt.Println(landPurchasePrefix(grid))
}

// Absolute value of an integer.
// 整数绝对值。
// Time: O(1), Space: O(1).
// 时间复杂度：O(1)，空间复杂度：O(1)。
//
// 返回整数绝对值；输入范围须保证负数取反不溢出。
// Return the absolute value; the input range must allow negation without overflow.
func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
