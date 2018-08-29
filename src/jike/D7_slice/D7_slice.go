package main

import "fmt"

func check_slice(s []int) {
	fmt.Printf("len(slice)=%d,cap(slice)=%d\n", len(s), cap(s))
}

func print_slice(s []int) {
	fmt.Printf("slice=%v\n", s)
}

func main() {
	fmt.Println("[Case1]Init slice with/without specific cap:")
	s1 := make([]int, 3)
	s2 := make([]int, 3, 8)
	check_slice(s1)
	check_slice(s2)
	fmt.Println("[Case2]View slice via new slice(eg. see from index n+1 to end) :")
	s1 = []int{1, 2, 3, 4, 5}
	s2 = s1[2:5]
	check_slice(s2)
	print_slice(s2)
	fmt.Println("[Case3]append  slice to check cap change")
	s3 := append(s2, 6, 7)
	check_slice(s3)
	print_slice(s3)
	s3 = append(s3, 10, 11, 12)
	check_slice(s3)
	print_slice(s3)
	fmt.Println("[Case4]View slice after decrease cap")
	//	s4 := append(s3[1:2], s3[4:cap(s3)]...)
	s3 = append(s3[1:2], s3[4:5]...)
	check_slice(s3)
	print_slice(s3)
}
