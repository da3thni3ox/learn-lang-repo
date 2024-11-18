package grorutine

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

func main() {
	res := make(chan string)

	for i := 1; i <= 3; i++ {
		wg.Add(1)

		go func(gorutineNum int) {
			defer wg.Done()
			for j := 1; j <= 10; j++ {
				res <- fmt.Sprintf("Gorutine %d numbers %d", gorutineNum, j)
			}
		}(i)
	}

	go func() {
		wg.Wait()
		close(res)
	}()

	for num := range res {
		fmt.Println(num)
	}

}
