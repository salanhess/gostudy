package main

import "fmt"

func select_sort(nums []int) []int {
	n := len(nums)
	for i, _ := range nums {
		// Wrong 1
		//		min := nums[i]
		//		for j := 1; j < n-i; j++ {
		//			fmt.Printf("i=%d,j=%d,cmp nums[j]:%d with min: %d \n", i, j, j, nums[minIndex])
		//			if nums[j] < min {
		//				min, nums[j] = nums[j], min
		//			}
		//			if nums[j] < nums[minIndex] {
		//				nums[minIndex], nums[j] = nums[j], nums[minIndex]
		//			}
		//		}
		// Correct Sample
		//	minIndex := i
		//	for j := i + 1; j < n; j++ {
		//	      fmt.Printf("i=%d,j=%d,cmp nums[j]:%d with min: %d \n", i, j, nums[j], nums[minIndex])
		//		if nums[j] < nums[minIndex] {
		//			nums[minIndex], nums[j] = nums[j], nums[minIndex]
		//		}
		//	}
		minIndex := i
		for j := i + 1; j < n; j++ {
			fmt.Printf("i=%d,j=%d,cmp nums[j]:%d with min: %d \n", i, j, nums[j], nums[minIndex])
			if nums[j] < nums[minIndex] {
				nums[minIndex], nums[j] = nums[j], nums[minIndex]
			}
		}
	}
	return nums
}
func main() {
	fmt.Println("vim-go")
	fmt.Println(select_sort([]int{1, 2, 4, 3}))
	fmt.Println(select_sort([]int{11, 2, 14, 3}))
	fmt.Println(select_sort([]int{100, 20, 4, 3}))
}
