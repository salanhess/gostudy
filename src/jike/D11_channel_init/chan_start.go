package main

import "fmt"

func basic_chan() {
	var ch1 = make(chan int, 2)
	ch1 <- 1
	ch1 <- 2
	elem := <-ch1
	fmt.Printf("basic_chan: elem %v \n", elem)
	elem = <-ch1
	fmt.Printf("basic_chan: elem %v \n", elem)
}

func invaild_chan1() {
	var ch1 = make(chan int, 1)
	ch1 <- 1
	//	ch1 <- 2
	elem := <-ch1
	fmt.Printf("elem %v \n", elem)
}

func invaild_chan2() {
	var ch1 = make(chan int, 1)
	ch1 <- 1
	//	_ = <-ch1
	elem := <-ch1
	fmt.Printf("elem %v \n", elem)
}

func invaild_chan3() {
	var ch1 chan int
	//	ch1 <- 1
	fmt.Println(ch1)
}

func main() {
	basic_chan()
	invaild_chan1()
	invaild_chan2()
	invaild_chan3()
}
