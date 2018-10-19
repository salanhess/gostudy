package main

//refer to https://stackoverflow.com/questions/29981050/concurrent-writing-to-a-file

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"
)

var mu sync.Mutex

func WriteToFile(text string, f *os.File) {
	mu.Lock()
	defer mu.Unlock()
	//write to the file
	fmt.Fprintf(f, "[%v] [%v]Printing out: %v\n", time.Now(), time.Now().Nanosecond(), text)
	//write to stdout
	fmt.Printf("[%v] [%v]Printing out: %v\n", time.Now(), time.Now().Nanosecond(), text)
}

func TaskSnap(snapid string, f *os.File) string {
	WriteToFile("TaskSnap "+snapid, f)
	fmt.Printf("tasksnap:%s\n", snapid)
	flag := rand.Intn(3)
	time.Sleep(time.Duration(flag) * time.Microsecond)
	return snapid
}

func DesSnap(snapid string, f *os.File) {
	//	flag := rand.Intn(3)
	//	time.Sleep(time.Duration(flag) * time.Nanosecond)
	WriteToFile("[DesSnap] "+snapid, f)
	fmt.Printf("Des snap:%s\n", snapid)
}

func main() {
	d, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	filename := d + "/log.txt"
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(f.Stat())
	defer f.Close()
	//Sender
	var snapidlist = []string{"snap1", "snap2", "snap3"}
	ch1 := make(chan string, 50)
	go func() {
		for index, snapid := range snapidlist {
			ch1 <- TaskSnap(snapid, f)
			fmt.Printf("Sender chan with index:%d snap id is %s\n", index, snapid)

		}
		fmt.Printf("Sender close channel\n")
		close(ch1)
	}()

	//Receiver
	for {
		elem, ok := <-ch1
		if !ok {
			fmt.Printf("Receiver close channel\n")
			break
		}
		DesSnap(elem, f)
		fmt.Printf("Receiver get chan with value:%v\n", elem)

	}
	fmt.Println("End...")
}
