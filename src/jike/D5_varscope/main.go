package main

import "fmt"
import . "jike/D5_varscope/varpkg"

var block = "package"
var container = "xxx"

//var container = []string{"zero", "one", "two"}

func main() {
	var block = 1
	var container = map[int]string{0: "zero", 1: "one", 2: "two"}
	//	Name = "Susan"
	//block := 1
	//	block := "function"
	{
		block := "inner"
		fmt.Printf("The block is %s.\n", block)
	}
	fmt.Printf("The block is %s.\n", block)
	fmt.Printf("The element is %q.\n", container[1])
	fmt.Printf("The block is %s.\n", block)
	fmt.Printf("The Name is %s.\n", Name)

}
