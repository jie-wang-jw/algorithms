package _7_backtracking

/*
37. 解数独 / Sudoku Solver

题目描述 / Problem Description
在 9×9 棋盘上填充 '.'，使每行、每列、每个 3×3 宫内 1-9 各出现一次。题目保证有唯一解。
Fill '.' cells on a 9 by 9 board so each row, column, and 3 by 3 box contains 1-9 once.
The puzzle has a unique solution.
*/

// 面试首选 / Interview first choice：37 解数独。
//
// 1. Two-dimensional backtracking: try digits 1-9 on each empty cell and return as soon as one filling works.
// 1. 二维回溯：在每个空格尝试 1-9，找到一种填法就立刻返回。
// Time: O(9^e) for e empty cells, Space: O(e) for the call stack.
// 时间复杂度：O(9^e)（e 为空格数），空间复杂度：O(e)，由递归栈产生。
//
// 找到第一个空格后逐个试填；一个候选引出完整解就保留棋盘并返回 true，否则恢复 '.' 再试下一个。
// Try values at the first empty cell; retain a complete solution or restore '.' before the next candidate.
// 该空格所有候选都失败就返回 false，不能跳过它去填别处；依赖合法的 9×9 初始棋盘，外层原地求解。
// If every candidate fails, return false rather than skip that cell; assume a valid 9×9 puzzle and solve in place.
func solveSudoku(board [][]byte) {
	solveSudokuBacktrack(board)
}

// Fill the next empty cell; return true when the remaining board has a valid filling.
// 填充下一个空格；剩下的棋盘有解时返回 true。
// Time: O(9^e), Space: O(e).
// 时间复杂度：O(9^e)，空间复杂度：O(e)。
//
// 找到第一个空格后逐个试填；一个候选引出完整解就保留棋盘并返回 true，否则恢复 '.' 再试下一个。
// Try values at the first empty cell; retain a complete solution or restore '.' before the next candidate.
// 该空格所有候选都失败就返回 false，不能跳过它去填别处；依赖合法的 9×9 初始棋盘，外层原地求解。
// If every candidate fails, return false rather than skip that cell; assume a valid 9×9 puzzle and solve in place.
func solveSudokuBacktrack(board [][]byte) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] != '.' {
				continue
			}
			for k := byte('1'); k <= '9'; k++ {
				if sudokuIsValid(i, j, k, board) {
					board[i][j] = k
					if solveSudokuBacktrack(board) {
						return true
					}
					board[i][j] = '.'
				}
			}
			return false
		}
	}
	return true
}

// A digit is valid when it is missing from the row, column, and 3 by 3 box.
// 数字在同行、同列、所在宫内都未出现时合法。
// Time: O(1), Space: O(1).
// 时间复杂度：O(1)，空间复杂度：O(1)。
//
// 只需检查同一行、同一列和同一宫；宫左上角为 (row/3*3,col/3*3)，除法先得到宫编号再换回坐标。
// Check the row, column and box; (row/3*3,col/3*3) converts the box indices back to its top-left coordinates.
func sudokuIsValid(row, col int, val byte, board [][]byte) bool {
	for i := 0; i < 9; i++ {
		if board[row][i] == val {
			return false
		}
	}
	for j := 0; j < 9; j++ {
		if board[j][col] == val {
			return false
		}
	}
	startRow := (row / 3) * 3
	startCol := (col / 3) * 3
	for i := startRow; i < startRow+3; i++ {
		for j := startCol; j < startCol+3; j++ {
			if board[i][j] == val {
				return false
			}
		}
	}
	return true
}
