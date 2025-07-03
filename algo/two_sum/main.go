package main

import "fmt"

func twoSum(nums []int, target int) []int {

	numsIdx := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		compliment := target - nums[i] // Второе число которое мне нужно найти

		if j, found := numsIdx[compliment]; found {
			return []int{i, j}

		}

		numsIdx[nums[i]] = i
	}

	return nil
}

func main() {

	numbers := []int{3, 2, 4}
	target := 6

	res := twoSum(numbers, target)

	fmt.Println(res)
}
