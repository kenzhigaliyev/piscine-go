package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func isValidLine(x int, y int, matrix_row []string) int {
	present_on_row := 1
	for index := range matrix_row[x] {
		if present_on_row == 1 && index != y && matrix_row[x][index] == matrix_row[x][y] && matrix_row[x][y] != byte('.') {
			fmt.Println("1")
			present_on_row = 0
		}
	}
	return present_on_row
}

func isValidColumn(x int, y int, matrix_column []string) int {
	present_on_column := 1
	if isValidLine(x, y, matrix_column) == 1 {
		for index := range matrix_column[y] {
			if present_on_column == 1 && index != x && matrix_column[y][index] == matrix_column[x][y] && matrix_column[x][y] != byte('.') {
				fmt.Println("2")
				present_on_column = 0
			}
		}
	}
	return present_on_column
}

func isValidSquare(x int, y int, matrix []string) int {
	present_in_square := 1
	if isValidLine(x, y, matrix) == 1 && isValidColumn(x, y, matrix) == 1 {
		square_row := (x / 3) * 3
		square_column := (y / 3) * 3
		for a := square_row; a < square_row+3; a++ {
			for b := square_column; b < square_column+3; b++ {
				if present_in_square == 1 {
					if a != x || b != y {
						if matrix[a][b] == matrix[x][y] && matrix[x][y] != byte('.') {
							fmt.Println("3")
							present_in_square = 0
						}
					}
				}
			}
		}
	}
	return present_in_square
}

func isValid(input []string) int {
	length := len(input)
	if length > 9 || length < 9 {
		fmt.Println("Error")
		fmt.Println("4")
		return 0
	} else {
		for _, numbers := range input {
			length_line := len(numbers)
			if length_line > 9 || length_line < 9 {
				fmt.Println("Error")
				fmt.Println("5")
				return 0
			}
		}
	}

	for x, values := range input {
		for y, digits := range values {
			if digits > '9' && digits < '1' {
				fmt.Println("Error")
				fmt.Println("6")
				return 0
			} else if isValidSquare(x, y, input) == 0 {
				fmt.Println("Error")
				fmt.Println("7")
				return 0
			}
		}
	}
	return 1
}

func PrintMatrix(matrix []string) {
	for _, row := range matrix {
		for index, column := range row {
			z01.PrintRune(column)
			if index < len(row)-1 {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}
}

func PointCounter(matrix []string) int {
	counter := 0
	for _, row := range matrix {
		for _, val := range row {
			if val == '.' {
				fmt.Println("8")
				counter++
			}
		}
	}
	return counter
}

func Sudoku(x int, y int, space int, counter int, matrix []string) {
	if space == counter {
		PrintMatrix(matrix)
		return
	} else {
		row := []rune(matrix[x])
		if row[y] == '.' {
			for a := '1'; a <= '9'; a++ {
				row[y] = rune(a)
				matrix[x] = string(row)
				fmt.Println("9")
				if isValidSquare(x, y, matrix) == 1 {
					fmt.Println("10")
					if x == 8 {
						counter++
						// y = 0
						Sudoku(x, y+1, space, counter, matrix)
						fmt.Println("11")
					} else {
						counter++
						// y++
						Sudoku(x+1, 0, space, counter, matrix)
						fmt.Println("12")
					}
				}
			}
			row[y] = '.'
			matrix[x] = string(row)
			fmt.Println("13")
		} else if y < 8 {
			y = 0
			Sudoku(x+1, y, space, counter, matrix)
			fmt.Println("14")
		} else {
			y++
			Sudoku(x, y, space, counter, matrix)
			fmt.Println("15")
		}

	}
}

func main() {
	matrix := os.Args[1:]
	x := 0
	y := 0
	counter_of_spaces := 0
	empty_spaces := PointCounter(matrix)
	if isValid(matrix) == 1 {
		Sudoku(x, y, empty_spaces, counter_of_spaces, matrix)
	}
}
