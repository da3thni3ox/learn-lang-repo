package main

import "fmt"

func main() {
	n := 10

	sum := 0

	for i := 1; i <= n; i++ {
		sum += i
	}

	fmt.Println(sum)
}
