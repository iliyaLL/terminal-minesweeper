package utils

import (
	"fmt"

	"github.com/alem-platform/ap"
)

func PrintInt(n int) {
	digits := []rune{}
	for n > 0 {
		digits = append([]rune{rune('0' + n%10)}, digits...)
		n /= 10
	}

	for _, r := range digits {
		ap.PutRune(r)
	}
}

func PrintString(str string) {
	for _, r := range str {
		ap.PutRune(r)
	}
}

func IsValidLine(line string) bool {
	for _, r := range line {
		if r != '.' && r != '*' {
			return false
		}
	}
	return true
}

func InputLine(w int) ([]rune, int) {
	var line string
	bombCounts := 0
	for {
		_, err := fmt.Scanf("%s", &line)
		if err != nil {
			PrintString("Input error! Invalid input. Re-write the line\n")
			continue
		} else if len(line) != w || len(line) > w || len(line) < w {
			PrintString("Input error! Invalid input. Re-write the line\n")
			continue
		} else if !IsValidLine(line) {
			PrintString("Input error! Enter the correct line with (* or .). Re-write the line\n")
		} else {
			break
		}
	}
	for _, r := range line {
		if r == '*' {
			bombCounts++
		}
	}
	return []rune(line), bombCounts
}

func StyleString(text string, color string, bold bool) string {
	var attr string
	if bold {
		attr = "\033[1m"
	} else {
		attr = ""
	}

	return color + attr + text + "\033[0m"
}
