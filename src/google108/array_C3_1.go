package google108

import (
	"fmt"
)

func study_C3_1() {
	//	arr := [3]int{-1, -100, 1000}
	//	g.GetMax(arr)
	//	g.Sumarr(arr)
	//	g.Printarr(arr)
	//	fmt.Println(arr)
	//	fmt.Println("===rray is value reference type in GOlang,and need specify total element Num")
	//	g.PrintarrRef(&arr)
	//	fmt.Println(arr)
}

func Printarr(arr []int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func PrintarrRef(arr *[3]int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func GetMax(arr [3]int) {
	maxi := 1
	maxvalue := -1
	for i, v := range arr {
		if v > maxvalue {
			maxi, maxvalue = i, v
		}
	}
	fmt.Printf("arr[%d]=%d is Max\n", maxi, maxvalue)
}

func Sumarr(arr [3]int) {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	fmt.Printf("arr %v sum var is %d\n", arr, sum)
}
