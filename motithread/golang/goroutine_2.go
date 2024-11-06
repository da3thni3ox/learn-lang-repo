package main

import (
	"fmt"
	"sync"
)

// func factorial(num int) {
// 	var res int = 1
// 	for i := 1; i <= num; i++ {
// 		res *= i
// 	}

// 	fmt.Println(res)
// }

var wg sync.WaitGroup

func sum(start, end int) {
	var sum int
	for i := start; i <= end; i++ {
		sum += i
	}
	defer wg.Done()

	fmt.Println(sum)
}

func main() {

	start := 0
	end := 0
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		if i == 1 {
			start = 1
			end = 10
		}

		if i == 2 {
			start = 11
			end = 20
		}

		if i == 3 {
			start = 21
			end = 30
		}
		go sum(start, end)

	}

	wg.Wait()
}

// Первая горутина считает сумму чисел от 1 до 10
// Вторая горутина считает сумму чисел от 11 до 20
// Треться горутина считает сумму чисел от 21 до 30
