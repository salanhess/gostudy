package main

import "fmt"

type Printer func(content string) (n int, err error)

func printToStd(content string) (bytesNum int, err error) {
	return fmt.Println(content)
}

func notFixedpara(args ...int) {
	for i, v := range args {
		fmt.Printf("Index[%d]=%d\n", i, v)
	}
}

func main() {
	var p Printer
	p = printToStd
	p("something")
	notFixedpara(1, 2, 3)
	notFixedpara(4, 5, 6, 7, 8)
}
