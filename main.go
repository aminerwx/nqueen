package main

import "fmt"

type Queen struct {
	x int
	y int
}

func main() {
	matrix := [][]int{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}
	out := [][]int{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}
	place(matrix, 0, &out)
	boardPrint(out)
	// fmt.Println(flag)
}

func nqueen(matrix [][]int, x, y int) {
	for i := 0; i < len(matrix); i++ {
		x++
		y++
	}
}

func board(n int) [][]int {
	return make([][]int, n)
}

func place(matrix [][]int, row int, output *[][]int) {
	if row == len(matrix)-1 {
		var r []int
		for i := 0; i < len(matrix); i++ {
			for j := 0; j < len(matrix); j++ {
				if matrix[i][j] == 9 {
					r = append(r, j+1)
					break
				}
			}
		}
		*output = append(*output, r)
		return
	}
	for j := 0; j < len(matrix); j++ {
		if check(matrix, row, j) {
			matrix[row][j] = 9
			place(matrix, row+1, output)
			matrix[row][j] = 0
		}
	}
}

func boardPrint(board [][]int) {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board); j++ {
			fmt.Printf("%d ", board[i][j])
		}
		fmt.Println()
	}
}

func check(matrix [][]int, x, y int) bool {
	if len(matrix) == 1 {
		fmt.Printf("found at: (%d, %d)\n", x, y)
		return matrix[0][0] == 9
	}

	// Set diagonal from top-left to bottom-right
	for i := -min(x, y); i <= min(len(matrix)-1-x, len(matrix)-1-y); i++ {
		if x+i == x && y+i == y {
			continue
		}
		if matrix[x+i][y+i] == 9 {
			fmt.Printf("found at: (%d, %d)\n", x+i, y+i)
			return false
		}
	}

	// Set diagonal from top-right to bottom-left
	for i := -min(x, len(matrix)-1-y); i <= min(len(matrix)-1-x, y); i++ {
		if x+i == x && y-i == y {
			continue
		}
		if matrix[x+i][y-i] == 9 {
			fmt.Printf("found at: (%d, %d)\n", x+i, y-i)
			return false
		}
	}

	// check rows from top to bottom
	for i := 0; i < len(matrix); i++ {
		if i != x && matrix[i][y] == 9 {
			fmt.Printf("found at: (%d, %d)\n", i, y)
			return false
		}
	}

	// check columns from left to right
	for i := 0; i < len(matrix); i++ {
		if i != y && matrix[x][i] == 9 {
			fmt.Printf("found at: (%d, %d)\n", i, y)
			return false
		}
	}
	fmt.Printf("nothing at (%d, %d)\n", x, y)
	return true
}
