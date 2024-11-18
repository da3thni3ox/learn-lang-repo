package gorutine

import (
	"fmt"
	"time"
)

func gorutineFirstTask() {

	for i := 1; i <= 5; i++ {
		go func() {
			for i := 1; i <= 5; i++ {
				fmt.Println(i)

				time.Sleep(1 * time.Second)
			}
		}()

	}

	time.Sleep(6 * time.Second)
}
