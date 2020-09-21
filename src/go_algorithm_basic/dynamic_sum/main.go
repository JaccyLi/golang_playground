package main

import "fmt"

func dynamicSum(nums []int) []int {
	runningSum := make([]int, len(nums))
	for i, _ := range nums {
		for j := 0; j <= i; j++ {
			runningSum[i] += nums[j]
		}
	}
	fmt.Println(runningSum)
	return runningSum
}

func main() {
	var nums []int = []int{3, 4, 1, 2, 6, 7, 22, 4353, 1211}
	fmt.Println("Input array: ", nums)
	fmt.Printf("Output     : %v", dynamicSum(nums))
}
