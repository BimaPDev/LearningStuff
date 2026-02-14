package main

import "fmt"

func isValidSudoku(board [][]byte) bool {
	seen := make(map[string]bool)
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			cell := board[row][col]
			if cell == '.' {
				continue
			}
			box := (row/3)*3 + (col / 3)
			keyRow := fmt.Sprintf("r%d-%c", row, cell)
			keyCol := fmt.Sprintf("c%d-%c", col, cell)
			keyBox := fmt.Sprintf("b%d-%c", box, cell)

			check := []string{keyRow, keyCol, keyBox}
			for _, k := range check {
				if seen[k] {
					return false
				}
				seen[k] = true
			}
		}
	}
	return true
}

func main() {
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}

	fmt.Println(isValidSudoku(board))
}
