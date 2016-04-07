package main

import (
	"encoding/json"
	"fmt"
	"math"
)

// Board represents the position of the Queens on the Board
type Board struct {
	Queens [8]int
}

func main() {
	var board Board

	// initialize
	for i := 0; i < 8; i++ {
		board.Queens[i] = -1
	}

	play(board, 0)
}

func play(board Board, column int) {

	for i := 0; i < 8; i++ {
		board.Queens[column] = i

		if !validate(board, column) {
			continue
		}

		if column == 7 {
			output(board)
			continue
		}

		play(board, column+1)
	}

	// reset
	board.Queens[column] = -1
}

func validate(board Board, column int) bool {

	// scan for entries in the same row
	for col := 0; col < column; col++ {
		if board.Queens[col] == board.Queens[column] {
			return false
		}
	}

	x := column
	y := board.Queens[column]

	// check diagonals
	for col := 0; col < column; col++ {
		deltaX := x - col
		deltaY := y - board.Queens[col]

		if math.Abs(float64(deltaX)) == math.Abs(float64(deltaY)) {
			return false
		}
	}

	return true
}

func output(board Board) {

	result := make([]int, 8)
	for i := 0; i < 8; i++ {
		result[i] = board.Queens[i] + 1
	}

	txt, _ := json.Marshal(result)
	fmt.Println(string(txt))
}
