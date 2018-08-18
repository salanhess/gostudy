package main_

import (
	"fmt"
	g "google108"
)

func beforeC3() {
	var a, b = 2, 3
	fmt.Printf("Current var a，b is %d, %d \n", a, b)
	//	g.Swap(&a, &b)
	//	g.SwapNoEffect(a, b)
	a, b = g.SwapReturn(a, b)
	fmt.Printf("After swap var a，b is %d, %d \n", a, b)
	fmt.Println("=====")
	fmt.Printf("Current var a is %d\n", a)
	g.Try_nopointer(a)
	fmt.Printf("Current var a is %d\n", a)
	g.Try_pointer(&a)
	fmt.Printf("Current var a is %d\n", a)

	//	q, r := div(11, 3)
	//	fmt.Println(q, r)
	//	sum(1, 2, 3)
	//	//fmt.Println(apply(pow, 3, 4))
	//	//lambda fun
	//	fmt.Println(apply(func(a, b int) int {
	//		return int(math.Pow(float64(a), float64(b)))
	//	}, 3, 4))
	//	fmt.Println(eval(1, 2, "+"))
	//	if _, err := eval(1, 2, "%"); err != nil {
	//		fmt.Println("Error!!")
	//	}
	//	//	forever("abc")
	//	const filename = "abc.txt"
	//	printfile(filename)
	//	//	readfile(filename)
	//	fmt.Println(grade(50))
	//	fmt.Println(grade(60))
	//	fmt.Println(grade(90))
	//	//	fmt.Println(grade(1100))

}

func Study_C3_5_leetcode() {
	g.LengthofNonRepeating("abc")
	g.LengthofNonRepeating("aabbccdd")
}

func Study_C3_6() {
	Study_C3_5_leetcode()
	g.Study_C3_6()
	g.LengthofNonRepeating_Rune("你不是我是大傻逼大傻逼")
}
func main() {

}
