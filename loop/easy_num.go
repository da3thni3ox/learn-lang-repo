package main

import "fmt"

func main() {
	checkNumber := 2
	isPrime := true

	if checkNumber <= 1 {
		isPrime = false
		fmt.Printf("Number %d is isPrime %v", checkNumber, isPrime)
		return
	}

	for i := 2; i < checkNumber; i++ {
		if checkNumber%i == 0 {
			isPrime = false
			break
		}
	}

	fmt.Printf("Number %d is easy %v", checkNumber, isPrime)
}
