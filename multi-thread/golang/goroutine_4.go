package gorutine

import (
	"fmt"
	"sync"
	"time"
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
	res := make(chan string, 20)
	fmt.Printf("Len=%d, Cap=%d\n", len(res), cap(res))

	for i := 1; i <= 3; i++ {
		wg.Add(1)

		go func(gorutineNum int) {
			defer wg.Done()
			for j := 1; j <= 10; j++ {
				fmt.Println("Goritine", gorutineNum)
				res <- fmt.Sprintf("Gorutine %d numbers %d len channel = %d", gorutineNum, j, len(res))
				time.Sleep(10 * time.Second)
			}
		}(i)
	}

	go func() {
		wg.Wait()
		close(res)
	}()

	fmt.Println("Hi, I'm ready to read from the channel")
	fmt.Printf("Len=%d, Cap=%d\n", len(res), cap(res))
	for num := range res {
		fmt.Println(num)
	}

}
