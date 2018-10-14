package main

import (
	"fmt"
	"unsafe"
)

type Dog struct {
	name string
	age  int
}

func (d *Dog) Setname(name string) {
	d.name = name
}

func main() {
	dog := Dog{"little dog", 11}
	dogP := &dog
	dogPtr := uintptr(unsafe.Pointer(dogP))
	fmt.Println(dogPtr)
	dogP.Setname("big dog")
	fmt.Printf("After setname, dog is %v\n", *dogP)
	namePtr := dogPtr + unsafe.Offsetof(dogP.name)
	nameP := (*string)(unsafe.Pointer(namePtr))
	fmt.Printf("namePtr is %v\n", namePtr)
	newDog := "Huge Ptr dog"
	fmt.Printf("newDog Ptr is %v\n", &newDog)
	//*nameP = "Huge Ptr dog"
	*nameP = newDog
	//nameP = &newDog // 这样操作不行，因为nameP是一个C type的传统指针，必须*用法
	fmt.Printf("After setname, dog is %v\n", *dogP)
	*nameP = "Monster"
	fmt.Printf("After setname, dog is %v\n", *dogP)
	fmt.Printf("nameP is %v\n", nameP)

}
