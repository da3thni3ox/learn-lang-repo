package gorutine

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

func sum(start, end int, res chan int) {
	var sum int
	for i := start; i <= end; i++ {
		sum += i
	}
	defer wg.Done()

	res <- sum
}

func main() {
	res := make(chan int, 3)

	start := 0
	end := 0
	for i := 1; i <= 3; i++ {
		wg.Add(1)

		switch i {
		case 1:
			start, end = 1, 10
		case 2:
			start, end = 11, 20
		case 3:
			start, end = 21, 30
		}
		go sum(start, end, res)

	}

	wg.Wait()
	close(res)

	var totalSum int
	for sum := range res {
		totalSum += sum
	}

	fmt.Println(totalSum)
}

// Первая горутина считает сумму чисел от 1 до 10
// Вторая горутина считает сумму чисел от 11 до 20
// Треться горутина считает сумму чисел от 21 до 30
