package main

import "fmt"

func main() {
	//Sender
	//	ch1 := make(chan int)
	ch1 := make(chan int, 3)
	go func() {
		for i := 0; i < 10; i++ {
			ch1 <- i
			fmt.Printf("Sender chan with value:%d\n", i)
		}
		fmt.Printf("Sender close channel\n")
		close(ch1)
	}()
	//Receiver
	for {
		elem, ok := <-ch1
		if !ok {
			fmt.Printf("Receiver close channel\n")
			break
		}
		fmt.Printf("Receiver get chan with value:%d\n", elem)

	}
	fmt.Println("End...")
}
