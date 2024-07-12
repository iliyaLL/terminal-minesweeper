package utils

import (
	"math/rand/v2"

	"github.com/alem-platform/ap"
)

func ColorPrint(val int) {
	colors := []string{

		"\033[34m", // blue
		"\033[32m", // green
		"\033[31m", // red
		"\033[35m", // magenta
		"\033[33m", // yellow
		"\033[36m", // light blue
		"\033[37m", // white
		"\033[90m", // black
	}

	if val >= 1 && val <= 8 {
		color := colors[val-1]
		for _, c := range color {
			ap.PutRune(c)
		}
		ap.PutRune(rune(val + '0'))
		reset := "\033[0m"
		for _, c := range reset {
			ap.PutRune(c)
		}
	} else {
		ap.PutRune(rune(val + '0'))
	}
}

func RandomMapGeneration(h, w int) [][]rune {
	input := make([][]rune, h)
	totalCells := h * w
	bombsToFill := totalCells * 30 / 100
	for i := 0; i < h; i++ {
		input[i] = make([]rune, w)
		for j := 0; j < w; j++ {
			input[i][j] = '.'
		}
	}
	for b := 0; b < 2; b++ {
		for {
			i := rand.IntN(h)
			j := rand.IntN(w)
			if input[i][j] == '.' {
				input[i][j] = '*'
				break
			}
		}
	}
	bombsPlaced := 2
	for bombsPlaced < bombsToFill {
		i := rand.IntN(h)
		j := rand.IntN(w)
		if input[i][j] == '.' {
			input[i][j] = '*'
			bombsPlaced++
		}
	}
	return input
}
