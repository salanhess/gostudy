package google108

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

//（UTF-8 可变字节）每个中文三字节，中文一字节 3*4+6=18
//string s : convert to []rune(s)  ,range 方式遍历
func Study_C3_6() {
	var s = "Yes!我爱go语言"
	fmt.Printf("%s len is %d\n ", []byte(s), len(s))
	for _, b := range []byte(s) {
		fmt.Printf("%X ", b)
	}
	fmt.Println("")
	for i, ch := range s { //ch is rune(alias of int32)
		fmt.Printf("(Index%d:%X) ", i, ch)
	}
	fmt.Println("")
	fmt.Printf("RuneCountInString:%d", utf8.RuneCountInString(s))
	fmt.Println("")

	for i, ch := range []rune(s) {
		fmt.Printf("(Index%d:%c ", i, ch)
	}
	fmt.Println("")
	for _, sub := range strings.Fields("abc abc efg abc") {
		fmt.Println(sub)
	}
	fmt.Println("============")
	fmt.Println(strings.Trim("abcdefg", "abc"))

}
