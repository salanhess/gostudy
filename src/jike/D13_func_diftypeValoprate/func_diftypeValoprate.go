package main

import "fmt"

/*
答案是：原数组不会改变。为什么呢？原因是，所有传给函数的参数值都会被复制，函数
在其内部使用的并不是参数值的原值，而是它的副本。
由于数组是值类型，所以每一次复制都会拷贝它，以及它的所有元素值。我在modify函
数中修改的只是原数组的副本而已，并不会对原数组造成任何影响。
注意，对于引用类型，比如：切片、字典、通道，像上面那样复制它们的值，只会拷贝它
们本身而已，并不会拷贝它们引用的底层数据。也就是说，这时只是浅表复制，而不是深
层复制。
以切片值为例，如此复制的时候，只是拷贝了它指向底层数组中某一个元素的指针，以及
它的长度值和容量值，而它的底层数组并不会被拷贝。
*/
func main() {
	fmt.Println("vim-go")
	//array will NOT change original val
	array1 := [3]string{"a", "b", "c"}
	array2 := modifyarr(array1)
	fmt.Printf("array1 = %v\n", array1)
	fmt.Printf("array2 = %v\n", array2)
	//slice will change original val
	slice1 := []string{"a", "b", "c"}
	slice2 := modifyslice(slice1)
	fmt.Printf("slice1= %v\n", slice1)
	fmt.Printf("slice2= %v\n", slice2)
	//complex typewill change original val
	comp1 := [3][]string{
		[]string{"e", "c", "f"},
		[]string{"g", "h", "i"},
		[]string{"j", "k", "l"},
	}
	comp2 := modifycomplextype(comp1)
	fmt.Printf("comp1= %v\n", comp1)
	fmt.Printf("comp2= %v\n", comp2)
}

func modifyarr(a [3]string) [3]string {
	a[1] = "x"
	return a
}

func modifyslice(s []string) []string {
	s[1] = "x"
	return s
}

func modifycomplextype(comp [3][]string) [3][]string {
	comp[1][1] = "X"
	return comp
}
