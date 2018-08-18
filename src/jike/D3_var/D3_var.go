package main

import "fmt"

func main() {

	var name = guessnametype()
	fmt.Printf("name=%v\n", *name)

	var name2 = guessnametype2()
	fmt.Printf("name2=%v\n", *name2)

	var var1 = 1
	fmt.Printf("Before call,var1 is %d, address is %s\n", var1, &var1)
	reusevarnam1(var1)
	fmt.Printf("After reusevarnam1,var1 is %d, address is %s\n\n", var1, &var1)

	var var2 = 2
	fmt.Printf("Before call,var2 is %d, address is %s\n", var2, &var2)
	reusevarnam2(var2)
	fmt.Printf("After reusevarnam1,var2 is %d, address is %s\n\n", var2, &var2)

	var var3 = 3
	fmt.Printf("Before call,var3 is %d, address is %s\n", var3, &var3)
	reusevarnam3(&var3)
	fmt.Printf("After reusevarnam3,var3 is %d, address is %s\n\n", var3, &var3)

}

func reusevarnam1(var1 int) {
	var1 = 3 + var1
	fmt.Printf("Inside reusevarnam1,var1 is %d, address is %s\n", var1, &var1)
}

func reusevarnam2(var2 int) {
	for var2 := 1; var2 < 3; var2++ {
		fmt.Println("reusevarnam2 ...")
		fmt.Printf("Inside reusevarnam2,var2 is %d, address is %s\n", var2, &var2)
	}
	fmt.Println("reusevarnam2")
}

func reusevarnam3(var3 *int) {
	*var3 = *var3 + 100
	fmt.Printf("Inside reusevarnam2,var3 is %d, address is %s\n", *var3, var3)
}

func guessnametype() *string {
	var n = "abc"
	return &n
}

func guessnametype2() *int {
	var n = 12
	return &n
}
