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
	response = fmt.Sprintf("echo: %s", request)
	return
}

func caller() {
	fmt.Println("Enter function caller.")
	panic(errors.New("something wrong")) // 正例。
	panic(fmt.Println)                   // 反例。
	fmt.Println("Exit function caller.")
}
func main() {
	// 示例1。
	for _, req := range []string{"", "hello!"} {
		fmt.Printf("request: %s\n", req)
		resp, err := echo(req)
		if err != nil {
			fmt.Printf("error: %s\n", err)
			continue
		}
		fmt.Printf("response: %s\n", resp)
	}
	fmt.Println()

	// 示例2。
	err1 := fmt.Errorf("invalid contents: %s", "#$%")
	err2 := errors.New(fmt.Sprintf("invalid contents: %s", "#$%"))
	if err1.Error() == err2.Error() {
		fmt.Println("The error messages in err1 and err2 are the same.")
	}
	s1 := []int{0, 1, 2, 3, 4}
	e5 := s1[5]
	_ = e5
}
