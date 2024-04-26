// Created by Daniel J. Rollins at 2024/04/16 20:12
// leetgo: 1.4.5
// https://leetcode.com/problems/valid-sudoku/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func isValidSudoku(board [][]byte) bool {

	for i := 0; i < 9; i++ {
		cell := board[i][0]
		if rows[cell] {
			return false
		}
		rows[cell] = true

		for j := 0; j < 9; j++ {
			cell := board[i][j]
			if cols[cell] {
				return false
			}
			cols[cell] = true

			sub := i % 3 + j % 3
			if subs[sub] 
		}
	}
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	board := Deserialize[[][]byte](ReadLine(stdin))
	ans := isValidSudoku(board)

	fmt.Println("\noutput:", Serialize(ans))
}
