package main

import "fmt"

func int16toint8(v int16) {
	var srcInt = int16(v)
	dstInt := int8(srcInt)
	fmt.Printf("srt type is %T,value is %d\ndst type is %T,value is %d\n", srcInt, srcInt, dstInt, dstInt)
	fmt.Println("==============int16toint8====================")
}

func inttostring(v int8) {
	fmt.Printf("srt type is %T,value is %d\ndst type is %T,value is %s\n", v, v, string(v), string(v))
	fmt.Println("==============int2String====================")
}

func slicebyte2sting_fail(v []byte) {
	fmt.Printf("srt type is %T,value is %v\ndst type is %T,value is %v\n", v, v, string(v), string(v))
	fmt.Println("==============slicebyte2sting====================")
}

func slicebyte2sting(v []byte) {
	fmt.Printf("srt type is %T,value is %v\ndst type is %T,value is %v\n", v, v, string(v[0]+v[1]+v[2]), string(v))
	fmt.Println("==============slicebyte2sting====================")
}

func slicerune2sting(v []rune) {
	fmt.Printf("srt type is %T,value is %v\ndst type is %T,value is %v\n", v, v, string(v), string(v))
	fmt.Println("==============slicerune2sting====================")
}
func main() {
	fmt.Println("vim-go")
	int16toint8(-255)
	inttostring(-1)
	inttostring(32)
	slicebyte2sting([]byte{'\xe4', '\xdb', '\xa0', '\xe5', '\xa5', '\xdb'}) //你好
	fmt.Printf("Utf-8 3 byte join together can display Chinese character %s %s \n", string("\xe4\xdb\xa0"), string("\xe5\xa5\xdb"))
	//	slicebyte2sting([]byte{'\xe4\xdb\xa0', '\xe5\xa5\xdb'}) //你好
	slicerune2sting([]rune{'\u4F60', '\u597D'}) //你好
	fmt.Println("len(你)=", len("你"))
}
