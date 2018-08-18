package google108

import (
	"fmt"
)

func Printslicestruct(slice []int) {
	fmt.Printf("len(slice)=%d cap(slice)=%d slice=%v\n", len(slice), cap(slice), slice)
}

func Printsliceindex(slice []int) {
	for i, v := range slice {
		fmt.Printf("slice[%d]=%d ", i, v)
	}
	fmt.Println("")
	fmt.Println("slice[:2]=", slice[:2], "slice[3:]=", slice[3:])
}

func Try_sliceDelete() {
	fmt.Println("delete slice element with slice+append method")
	var slice = make([]int, 5, 10)
	slice = []int{1, 2, 3, 4, 5, 6}
	Printsliceindex(slice)
	Printslicestruct(slice)
	slice = append(slice[:2], slice[3:]...)
	Printslicestruct(slice)
	fmt.Println("pop element from front")
	slice = slice[1:]
	Printslicestruct(slice)
	fmt.Println("pop element from end")
	slice = slice[:len(slice)-1]
	Printslicestruct(slice)
}

func Try_sliceMake() {
	fmt.Println("Make slice with length,cap;copy slice(dst,src) ")
	var slice = make([]int, 2, 4)
	slice = []int{2, 4, 6, 8}
	Printslicestruct(slice)
	var s2 = make([]int, 8, 8)
	Printslicestruct(s2)
	copy(s2, slice)
	Printslicestruct(s2)
}

func Try_sliceLoop() {
	var slice []int
	for i := 0; i < 100; i++ {
		slice = append(slice, i*2+1)
		fmt.Printf("len(slice)=%d cap(slice)=%d\n", len(slice), cap(slice))
	}
	fmt.Println(slice)
}

func Try_sliceCap() {
	//	s2 = s2[:6]
	arr := [...]int{0, 1, 2, 3, 4, 5, 6}
	s1 := arr[2:6]
	fmt.Println("roll bak")
	fmt.Println("s1=:", s1, "cap(s1):", cap(s1), "len(s1):", len(s1))
	s2 := s1[3:5]
	fmt.Println("s2=:", s2, "cap(s2):", cap(s2), "len(s2):", len(s2))

	arr2 := [...]int{1, 2, 3}
	fmt.Println("arr2=", arr2, "cap(arr2)=", cap(arr2))
	s4 := append(arr2[:2], 4)
	fmt.Println("arr2=", arr2, "cap(arr2)=", cap(arr2), "cap(s4)=", cap(s4))
	s5 := append(s4[:], 5)
	fmt.Println("arr2=", arr2, "cap(s5)=", cap(s5))
	fmt.Println(s5)
}

func study_C3_3() {
	//	g.Try_sliceCap()
	//	g.Try_sliceLoop()
	//	g.Try_sliceMake()
	//	g.Try_sliceDelete()
}
