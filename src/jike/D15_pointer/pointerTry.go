package main

import "fmt"

/*
今天的问题是：你能列举出 Go 语言中的哪些值是不可寻址的吗？
这道题的典型回答是以下列表中的值都是不可寻址的。Case:
1.常量的值。
2.基本类型值的字面量。
3.算术操作的结果值。
4.对各种字面量的索引表达式和切片表达式的结果值。不过有一个例外，对切片字面量的索引结果值却是可寻址的。
5.对字符串变量的索引表达式和切片表达式的结果值。
6.对字典变量的索引表达式的结果值。
7.函数字面量和方法字面量，以及对它们的调用表达式的结果值。
8.结构体字面量的字段值，也就是对结构体字面量的选择表达式的结果值。
9.类型转换表达式的结果值。
10.类型断言表达式的结果值。
11.接收表达式的结果值。
*/

func main() {
	fmt.Println("vim-go")
	const Mycst = 1
	//Case1
	//fmt.Printf("const Mycst is %v, address is %d\n", Mycst, &Mycst)
	fmt.Printf("const Mycst is %v\n", Mycst)
	//Case2
	fmt.Printf("123 is %v\n", 123)
	//fmt.Printf("123 is %v, address is %d\n", 123, &123)
	//Case3
	fmt.Printf("1+2 is %v\n", 1+2)
	//fmt.Printf("1+2 is %v, address is %d\n", 1+2, &(1 + 2))
	//Case4
	Mystr := "abc"
	fmt.Printf("Mystr[0] is %v\n", Mystr[0])
	//fmt.Printf("Mystr[0] is %v, address is %d\n", Mystr[0], &Mystr[0])
	fmt.Printf("Mystr[0:1] is %v\n", Mystr[0:1])
	//fmt.Printf("Mystr[0:1] is %v, address is %d\n", Mystr[0:1], &Mystr[0:1])
	//Case4-Exception
	Myslice := []string{"abc", "efg"}
	fmt.Printf("Myslice[0:1] is %v\n", Myslice[0:1])
	//fmt.Printf("Myslice[0:1] is %v, address is %d\n", Myslice[0:1], &Myslice[0:1])
	v := Myslice[0:1]
	fmt.Printf("Myslice[0:1] is %v, address is %s\n", Myslice[0:1], &v)
	//Case6
	Mymap := map[int]string{1: "malaixiya", 2: "singapore"}
	//	fmt.Printf("Mymap is %v, address is %d\n", Mymap[1], &Mymap[1])
	fmt.Printf("Mymap is %v\n", Mymap[1])
	//Case7
	Myfunc := func() {
		fmt.Println("My lamda func!")
	}
	Myfunc()
	fmt.Printf("Lamda func is %v\n", &Myfunc)
	//fmt.Printf("Myfunc is %v\n", &func() {
	//	fmt.Println("Hello world")
	//})
	//Case8
	type Mystruct struct {
		Smystr string
		Smyint int
	}
	mystruct := Mystruct{"Carol", 22}
	//mystruct := Mystruct{Smystr: "Carol", Smyint: 22}
	fmt.Printf("mystruct is %v,address is %s\n", mystruct, &mystruct.Smyint)
	fmt.Printf("mystruct direct value is %s\n\n", Mystruct{"Susan", 33}.Smystr)
	//fmt.Printf("mystruct direct value is %s\n, address is %s\n", Mystruct{"Susan", 33}.Smystr, &Mystruct{"Susan", 33}.Smystr)
}
