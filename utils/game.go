package utils

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/alem-platform/ap"
)

func GameLoop(matrix [][]rune, h, w, bombCounts int) {
	moveCount := 0
	rand.Seed(time.Now().UnixNano())
	intMatrix := convertRuneMatrixToInt(matrix)
	revealed := make([][]bool, h)
	for i := 0; i < h; i++ {
		revealed[i] = make([]bool, w)
	}
	for {
		displayMatrix(intMatrix, h, w, revealed)
		var x, y int
		PrintString("Enter coordinates: ")
		for {
			_, err := fmt.Scanf("%d%d", &x, &y)
			if err != nil {
				PrintString("Input error! Invalid input. Enter the coordinates again:\n ")
				continue
			} else if x < 0 || x > h || y < 0 || y > w {
				PrintString("Input error! Coordinates out of bounds. Please enter valid coordinates.\n")
				continue
			} else {
				break
			}
		}
		x--
		y--
		ap.PutRune('\n')

		if revealed[x][y] {
			PrintString("You have already stepped on this cell. Please provide another...\n")
			continue
		}
		if moveCount == 0 && intMatrix[x][y] == -1 {
			matrix = relocateBomb(matrix, h, w, x, y)
			intMatrix = convertRuneMatrixToInt(matrix)
			cascadeReveal(intMatrix, revealed, x, y)
			moveCount++
			continue
		}
		moveCount++
		if intMatrix[x][y] == -1 {
			revealAllBombs(intMatrix, revealed, h, w)
			displayMatrix(intMatrix, h, w, revealed)
			PrintString(StyleString("Game Over!", "\033[31m", true))
			ap.PutRune('\n')

			PrintStatistics(h, w, bombCounts, moveCount)
			break
		}
		cascadeReveal(intMatrix, revealed, x, y)
		if checkWin(intMatrix, revealed, h, w, bombCounts) {
			displayMatrix(intMatrix, h, w, revealed)

			PrintString(StyleString("You Win!", "\033[32m", true))
			ap.PutRune('\n')

			PrintStatistics(h, w, bombCounts, moveCount)
			break
		}
	}
}

func PrintStatistics(h, w, bombCounts, moveCount int) {
	PrintString("Your statistics:")
	ap.PutRune('\n')
	PrintString("- Field size: ")
	PrintInt(h)
	ap.PutRune('x')
	PrintInt(w)
	ap.PutRune('\n')
	PrintString("- Number of bombs: ")
	PrintInt(bombCounts)
	ap.PutRune('\n')
	PrintString("- Number of moves: ")
	PrintInt(moveCount)
	ap.PutRune('\n')
}

func relocateBomb(matrix [][]rune, h, w, x, y int) [][]rune {
	for {
		newX := rand.Intn(h)
		newY := rand.Intn(w)
		if matrix[newX][newY] != '*' {
			matrix[newX][newY] = '*'
			matrix[x][y] = '.'
			break
		}
	}
	return matrix
}

func convertRuneMatrixToInt(matrix [][]rune) [][]int {
	rows := len(matrix)
	cols := len(matrix[0])

	intGrid := make([][]int, rows)
	for i := 0; i < rows; i++ {
		intGrid[i] = make([]int, cols)
		for j := 0; j < cols; j++ {
			if matrix[i][j] == '*' {
				intGrid[i][j] = -1
				continue
			}
			count := 0
			for di := -1; di <= 1; di++ {
				for dj := -1; dj <= 1; dj++ {
					ni := i + di
					nj := j + dj
					if ni >= 0 && ni < rows && nj >= 0 && nj < cols && !(di == 0 && dj == 0) {
						if matrix[ni][nj] == '*' {
							count++
						}
					}
				}
			}
			intGrid[i][j] = count
		}
	}
	return intGrid
}

func revealAllBombs(intMatrix [][]int, revealed [][]bool, h, w int) {
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if intMatrix[i][j] == -1 {
				revealed[i][j] = true
			}
		}
	}
}

func checkWin(intMatrix [][]int, revealed [][]bool, h, w, bombCounts int) bool {
	totalCells := h * w
	unrevealedSafeCells := totalCells - bombCounts

	count := 0
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if revealed[i][j] {
				count++
			}
		}
	}

	return count == unrevealedSafeCells
}

func cascadeReveal(matrix [][]int, revealed [][]bool, x, y int) {
	rows := len(matrix)
	cols := len(matrix[0])

	if x < 0 || x >= rows || y < 0 || y >= cols || revealed[x][y] {
		return
	}

	revealed[x][y] = true
	if matrix[x][y] == 0 {
		for di := -1; di <= 1; di++ { 
			for dj := -1; dj <= 1; dj++ { 
				if di != 0 || dj != 0 { 
					cascadeReveal(matrix, revealed, x+di, y+dj)
				}
			}
		}
	}
}
