package main

import (
	"crunch02/utils"
	"fmt"
)

func main() {
	var inputError string = utils.StyleString("Input Error! ", "\033[31m", true)
	utils.PrintString("Welcome to " + utils.StyleString("Minesweeper Game", "\033[36m", true) + "!\n\n")
	utils.PrintString("Choose a mode: \n1. Enter a custom map \n2. Generate a random map\nEnter your choice:")
	var answer int
	for {
		_, err := fmt.Scanf("%d", &answer)
		if err != nil {
			utils.PrintString(inputError + "Please enter the INTEGER number!\n")
			continue
		} else if answer < 1 || answer > 2 {
			utils.PrintString(inputError + "Please enter the correct number (1 or 2)!\n")
			continue
		} else {
			break
		}
	}
	var h, w int
	if answer == 1 {
		utils.PrintString("Enter the size of the grid:")
		for {
			_, err := fmt.Scanf("%d %d", &h, &w)
			if err != nil {
				utils.PrintString(inputError + "Please enter 2 numbers more! Enter the size again:\n")
				continue
			} else if h <= 2 && w > 2 {
				utils.PrintString(inputError + "Height should be more than 3! Enter the size again:\n")
				continue
			} else if w <= 2 && h > 2 {
				utils.PrintString(inputError + "Width should be more than 3! Enter the size again:\n")
				continue
			} else if h <= 2 && w <= 2 {
				utils.PrintString(inputError + "Height and width should be nore that 3! Enter the size again:\n")
			} else {
				break
			}
		}
		utils.PrintString("Please input state of the cells: \n'.' empty cell. \n'*' bomb\nOne line of input is state of one line of map:\n")
		bombCounts := 0
		matrix := make([][]rune, h)
		for i := 0; i < h; i++ {
			line, bombn := utils.InputLine(w)
			bombCounts += bombn
			matrix[i] = line
		}
		if bombCounts < 2 {
			panic("Invalid grid. The amount of bombs is less than 2")
		}

		utils.GameLoop(matrix, h, w, bombCounts)
	} else if answer == 2 {
		utils.PrintString("Enter the size of the grid:")
		for {
			_, err := fmt.Scanf("%d %d", &h, &w)
			if err != nil {
				utils.PrintString(inputError + "Please enter 2 numbers more! Enter the size again:\n ")
				continue
			} else if h <= 2 && w > 2 {
				utils.PrintString(inputError + "Height should be more than 3! Enter the size again:\n")
				continue
			} else if w <= 2 && h > 2 {
				utils.PrintString(inputError + "Width should be more than 3! Enter the size again:\n")
				continue
			} else if h <= 2 && w <= 2 {
				utils.PrintString(inputError + "Height and width should be nore that 3! Enter the size again:\n")
			} else {
				break
			}
		}
		matrix := utils.RandomMapGeneration(h, w)

		utils.GameLoop(matrix, h, w, (h*w)*30/100)
	}
}
