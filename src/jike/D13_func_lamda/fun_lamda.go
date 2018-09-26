package main

import (
	"errors"
	"fmt"
)

func simplelamda() {
	var j int = 5
	a := func() func() {
		var i int = 10
		return func() {
			fmt.Printf("i,j:%d,%d\n", i, j)
		}
	}()
	fmt.Println("lamda func can use i always and use j dynamic")
	a()
	j *= 2
	a()
}

type operate func(x, y int) int

/*
接受其他的函数作为参数传入:
通过编写calculate函数来实现两个整数间的加减乘除运算，
但是希望两个整数和具体的操作都由该函数的调用方给出，那么，这样一个函数应该怎样编写:
*/

func calculate(x, y int, op operate) (int, error) {
	if op == nil {
		return 0, errors.New("invalid operation")
	}
	return op(x, y), nil
}

/*
另一个特点，把其他的函数作为结果返回，又是怎么玩的呢？你可以看看我在
demo27.go 文件中声明的函数类型calculateFunc和函数genCalculator。其中，
genCalculator函数的唯一结果的类型就是calculateFunc
*/
type calculateFunc func(x int, y int) (int, error)

func genCalculator(op operate) calculateFunc {
	return func(x, y int) (int, error) {
		if op == nil {
			return 0, errors.New("invalid operation")
		}
		return op(x, y), nil
	}
}

func main() {
	simplelamda()
	var x, y = 3, 4
	opadd := func(x, y int) int {
		fmt.Printf("[op] get x:%d,y:%d\n", x, y)
		return x + y
	}

	opminus := func(x, y int) int {
		fmt.Printf("[Subtraction-minus] get x:%d,y:%d\n", x, y)
		return x - y
	}

	result, err := calculate(x, y, opadd)
	fmt.Printf("===result is:%d (error %v)\n", result, err)
	result, err = calculate(x, y, opminus)
	fmt.Printf("Subtraction-minus ===result is:%d (error %v)\n", result, err)
	x, y = 56, 78
	add := genCalculator(opadd)
	result, err = add(x, y)
	fmt.Printf("The result: %d (error: %v)\n", result, err)
	minus := genCalculator(opminus)
	result, err = minus(x, y)
	fmt.Printf("Subtraction-minus:The result: %d (error: %v)\n", result, err)
}
