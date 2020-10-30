package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println(max1(2, 3, 5, 6))
	fmt.Println(max2(2, 3, 5, 6))
	fmt.Println(min1(2, 3, 5, 6))
	fmt.Println(min2(2, 3, 5, 6))
}

func max1(nums ...int) int {
	if len(nums) == 0 {
		log.Println("max1() called with no args.")
		return 0
	}
	res := nums[0]
	for _, num := range nums[1:] {
		res = max(res, num)
	}
	return res
}

func max2(n int, nums ...int) int {
	res := n
	for _, num := range nums {
		res = max(res, num)
	}
	return res
}

func max(num1, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}

func min1(nums ...int) int {
	if len(nums) == 0 {
		log.Println("min1() called with no args.")
		return 0
	}
	res := nums[0]
	for _, num := range nums[1:] {
		res = min(res, num)
	}
	return res
}

func min2(n int, nums ...int) int {
	res := n
	for _, num := range nums {
		res = min(res, num)
	}
	return res
}

func min(num1, num2 int) int {
	if num1 < num2 {
		return num1
	}
	return num2
}
