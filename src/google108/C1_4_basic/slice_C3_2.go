package google108

import (
	"fmt"
)

func Printslice(arr []int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func PrintsliceRef(s1 []int) {
	arr := s1[:]
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func GetsliceMax(arr []int) {
	maxi := 1
	maxvalue := -1
	for i, v := range arr {
		if v > maxvalue {
			maxi, maxvalue = i, v
		}
	}
	fmt.Printf("arr[%d]=%d is Max\n", maxi, maxvalue)
}

func Sumslice(arr []int) {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	fmt.Printf("arr %v sum var is %d\n", arr, sum)
}
