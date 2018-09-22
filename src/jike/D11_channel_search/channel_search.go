package main

import "fmt"
import "time"

//Refer to https://www.studygolang.com/articles/11263
//Question: 1. why -5,17,12, according print,should be 17? -5? 12

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Millisecond * 100)
		fmt.Println(s)
	}
}

func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	c <- sum
}
func main() {
	//	go say("word")
	//	fmt.Println("vim-go")
	//	say("hello")
	a := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int, 3)
	fmt.Printf("a=%v,a[:3]=%v,a[3:]=%v\n", a, a[:3], a[3:])
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	x, y := <-c, <-c //Get from chan c
	fmt.Println(x, y, x+y)
}
