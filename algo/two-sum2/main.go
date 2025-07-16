package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	target := 9

	leftIdx := 0
	rightIdx := len(nums) - 1
	resNums := []int{}

	for leftIdx < rightIdx {

		if nums[leftIdx]+nums[rightIdx] < target {
			leftIdx++
		}

		if nums[leftIdx]+nums[rightIdx] > target {
			rightIdx--
		} else if nums[leftIdx]+nums[rightIdx] == target {
			resNums = append(resNums, leftIdx, rightIdx)
			break
		}
	}
	fmt.Println(resNums)
}
