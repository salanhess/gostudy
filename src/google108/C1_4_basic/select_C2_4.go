package google108

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"reflect"
	"runtime"
)

func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		return a / b, nil
	default:
		return 0, fmt.Errorf("Unsupported operation: " + op)
	}
}

func readfile(filename string) {
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println("[Error]", "cannot print file contents", err)
	} else {
		fmt.Println(string(contents))
	}
}

func printfile(filename string) {
	if file, err := os.Open(filename); err != nil {
		panic(err)
	} else {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}
}

func forever(str string) {
	for {
		fmt.Println(str)
	}
}

func div(a, b int) (q, r int) {
	q = a / b
	r = a % b
	return
}

func grade(score int) string {
	g := ""
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf(
			"Wrong score: %d", score))
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	}
	return g
}

//rewrite func for above operation
func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with parameters: "+
		"(%d, %d)\n", opName, a, b)
	return op(a, b)
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func sum(numbers ...int) {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	fmt.Println("=====", s)

}
