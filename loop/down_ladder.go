package main

import "fmt"

func main() {

	n := 5

	for i := n; i >= 1; i-- {
		for j := 1; j <= i; j++ {
			fmt.Print("x")
		}
		fmt.Print("\n")
	}
}
