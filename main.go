package main

import (
	"fmt"
	"math"
	"os"

	"golang.org/x/term"
)

func main() {
	term.MakeRaw(int(os.Stdin.Fd()))
	clear()

	// oLen := 30.0
	oHeight := 30.0
	oWidth := 40.0

	rot := math.Pi / 6

	for {
		b := make([]byte, 1)
		os.Stdin.Read(b)
		clear()

		switch b[0] {
		case 'W', 'w':
			rot += 0.1
		case 'S', 's':
			rot -= 0.1
		case 3:
			os.Exit(0)
		}

		print(func() (width float64, pWidth float64) {
			// todo: calculate width

			width = oWidth
			pWidth = 1
			return
		}, func() (height float64, pHeight float64) {
			height = oHeight * math.Abs(math.Cos(rot))
			pHeight = (oHeight - height) / 2
			return
		})

	}
}

func clear() {
	fmt.Print("\033[H")
	fmt.Print("\033[2J")
}

func print(
	calcWidth func() (float64, float64),
	calcHeight func() (float64, float64),
) {
	height, pHeight := calcHeight()
	for ph := 0; ph < int(pHeight); ph++ {
		fmt.Println()
	}

	for h := 0; h < int(height); h++ {
		width, pWidth := calcWidth()
		for pw := 0; pw < int(pWidth); pw++ {
			fmt.Print(" ")
		}

		for w := 0; w < int(width); w++ {
			fmt.Print("#")
		}
		fmt.Println()
	}
}
