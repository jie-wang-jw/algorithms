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

解法一：行列前缀和 / Method 1: Row and Column Prefix Sums
先求总和、每一行和、每一列和。再依次累加前若干行（或前若干列）作为一块，另一块是总和减去这一块。
差值为 |sum - 2*cut|。每个子区域至少一个区块，因此不在切完整张图之后更新答案。
Compute the total sum and each row/column sum. Accumulate the first i rows (or first j columns) as one block;
the other block is total minus that prefix. The difference is |sum - 2*cut|.
Skip a cut that would leave one side empty.

解法二：遍历时累加（优化暴力） / Method 2: Accumulate While Scanning
不必单独存行列和。按行扫描，每到行尾用当前累计和更新答案；再按列扫描，每到列尾更新一次。
There is no need to store row and column sums separately. Scan row-major and update at each row end;
then scan column-major and update at each column end.

时间与空间复杂度 / Time and Space Complexity
n 为行数，m 为列数。两种解法都扫描整个矩阵常数次，时间 O(n*m)。
解法一额外 O(n+m) 存行列和；解法二除输入外 O(1)。返回一个整数。
For n rows and m columns both methods scan the matrix a constant number of times, so time is O(n*m).
Method 1 uses O(n+m) extra space for row and column sums; method 2 uses O(1) besides the input.
Both return a single integer.
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
// 步骤与要点 / Steps and notes:
//  1. Stop before the last row so the bottom block stays nonempty.
//     不切在最后一行之后，保证下半块至少有一行。
//  2. |sum-2*cut| == |top-bottom| because bottom = sum - cut.
//     |sum-2*cut| 等于 |上半-下半|，因为下半 = sum - cut。
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
// 步骤与要点 / Steps and notes:
//  1. Valid horizontal cut only when some rows remain below.
//     只有下面还剩行时，这一行的结束才是合法横向切割。
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
func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
