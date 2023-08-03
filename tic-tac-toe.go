package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Board [3][3]string

type Coordinates struct {
	x int
	y int
}

func newBoard() *Board {
	return &Board{
		{".", ".", "."},
		{".", ".", "."},
		{".", ".", "."},
	}
}

func (b Board) checkHorizontal() bool {
	if b[0][0] == b[0][1] && b[0][0] == b[0][2] {
		if b[0][0] != "." && b[0][1] != "." && b[0][2] != "." {
			fmt.Println("top row")
			return true
		}
	}
	if b[1][0] == b[1][1] && b[1][0] == b[1][2] {
		if b[1][0] != "." && b[1][1] != "." && b[1][2] != "." {
			fmt.Println("mid row")
			return true
		}
	}
	if b[2][0] == b[2][1] && b[2][0] == b[2][2] {
		if b[2][0] != "." && b[2][1] != "." && b[2][2] != "." {
			fmt.Println("bot row")
			return true
		}
	}
	return false
}

func (b Board) checkVertical() bool {
	if b[0][0] == b[1][0] && b[0][0] == b[2][0] {
		if b[0][0] != "." && b[1][0] != "." && b[2][0] != "." {
			fmt.Println("lft col")
			return true
		}
	}
	if b[0][1] == b[1][1] && b[0][1] == b[2][1] {
		if b[0][1] != "." && b[1][1] != "." && b[2][1] != "." {
			fmt.Println("mid col")
			return true
		}
	}
	if b[0][2] == b[1][2] && b[0][2] == b[2][2] {
		if b[0][2] != "." && b[1][2] != "." && b[2][2] != "." {
			fmt.Println("rt col")
			return true
		}
	}
	return false
}

func (b Board) checkDiagonal() bool {
	if b[0][0] == b[1][1] && b[0][0] == b[2][2] {
		if b[0][0] != "." && b[1][1] != "." && b[2][2] != "." {
			return true
		}
	}
	if b[0][2] == b[1][1] && b[0][2] == b[2][0] {
		if b[0][2] != "." && b[1][1] != "." && b[2][0] != "." {
			return true
		}
	}
	return false
}

func (b Board) checkWin() bool {
	switch {
	case b.checkHorizontal():
		fmt.Println("won hor")
		return true
	case b.checkVertical():
		fmt.Println("won ver")
		return true
	case b.checkDiagonal():
		fmt.Println("won diag")
		return true
	}

	return false
}

func (b *Board) insert(str string, cords Coordinates) {
	if str != "X" && str != "O" {
		fmt.Println("You must use X or O.")
		return
	}

	if b[cords.y][cords.x] != "." {
		fmt.Println("Please choose open spot.")
		return
	}

	b[cords.y][cords.x] = str

	return
}

func (b Board) print() {
	fmt.Println(b[0])
	fmt.Println(b[1])
	fmt.Println(b[2])
}

func main() {
	turn := 0
	player := 1
	board := newBoard()

	fmt.Println("This is a new board:")
	board.print()

	for {
		rand.Seed(time.Now().UnixNano())
		if player == 1 {
			cor := Coordinates{
				x: rand.Intn(3),
				y: rand.Intn(3),
			}

			board.insert("X", cor)
			turn++

			if board.checkWin() {
				fmt.Printf("After turn %d X's win!\n", turn)
				board.print()
				break
			} else {
				fmt.Printf("Board after turn %d\n", turn)
				board.print()
				player = 2
			}
		}
		if player == 2 {
			cor := Coordinates{
				x: rand.Intn(3),
				y: rand.Intn(3),
			}

			board.insert("O", cor)
			turn++

			if board.checkWin() {
				fmt.Printf("After turn %d O's win!\n", turn)
				board.print()
				break
			} else {
				fmt.Printf("Board after turn %d\n", turn)
				board.print()
				player = 1
			}
		}
	}
}
