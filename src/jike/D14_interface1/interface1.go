package main

import "fmt"

type Pet interface {
	SetName(name string)
	Name() string
	Category() string
}

type Cat struct {
	name string
}

type Dog struct {
	name string
	ctg  string
}

func (cat *Cat) SetName(name string) {
	cat.name = name
}

func (cat Cat) Name() string {
	return cat.name
}

func (cat Cat) Category() string {
	return "Cat"
}

func (dog *Dog) SetName(name string) {
	dog.name = "Dog: " + name
}

func (dog Dog) Name() string {
	return dog.name
}

func (dog Dog) Category() string {
	return "Dog XXX"
}

func main() {
	fmt.Println("vim-go")
	litteCat := Cat{name: "miaomiao"}
	/*
		首先，由于dog的SetName方法是指针方法，所以该方法持有的接收者就是指向dog的指针
		值的副本，因而其中对接收者的name字段的设置就是对变量dog的改动。那么当
		dog.SetName("monster")执行之后， dog的name字段的值就一定是"monster"。如果你理
		解到了这一层，那么请小心前方的陷阱。
		为什么dog的name字段值变了，而pet的却没有呢？这里有一条通用的规则需要你知晓：
		如果我们使用一个变量给另外一个变量赋值，那么真正赋给后者的，并不是前者持有的那
		个值，而是该值的一个副本。
	*/
	pet := litteCat
	if _, ok := interface{}(litteCat).(Pet); ok {
		fmt.Println("litteCat inplement interface Pet")
	}
	if _, ok := interface{}(&litteCat).(Pet); ok {
		fmt.Println("*litteCat inplement interface Pet")
	}
	fmt.Println(litteCat.Name())
	fmt.Printf("pet is %v  pet.Name %s pet.Category %s\n", pet, pet.Name(), pet.Category())
	litteCat.SetName("xxNewcat")
	fmt.Printf("litteCat is %v\n", litteCat)
	fmt.Printf("[After litteCat SetName]pet is %v  pet.Name %s pet.Category %s\n", pet, pet.Name(), pet.Category())
	/*
		当我们给一个接口变量赋值的时候，该变量的动态类型会与它的动态值一起被存储在一个
		专用的数据结构中。
		严格来讲，这样一个变量的值其实是这个专用数据结构的一个实例，而不是我们赋给该变
		量的那个实际的值。所以我才说， pet的值与dog的值肯定是不同的，无论是从它们存储
		的内容，还是存储的结构上来看都是如此。不过，我们可以认为，这时pet的值中包含了
		dog值的副本。我们就把这个专用的数据结构叫做iface吧，在 Go 语言的runtime包中它其实就叫这个
		名字。
		iface的实例会包含两个指针，一个是指向类型信息的指针，另一个是指向动态值的指
		针。这里的类型信息是由另一个专用数据结构的实例承载的，其中包含了动态值的类型，
		以及使它实现了接口的方法和调用它们的途径，等等。
		总之，接口变量被赋予动态值的时候，存储的是包含了这个动态值的副本的一个结构更加
		复杂的值。你明白了吗？
	*/
	bigDog := Dog{name: "manxiong"}
	pet2 := &bigDog
	/*
		我声明的类型Dog附带了 3 个方法。其中有 2 个值方法，分别是Name和Category，另外
		还有一个指针方法SetName。
		这就意味着， Dog类型本身的方法集合中只包含了 2 个方法，也就是所有的值方法。而它
		的指针类型*Dog方法集合却包含了 3 个方法，
		也就是说，它拥有Dog类型附带的所有值方法和指针方法。又由于这 3 个方法恰恰分别是
		Pet接口中某个方法的实现，所以*Dog类型就成为了Pet接口的实现类型。
	*/
	fmt.Printf("[Type]pet2 type is %T", pet2)
	fmt.Printf("pet2 is %v  pet2.Name %s pet2.Category %s\n", pet2, pet2.Name(), pet2.Category())
	fmt.Println(bigDog.Name())
	bigDog.SetName("xxNewDog")
	fmt.Printf("[After bigDog SetName]pet2 is %v  pet2.Name %s pet2.Category %s\n", pet2, pet2.Name(), pet2.Category())
}

/*
[root@izhp3c4nnfdk1zyhek9f3ez D14_interface1]# go run interface1.go
vim-go
*litteCat inplement interface Pet
miaomiao
pet is {miaomiao}  pet.Name miaomiao pet.Category Cat
litteCat is {xxNewcat}
[After litteCat SetName]pet is {miaomiao}  pet.Name miaomiao pet.Category Cat
[Type]pet2 type is *main.Dogpet2 is &{manxiong }  pet2.Name manxiong pet2.Category Dog XXX
manxiong
[After bigDog SetName]pet2 is &{Dog: xxNewDog }  pet2.Name Dog: xxNewDog pet2.Category Dog XXX
*/
