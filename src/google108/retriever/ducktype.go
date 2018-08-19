package main

import (
	"fmt"
	"google108/retriever/mock"
	"google108/retriever/real"
	"time"
)

const url = "http://www.imooc.com"

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

type RetrieverPoster interface {
	Retriever
	Poster
}

func download(poster Poster) {
	poster.Post(url, map[string]string{"name": "imooc", "couse": "golang"})
}

func session(s RetrieverPoster) string {
	//	s.Get(url)
	s.Post(url, map[string]string{"Context": "another faked imooc.com"})
	return s.Get(url)

}

func inspect(r Retriever) {
	fmt.Println("Inspecting", r)
	fmt.Printf("> %T %v\n", r, r)
	fmt.Print(" > Type switch:")
	switch v := r.(type) {
	case *mock.Retreiver:
		fmt.Println(v.Context, "[switchtype]=======")
	case *real.Retreiver:
		fmt.Println(v.TimeOut, "[switchtype]**********")
	default:
		panic("Unknown type")
	}
}

//func inspect2(r Retriever) {
//	if v, ok := r.(*mock.Retreiver); ok {
//		fmt.Println(v.Context, "[assert]=======")
//	}
//	if v, ok := r.(*real.Retreiver); ok {
//		fmt.Println(v.UserAgent, "[assert]******")
//	}
//}

func main() {
	var r Retriever
	fmt.Printf("[After init]Retriever is %T,value is %v\n", r, r)
	r = &mock.Retreiver{"I love moc"}
	inspect(r)
	//	inspect2(r)
	r = &real.Retreiver{UserAgent: "Mozilla 2.0", TimeOut: time.Minute}
	inspect(r)
	//	inspect2(r)
	fmt.Println(r.Get("http://www.chu6.top"))

	retriever := &mock.Retreiver{url}
	fmt.Println("===========", session(retriever))

	download(retriever)
	inspect(retriever)

}
