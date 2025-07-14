package main

import "fmt"

func main() {
	var mltable [][]int
	var tempMlTable []int

	size := 2
	for i := 1; i <= size; i++ {
		tempMlTable = make([]int, 3, 3)
		for j := 1; j <= size; j++ {
			res := i * j
			tempMlTable[j-1] = res
		}

		mltable = append(mltable, tempMlTable)
	}

	fmt.Println(mltable)
}

//https://www.codewars.com/kata/534d2f5b5371ecf8d2000a08/train/go
//Multiplication table
