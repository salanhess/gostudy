package google108

import (
	"fmt"
)

func PrintMap1(m map[string]string) {
	fmt.Printf("len(m)=%d m=%v\n", len(m), m)
}

func BrowseMap1(m map[string]string) {
	for key, v := range m {
		fmt.Printf("m[%s]=%s\n", key, v)
	}
}

func ScanMap1(m map[string]string, keyflag string) {
	if keyval, ok := m[keyflag]; ok {
		fmt.Printf("m[%s]=%s\n", keyflag, keyval)
	} else {
		fmt.Println("Not find ", keyflag, " in map", m)
	}
}

func Study_C3_4() {
	m := map[string]string{"name": "carol", "age": "12", "sex": "male"}
	PrintMap1(m)
	m2 := make(map[string]string)
	PrintMap1(m2)
	BrowseMap1(m)
	fmt.Println("============")
	ScanMap1(m, "name")
	ScanMap1(m, "newname")
	m["name"] = "susan"
	m["age"] = "33"
	ScanMap1(m, "name")
	ScanMap1(m, "age")
	BrowseMap1(m)
	fmt.Println("====delete map element========")
	delete(m, "name")
	BrowseMap1(m)
}
