package main

import (
	"errors"
	"fmt"
)

func echo(request string) (response string, err error) {
	if request == "" {
		err = errors.New("empty request")
		return
	}
	response = fmt.Sprintf("echo : %s", request)
	return
}

func main() {
	fmt.Println("vim-go")
	numbers2 := [...]int{1, 2, 3, 4, 5, 6}
	maxIndex2 := len(numbers2) - 1
	for i, e := range numbers2 {
		if i == maxIndex2 {
			numbers2[0] += e
		} else {
			numbers2[i+1] += e
		}
	}
	fmt.Println(numbers2)

	fmt.Println("xxxxxx slice")
	numbers3 := []int{1, 2, 3, 4, 5, 6}
	maxIndex3 := len(numbers3) - 1
	for i, e := range numbers3 {
		if i == maxIndex3 {
			numbers3[0] += e
		} else {
			numbers3[i+1] += e
		}
	}
	fmt.Println(numbers3)
	err, out := echo("")
	fmt.Println("========", err, out)
	err, out = echo("hello")
	fmt.Println("========", err, out)
}
