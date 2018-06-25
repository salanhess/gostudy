// restAPItest

package main

import (
	"fmt"
)

func checkSlice(s []int) {
	fmt.Printf("s=%v,len(s)=%d,cap(s)=%d\n", s, len(s), cap(s))
}

func lenOfNonRepeat(s string) int {
	lastOccured := make(map[byte]int)
	start := 0
	maxLenth := 0

	for i, ch := range []byte(s) {
		lastI, ok := lastOccured[ch]
		if ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLenth {
			maxLenth = i - start + 1
		}
		lastOccured[ch] = i
	}
	return maxLenth
}

func stringview(s string) {
	for i, v := range s {
		fmt.Printf("Not convert,str=%s,s=%d,v=%c\n", s, i, v)
	}

	for i, v := range []byte(s) {
		fmt.Printf("Convert to []byte,str=%s,s=%d,v=%c\n", s, i, v)
	}

	for i, v := range []rune(s) {
		fmt.Printf("Convert to []rune,str=%s,s=%d,v=%c\n", s, i, v)
	}

}
func main() {
	var s1 = []int{}
	for i := 1; i < 10; i++ {
		s1 = append(s1, i)
		checkSlice(s1)
	}
	s2 := "aabbcc"
	s3 := "我爱吃肉"
	fmt.Println(lenOfNonRepeat(s2))
	fmt.Println(lenOfNonRepeat(s3))
	stringview(s2)

	stringview(s3)

}
