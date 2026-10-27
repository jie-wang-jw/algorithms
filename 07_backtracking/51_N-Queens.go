package _7_backtracking

/*
51. N 皇后 / N-Queens

题目描述 / Problem Description
在 n×n 棋盘上放置 n 个皇后，使它们不能互相攻击（不同行、列、对角线），返回所有棋盘布局。
Queen 用 'Q' 表示，空位用 '.'。
Place n queens on an n×n board so that none share a row, column, or diagonal, and return every board.
Queens are 'Q'; empty squares are '.'.

示例 / Example
n=4 有两解。其一（1-indexed 行列）：(1,2),(2,4),(3,1),(4,3)；另一组对称。
n=4 has two solutions. One (1-based): (1,2),(2,4),(3,1),(4,3); the other is symmetric.
*/

// 面试首选 / Interview first choice：51 N 皇后。
//
// 1. Backtracking: place one queen per row, skipping columns and diagonals already attacked.
// 1. 回溯：每一行放一个皇后，跳过已占用的列和两条对角线。
// Time: O(n*n! + Q*n^2), Q solutions; auxiliary space O(n^2), output O(Q*n^2).
// Q 为解数；搜索及复制结果时间 O(n*n!+Q*n²)，辅助空间 O(n²) 含棋盘，输出 O(Q*n²)。
//
// 每行放一枚，行冲突自然排除；colUsed 管列，row-col+n 管主对角线，row+col 管副对角线。
// Place one queen per row; track columns, main diagonals by row-col+n, and anti-diagonals by row+col.
// 同一对角线上差值或和恒定；选择时同步标记，退出时全部恢复，满 n 行才复制棋盘。
// Differences or sums are constant along a diagonal; mark on choice, restore on return, and copy the board after n rows.
//
// 边界：按题意 n>=1；n=1 有一个解，n=2、3 无解；只有放满 n 行才保存结果。
func solveNQueens(n int) [][]string {
	result := [][]string{}
	board := make([][]byte, n)
	for i := 0; i < n; i++ {
		row := make([]byte, n)
		for j := 0; j < n; j++ {
			row[j] = '.'
		}
		board[i] = row
	}
	colUsed := make([]bool, n)
	diagUsed := make([]bool, 2*n)
	antiDiagUsed := make([]bool, 2*n)

	var backtrack func(int)
	backtrack = func(row int) {
		if row == n {
			solution := make([]string, n)
			for i := 0; i < n; i++ {
				solution[i] = string(board[i])
			}
			result = append(result, solution)
			return
		}
		for col := 0; col < n; col++ {
			diagIndex := row - col + n
			antiDiagIndex := row + col
			if colUsed[col] || diagUsed[diagIndex] || antiDiagUsed[antiDiagIndex] {
				continue
			}
			board[row][col] = 'Q'
			colUsed[col], diagUsed[diagIndex], antiDiagUsed[antiDiagIndex] = true, true, true
			backtrack(row + 1)
			board[row][col] = '.'
			colUsed[col], diagUsed[diagIndex], antiDiagUsed[antiDiagIndex] = false, false, false
		}
	}

	backtrack(0)
	return result
}
