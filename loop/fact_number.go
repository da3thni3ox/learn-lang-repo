package main

import "fmt"

func main() {
	n := 3

	fact := 0

	for i := 1; i <= n; i++ {
		num := i

		fact = num * i
	}

	fmt.Println(fact)
}
