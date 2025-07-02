package main

import "fmt"

func drawTriagle(row int, maxWidth int) {

	for i := 0; i < row; i++ {
		spaces := maxWidth - (2*i + 1)

		for space := 0; space < spaces/2; space++ {
			fmt.Print(" ")
		}
		for star := 0; star < 2*i+1; star++ {
			fmt.Print("*")
		}
		fmt.Print("\n")
	}
}

func main() {
	level := 3

	lastLevelRows := level + 1
	maxStars := 2*(lastLevelRows-1) + 1

	for i := 1; i <= level; i++ {
		drawTriagle(i+1, maxStars)
	}

	for i := 0; i < 1; i++ {
		padding := (maxStars - 3) / 2
		for space := 0; space < padding; space++ {
			fmt.Print(" ")
		}

		fmt.Print("| |\n")
		if i == 0 {
			for space := 0; space < padding; space++ {
				fmt.Print(" ")
			}
			fmt.Println("---")
		}

	}
}

/*
	TODO Расположить игрушки

	Первая строка 1
	вторая строка 3
	третья строка 5


*/
