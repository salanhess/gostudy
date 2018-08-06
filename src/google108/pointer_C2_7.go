package google108

import (
	_ "fmt"
)

func Try_nopointer(a int) {
	a++
}

func Try_pointer(a *int) {
	*a += 1
}

func SwapNoEffect(a, b int) {
	a, b = b, a
}

func Swap(a, b *int) {
	*a, *b = *b, *a
}

func SwapReturn(a, b int) (int, int) {
	return b, a
}
