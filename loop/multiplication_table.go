package main

import "fmt"

func main() {
	n := 10

	for i := 1; i < n; i++ {
		for j := 1; j < n; j++ {
			res := i * j

			fmt.Printf(" %d * %d = %d\t", j, i, res)
		}
		fmt.Print("\n")
	}
}
