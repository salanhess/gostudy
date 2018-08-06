package google108

import (
	"fmt"
)

func LengthofNonRepeating_rune(s string) {
	for i, rune := range s {
		fmt.Println(i, rune)
	}

}

//leetcode:寻找最长的不重复字符串的长度（位置在start，字符起始点记录在start）
//Case1: abcda
//Case2: aabbcc
func LengthofNonRepeating(s string) {
	lastoccured := make(map[byte]int)
	start := 0
	maxlength := 0
	for i, ch := range []byte(s) {
		lastI, ok := lastoccured[ch]
		if ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxlength {
			maxlength = i - start + 1
		}
		lastoccured[ch] = i
	}
	fmt.Printf("maxlength=%d\n", maxlength)
}

func LengthofNonRepeating_Rune(s string) {
	lastoccured := make(map[rune]int)
	start := 0
	maxlength := 0
	for i, ch := range []rune(s) {
		lastI, ok := lastoccured[ch]
		if ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxlength {
			maxlength = i - start + 1
		}
		lastoccured[ch] = i
	}
	fmt.Printf("maxlength=%d\n", maxlength)
}
