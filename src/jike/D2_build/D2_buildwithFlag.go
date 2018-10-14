package main

import (
	"flag"
	"fmt"
	"os"
)

var name string
var age int

func init() {
	fmt.Println("this is D2 build with flag init")
	flag.CommandLine = flag.NewFlagSet("", flag.ExitOnError) //PanicOnError also
	flag.CommandLine.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage of %s:\n", "question")
		flag.PrintDefaults()
	}
	flag.StringVar(&name, "name", "everyone", "use for name as para")
	flag.IntVar(&age, "age", 16, "use for age as para")

}

func main() {
	//	flag.Usage = func() {
	//		fmt.Fprintf(os.Stderr, "usage of %s:\n", "question")
	//		flag.PrintDefaults()
	//	}

	flag.Parse()
	fmt.Printf("Hello, %s! you are %d years old.\n", name, age)
}
