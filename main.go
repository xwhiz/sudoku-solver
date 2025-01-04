package main

import (
	"fmt"
	"log"
	"slices"
)

type Value struct {
	row           int
	col           int
	possibilities []int
}

func main() {
	board := [][]int{
		{7, 0, 3, 0, 0, 9, 1, 0, 5},
		{0, 1, 5, 0, 0, 8, 0, 6, 0},
		{9, 0, 0, 5, 6, 0, 0, 7, 4},
		{0, 0, 0, 8, 0, 2, 4, 0, 0},
		{3, 9, 0, 0, 0, 0, 0, 0, 0},
		{0, 8, 0, 1, 9, 6, 0, 5, 3},
		{5, 4, 9, 0, 1, 3, 0, 0, 7},
		{0, 0, 0, 9, 0, 4, 0, 3, 0},
		{0, 0, 0, 9, 0, 4, 0, 3, 0},
	}

	const logBoard = false
	const maxIters int = 100
	i := 0

	for !hasWon(board) && i < maxIters {
		possibilities := findNewPossibilities(board)
		slices.SortFunc(possibilities, func(a Value, b Value) int {
			return len(a.possibilities) - len(b.possibilities)
		})

		if len(possibilities) == 0 {
			break
		}

		top := possibilities[0]

		if len(top.possibilities) != 1 {
			fmt.Println("Logging the possibilities array")
			for _, p := range possibilities {
				fmt.Println(p)
			}
			log.Fatalf("Unable to find any cell with possibility of only 1 number.")
		}

		fmt.Printf("Change %d %d to %d\n", top.row+1, top.col+1, top.possibilities[0])
		board[top.row][top.col] = top.possibilities[0]
		i += 1

		if logBoard {
			for _, row := range board {
				fmt.Println(row)
			}
		}
	}

	if hasWon(board) {
		fmt.Println("Completed the board.")
		for _, row := range board {
			fmt.Println(row)
		}
	}

}

func findNewPossibilities(board [][]int) []Value {
	possibilities := []Value{}
	for i, row := range board {
		for j, v := range row {
			if v != 0 {
				continue
			}

			possibleNums := []int{}
			for num := 1; num <= 9; num++ {
				// check for row
				if slices.Contains(row, num) {
					continue
				}

				// check for column
				isValueInCurrentCol := false
				for k := range board {
					if board[k][j] == num {
						isValueInCurrentCol = true
					}
				}
				if isValueInCurrentCol {
					continue
				}

				// check for in current block
				if slices.Contains(getItemsInCurrentBlock(board, i, j), num) {
					continue
				}

				possibleNums = append(possibleNums, num)
			}

			possibility := Value{row: i, col: j, possibilities: possibleNums}
			possibilities = append(possibilities, possibility)
		}
	}
	return possibilities
}

func getItemsInCurrentBlock(board [][]int, row int, col int) []int {
	rowStart := int(row/3) * 3
	colStart := int(col/3) * 3

	items := []int{}
	for i := rowStart; i < rowStart+3; i++ {
		for j := colStart; j < colStart+3; j++ {
			v := board[i][j]
			if v == 0 {
				continue
			}
			items = append(items, v)
		}
	}

	return items
}

func hasWon(board [][]int) bool {
	for _, row := range board {
		for _, v := range row {
			if v == 0 {
				return false
			}
		}
	}
	return true
}

func lo