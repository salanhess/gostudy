package main

/*
有了这个偏移量，又有了结构体值在内存中的起始存储地址（这里由dogPtr变量代
表），把它们相加我们就可以得到dogP的name字段值的起始存储地址了。这个地址由变
量namePtr代表。
此后，我们可以再通过两次类型转换把namePtr的值转换成一个*string类型的值，这样
就得到了指向dogP的name字段值的指针值。
你可能会问，我直接用取址表达式&(dogP.name)不就能拿到这个指针值了吗？干嘛绕这
么大一圈呢？你可以想象一下，如果我们根本就不知道这个结构体类型是什么，也拿不到
dogP这个变量，那么还能去访问它的name字段吗？
答案是，只要有namePtr就可以。它就是一个无符号整数，但同时也是一个指向了程序内
部数据的内存地址。它可能会给我们带来一些好处，比如可以直接修改埋藏得很深的内部
数据
*/
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
	fmt.Printf("[nameP] is %v\n", nameP)
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
