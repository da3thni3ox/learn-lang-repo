package main

import "fmt"

func main() {
	prices := []int{1, 1, 1, 1, 1}

	buyPriceMin := prices[0]
	benefit := 0
	maxBenefit := 0
	for i := 0; i < len(prices); i++ {
		currentPrice := prices[i]

		if buyPriceMin > currentPrice {
			buyPriceMin = currentPrice
		}

		benefit = currentPrice - buyPriceMin

		if benefit > maxBenefit {
			maxBenefit = benefit
		}
	}

	fmt.Println(maxBenefit)

}
