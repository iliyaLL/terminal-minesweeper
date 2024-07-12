package utils

import (
	"fmt"
	"strconv"

	"github.com/alem-platform/ap"
)

func drawHorizontalNotation(w int) {
	res := ""
	for i := 1; i <= w; i++ {
		if i < 10 {
			res += "       "
			res += strconv.Itoa(i)
		} else if i < 100 {
			res += "       "
			res += strconv.Itoa(i)
		} else {
			res += "       "
			res += strconv.Itoa(i)
		}
	}
	res += "\n    "
	for i := 0; i < w*8-1; i++ {
		res += "_"
	}
	res += "\n"
	fmt.Print(res)
}

func displayMatrix(matrix [][]int, h, w int, revealed [][]bool) {
	drawHorizontalNotation(w)

	for i := 0; i < h; i++ {
		PrintString("   ")
		for j := 0; j < w; j++ {
			if revealed[i][j] {
				PrintString("|       ")
			} else {
				PrintString("|XXXXXXX")
			}
		}
		PrintString("|\n")

		if i+1 < 10 {
			PrintInt(i + 1)
			PrintString("  ")
		} else if i < 100 {
			PrintInt(i + 1)
			ap.PutRune(' ')
		} else {
			PrintInt(i + 1)
		}

		for j := 0; j < w; j++ {
			if revealed[i][j] {
				val := matrix[i][j]
				if val == -1 {
					PrintString("|   *   ")
				} else {
					if val == 0 {
						PrintString("|       ")
					} else {
						PrintString("|   ")
						ColorPrint(val)
						PrintString("   ")
					}
				}
			} else {
				PrintString("|XXXXXXX")
			}
		}
		PrintString("|\n   ")

		for j := 0; j < w; j++ {
			if revealed[i][j] {
				PrintString("|_______")
			} else {
				PrintString("|-------")
			}
		}
		PrintString("|\n")
	}
}
