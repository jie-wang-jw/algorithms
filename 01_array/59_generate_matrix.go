package _1_array

/*
题目描述 / Problem Description
给定正整数 n，生成一个 n x n 矩阵，按照顺时针螺旋顺序填入从 1 到 n² 的所有整数。
Given a positive integer n, generate an n x n matrix filled with the integers from 1 to n² in clockwise spiral order.

解题思路 / Solution Approach
维护上、下、左、右四条未填充边界。每轮依次填充上边、右边、下边和左边，然后将四条边界向内收缩。
Maintain top, bottom, left, and right boundaries for the unfilled area. Fill the four sides clockwise, then move all boundaries inward.

关键逻辑：为什么这样做 / Why This Works
四条边围住尚未填写的区域。每填完一条边就立即收缩该边界，而不是等四条边都写完才收缩：例如 top++ 后，右边从下一行开始，右上角不会重复填写；
其他角同理。后两条边先检查剩余行/列，避免处理已经消失的区域。n=3 填完外圈后四个边界都为 1，下一轮只填中心的 9，随后边界交错，循环结束。
The four boundaries enclose unfilled cells. Shrink each boundary immediately after filling that side, not after all four sides:
top++ makes the right column start one row lower, avoiding a second write to the top-right corner. Other corners follow the same rule.
Check remaining rows/columns before the last two sides. For n=3, all boundaries become 1 after the outer ring;
the next round fills only the center with 9, then the boundaries cross.

时间与空间复杂度 / Time and Space Complexity
n 为矩阵边长。时间 O(n²)，每个格子填入一次。除返回矩阵外辅助空间 O(1)；返回矩阵占 O(n²)，包含结果的总空间为 O(n²)。
n is the matrix side length. Time O(n²), filling each cell once. Auxiliary space excluding the output is O(1);
the returned matrix and total space take O(n²).

补充解法 / Additional Approaches
generateMatrixSimulation：按右、下、左、上移动；下一格越界或非零时顺时针转向。
已填数字从 1 开始，因此 0 可当未访问标记，无需额外 visited 数组。
Move right, down, left, up; turn when the next cell is outside or already filled.
Values start at 1, so zero identifies an unvisited cell without a separate visited array.
generateMatrixByLayers：逐圈填四条边，每条边包含起点、不包含终点，让下一条边负责角点。
奇数阶最后剩中心格，单独填入；避免同一个角写两次。
Fill each ring with start-inclusive, end-exclusive edges: each corner belongs to exactly one edge.
For odd n, fill the remaining center separately.
两者时间 O(n²)，除 O(n²) 输出矩阵外辅助空间 O(1)。边界收缩版仍适合优先掌握。
Both take O(n²) time and O(1) auxiliary space beyond the O(n²) output.
*/

// 1. Shrinking boundaries: fill one ring per round and move each boundary inward right away.
// 1. 边界收缩法：每轮填完一圈，并且填完一条边就立刻收缩对应边界。
// Time: O(n²), Space: O(1) auxiliary beyond the O(n²) output matrix.
// 时间复杂度：O(n²)，除 O(n²) 输出矩阵外辅助空间 O(1)。
func generateMatrix(n int) [][]int {
	// Create an n by n matrix initialized with zeros.
	// 创建一个 n x n 的二维数组，初始值都为 0。
	ans := make([][]int, n)
	for i := range ans {
		ans[i] = make([]int, n)
	}

	// These four boundaries describe the unfilled rectangle.
	// 这四个边界围住当前尚未填充的矩形区域。
	top, bottom := 0, n-1
	left, right := 0, n-1
	// num is the next value to write into the matrix.
	// num 是下一个要写入矩阵的数字。
	num := 1

	// Fill one outer layer clockwise, then move all boundaries inward.
	// 每轮顺时针填完一层，再把四条边界向内收缩。
	for top <= bottom && left <= right {
		// 1. Fill the top row from left to right. / 从左到右填充 top 行。
		for i := left; i <= right; i++ {
			ans[top][i] = num
			num++
		}
		top++

		// 2. Fill the right column from top to bottom. / 从上到下填充 right 列。
		for i := top; i <= bottom; i++ {
			ans[i][right] = num
			num++
		}
		right--

		// 3. Fill the bottom row from right to left. / 从右到左填充 bottom 行。
		// Check first because the top row may have used the last remaining row.
		// 先判断是否还有行，因为 top 行可能已经填掉最后一行。
		if top <= bottom {
			for i := right; i >= left; i-- {
				ans[bottom][i] = num
				num++
			}
			bottom--
		}

		// 4. Fill the left column from bottom to top. / 从下到上填充 left 列。
		// Check first because the right column may have used the last remaining column.
		// 先判断是否还有列，因为 right 列可能已经填掉最后一列。
		if left <= right {
			for i := bottom; i >= top; i-- {
				ans[i][left] = num
				num++
			}
			left++
		}
	}

	return ans
}

// 2. Direction simulation: walk right, down, left, up and turn clockwise when blocked.
// 2. 方向模拟法：按右、下、左、上行走，遇到阻挡就顺时针转向。
// Time: O(n²), Space: O(1) auxiliary beyond the O(n²) output matrix.
// 时间复杂度：O(n²)，除 O(n²) 输出矩阵外辅助空间 O(1)。
func generateMatrixSimulation(n int) [][]int {
	result := make([][]int, n)
	for i := range result {
		result[i] = make([]int, n)
	}

	// The four offsets are ordered clockwise: right, down, left, up.
	// 四个方向偏移按顺时针排列：右、下、左、上。
	directions := [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	row, col, direction := 0, 0, 0

	for value := 1; value <= n*n; value++ {
		result[row][col] = value

		// The last value needs no move, and moving would also find no free cell.
		// 最后一个数字不需要再移动，此时也已经没有空位可走。
		if value == n*n {
			break
		}

		nextRow, nextCol := row+directions[direction][0], col+directions[direction][1]

		// A filled cell is a wall, just like the matrix boundary.
		// 已填写的格子和矩阵边界一样，都表示应该转弯。
		// Values start at 1, so a zero reliably marks an unvisited cell.
		// 填入的数字从 1 开始，所以 0 可以可靠地表示未访问过的格子。
		if nextRow < 0 || nextRow >= n || nextCol < 0 || nextCol >= n || result[nextRow][nextCol] != 0 {
			direction = (direction + 1) % 4
		}

		row += directions[direction][0]
		col += directions[direction][1]
	}

	return result
}

// 3. Ring by ring: fill four half-open edges per ring, then the center when n is odd.
// 3. 逐圈填充法：每圈填四条左闭右开的边，n 为奇数时最后单独填中心格。
// Time: O(n²), Space: O(1) auxiliary beyond the O(n²) output matrix.
// 时间复杂度：O(n²)，除 O(n²) 输出矩阵外辅助空间 O(1)。
func generateMatrixByLayers(n int) [][]int {
	result := make([][]int, n)
	for i := range result {
		result[i] = make([]int, n)
	}

	value := 1

	// n/2 rings have four full edges; an odd n leaves a single center cell.
	// 共有 n/2 圈拥有完整的四条边；n 为奇数时还会剩下一个中心格。
	for start := 0; start < n/2; start++ {
		end := n - 1 - start

		// Each edge excludes its ending corner, so every corner belongs to exactly one edge.
		// 每条边不含终点角，因此每个角点只被恰好一条边填写。
		// Top edge: left to right. / 上边：从左到右。
		for col := start; col < end; col++ {
			result[start][col] = value
			value++
		}

		// Right edge: top to bottom. / 右边：从上到下。
		for row := start; row < end; row++ {
			result[row][end] = value
			value++
		}

		// Bottom edge: right to left. / 下边：从右到左。
		for col := end; col > start; col-- {
			result[end][col] = value
			value++
		}

		// Left edge: bottom to top. / 左边：从下到上。
		for row := end; row > start; row-- {
			result[row][start] = value
			value++
		}
	}

	// An odd side length leaves one uncovered center cell.
	// 奇数阶矩阵的正中心格不属于任何一圈的四条边，需要单独填写。
	if n%2 == 1 {
		result[n/2][n/2] = value
	}

	return result
}
