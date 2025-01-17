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
	// domainBasedSolution(board, logBoard)
	backtrackingSolution(board)
}

func backtrackingSolution(board [][]int) bool {
	if hasWon(board) {
		printBoard(board)
		return true
	}

	for i, row := range board {
		for j, v := range row {
			if v != 0 {
				continue
			}

			cellDomain := getCellDomain(board, i, j)

			for _, num := range cellDomain {
				boardCopy := make([][]int, len(board))
				for i := range board {
					boardCopy[i] = make([]int, len(board[i]))
					copy(boardCopy[i], board[i])
				}
				if isInvalidatingAnyDomain(boardCopy, i, j, num) {
					continue
				}

				if backtrackingSolution(board) {
					return true
				}
			}
		}
	}
	return false
}

func isInvalidatingAnyDomain(board [][]int, row int, col int, num int) bool {
	board[row][col] = num

	for i := range board {
		for j := range row {
			if len(getCellDomain(board, i, j)) == 0 {
				return true
			}
		}
	}
	return false
}

func domainBasedSolution(board [][]int, logBoard bool) {
	for !hasWon(board) {
		possibilities := findNewPossibilities(board)
		slices.SortFunc(possibilities, func(a Value, b Value) int {
			return len(a.possibilities) - len(b.possibilities)
		})

		if len(possibilities) == 0 {
			break
		}

		if len(possibilities[0].possibilities) > 1 {
			fmt.Println("Possibilities")
			for _, p := range possibilities {
				fmt.Println(p)
			}
			log.Fatalf("Please enter a valid board. The given board has no solution.")
		}

		for _, p := range possibilities {
			if len(p.possibilities) == 0 {
				log.Fatalf("No solution is possible for given board. Please enter a valid board.")
			}

			if len(p.possibilities) == 1 {
				fmt.Printf("Change %d %d to %d\n", p.row+1, p.col+1, p.possibilities[0])
				board[p.row][p.col] = p.possibilities[0]
			}
		}

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
			if v != 0 {
				continue
			}
			possibleNums := getCellDomain(board, i, j)
			possibilities = append(possibilities, Value{row: i, col: j, possibilities: possibleNums})
		}
	}
	return possibilities
}

func getCellDomain(board [][]int, row int, col int) []int {
	domain := []int{}
	rowValues := board[row]
	colValues := getItemsInCol(board, col)
	blockValues := getItemsInCurrentBlock(board, row, col)
	for num := 1; num <= 9; num++ {
		if slices.Contains(rowValues, num) || slices.Contains(colValues, num) || slices.Contains(blockValues, num) {
			continue
		}
		domain = append(domain, num)
	}
	return domain
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
