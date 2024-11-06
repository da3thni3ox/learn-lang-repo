package main

import (
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

func sum(group int) {
	var sum int
	for i := 1; i <= 10; i++ {
		sum += i
	}
	defer wg.Done()
}

func main() {

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go sum(i)

	}

	wg.Wait()
}

// Первая горутина считает сумму чисел от 1 до 10
// Вторая горутина считает сумму чисел от 11 до 20
// Треться горутина считает сумму чисел от 21 до 30
