package _1_array

/*
题目描述 / Problem Description
给定正整数 n，生成一个 n x n 矩阵，按照顺时针螺旋顺序填入从 1 到 n² 的所有整数。
Given a positive integer n, generate an n x n matrix filled with the integers from 1 to n² in clockwise spiral order.
*/

// 1. Shrinking boundaries: fill one ring per round and move each boundary inward right away.
// 1. 边界收缩法：每轮填完一圈，并且填完一条边就立刻收缩对应边界。
// Time: O(n²), Space: O(1) auxiliary beyond the O(n²) output matrix.
// 时间复杂度：O(n²)，除 O(n²) 输出矩阵外辅助空间 O(1)。
//
// top/bottom/left/right 围住尚未填充的区域；每填完一条边就收缩对应边界，角点不会重复写入。
// The four bounds enclose unfilled cells; shrink each bound after filling its edge to avoid repeating corners.
// 填下边和左边前重新判定区域是否还存在，防止最后只剩一行或一列时重复填充；n>=0。
// Recheck before the bottom and left edges so a final single row or column is not filled twice; require n>=0.
func generateMatrix(n int) [][]int {
	ans := make([][]int, n)
	for i := range ans {
		ans[i] = make([]int, n)
	}

	top, bottom := 0, n-1
	left, right := 0, n-1
	num := 1

	for top <= bottom && left <= right {
		for i := left; i <= right; i++ {
			ans[top][i] = num
			num++
		}
		top++

		for i := top; i <= bottom; i++ {
			ans[i][right] = num
			num++
		}
		right--

		if top <= bottom {
			for i := right; i >= left; i-- {
				ans[bottom][i] = num
				num++
			}
			bottom--
		}

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
//
// 零表示未填；按右、下、左、上行走，下一格越界或已填时顺时针转向。
// Zero marks an unfilled cell; walk right/down/left/up and turn clockwise at a boundary or filled cell.
// 写完 n*n 后直接停止，避免再移动到无空位的区域；n>=0，n=0 时循环不执行。
// Stop after n*n writes because no next empty cell remains; for n=0 the loop never runs.
func generateMatrixSimulation(n int) [][]int {
	result := make([][]int, n)
	for i := range result {
		result[i] = make([]int, n)
	}

	directions := [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	row, col, direction := 0, 0, 0

	for value := 1; value <= n*n; value++ {
		result[row][col] = value

		if value == n*n {
			break
		}

		nextRow, nextCol := row+directions[direction][0], col+directions[direction][1]

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
//
// 每圈四条边都“含起点、不含终点”，让四个角各归一条边；start/end 是本圈的两端下标。
// Each edge includes its start and excludes its end, assigning each corner once; start/end bound the layer.
// 完整圈只有 n/2 个；奇数阶最后剩中心一格，单独填入。要求 n>=0。
// There are n/2 complete layers; an odd-sized matrix leaves one center cell. Require n>=0.
func generateMatrixByLayers(n int) [][]int {
	result := make([][]int, n)
	for i := range result {
		result[i] = make([]int, n)
	}

	value := 1

	for start := 0; start < n/2; start++ {
		end := n - 1 - start

		for col := start; col < end; col++ {
			result[start][col] = value
			value++
		}

		for row := start; row < end; row++ {
			result[row][end] = value
			value++
		}

		for col := end; col > start; col-- {
			result[end][col] = value
			value++
		}

		for row := end; row > start; row-- {
			result[row][start] = value
			value++
		}
	}

	if n%2 == 1 {
		result[n/2][n/2] = value
	}

	return result
}
