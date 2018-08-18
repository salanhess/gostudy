package main

import (
	"fmt"
	"google108/retriever/mock"
	"google108/retriever/real"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("http://www.imooc.com")
}

func main() {
	var r Retriever
	r = mock.Retriever{"this is fake imooc.com"}
	fmt.Println(download(r))
	//fmt.Println(download(mock.Retriever{"this is fake imooc.com"}))

	r = real.Retriever{}
	fmt.Println(download(r))
}
