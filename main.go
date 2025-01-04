package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Value struct {
	row           int
	col           int
	possibilities []int
}

func main() {
	board := readBoard("board.txt")

	const logBoard = false
	const maxIters int = 81
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
		for i := 0; i < len(possibilities)-1; i++ {
			if len(possibilities[i].possibilities) == 0 {
				top = possibilities[i+1]
			}
		}

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
			printBoard(board)
		}
	}

	if hasWon(board) {
		fmt.Println("Completed the board.")
		printBoard(board)
	}

}

func readBoard(name string) [][]int {
	f, err := os.Open(name)
	if err != nil {
		log.Fatalf("Unable to open the file\n")
	}

	result := [][]int{}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()

		arr := strings.Split(line, " ")
		row := []int{}
		for _, item := range arr {
			num, err := strconv.Atoi(item)
			if err != nil {
				log.Fatalf("Unable to convert string to number\n")
			}
			row = append(row, num)
		}
		result = append(result, [][]int{row}...)
	}

	return result
}

func findNewPossibilities(board [][]int) []Value {
	possibilities := []Value{}
	for i, row := range board {
		for j, v := range row {
			currentColumnValues := getItemsInCol(board, j)
			if v != 0 {
				continue
			}

			possibleNums := []int{}
			for num := 1; num <= 9; num++ {
				if slices.Contains(row, num) {
					continue
				}
				if slices.Contains(currentColumnValues, num) {
					continue
				}
				if slices.Contains(getItemsInCurrentBlock(board, i, j), num) {
					continue
				}
				possibleNums = append(possibleNums, num)
			}
			possibilities = append(possibilities, Value{row: i, col: j, possibilities: possibleNums})
		}
	}
	return possibilities
}

func getItemsInCol(board [][]int, col int) []int {
	items := []int{}
	for k := 0; k < 9; k++ {
		v := board[k][col]
		if v == 0 {
			continue
		}
		items = append(items, v)
	}
	return items
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

func printBoard(board [][]int) {
	for i, row := range board {
		if i%3 == 0 || i == 0 {
			for range 25 {
				fmt.Print("-")
			}
			fmt.Println()
		}

		for j, v := range row {
			if (j)%3 == 0 {
				fmt.Print("| ")
			}
			fmt.Printf("%d ", v)

			if j == len(row)-1 {
				fmt.Print("| ")
			}
		}
		fmt.Println()

		if i == len(board)-1 {
			for range 25 {
				fmt.Print("-")
			}
			fmt.Println()
		}

	}
}
