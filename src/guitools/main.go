package main

//refer to https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/03.2.md
//refer to http://wiki.jikexueyuan.com/project/the-way-to-go/15.3.html
//env in
// [root@hc172-80 delvol]# pwd
// /export/baihao/jcloud-zbs/src/jd.com/zbs/zbs-test/tools/delvol
// write file: https://gocn.io/article/40

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //获取请求的方法
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		log.Println(t.Execute(w, nil))
	} else {
		//请求的是登录数据，那么执行登录的逻辑判断
		r.ParseForm()
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
	}
}

func cleanup(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //获取请求的方法
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))
		t, _ := template.ParseFiles("cleanup.gtpl")
		log.Println(t.Execute(w, token))
	} else {
		file, err := os.OpenFile("log/operation.log", os.O_APPEND|os.O_WRONLY, 0666)
		check(err)
		defer file.Close()
		//写文件
		byteSlice := []byte(time.Now().String() + " Start to write Bytes log!\n")
		bytesWritten, err := file.Write(byteSlice)
		check(err)
		fmt.Println("Wrote %d bytes.\n", bytesWritten)
		//请求的是登录数据，那么执行登录的逻辑判断
		r.ParseForm()
		token := r.Form.Get("token")
		if token != "" {
			//验证token的合法性
			fmt.Println("Verify token wether correct...")
		} else {
			//不存在token报错
			fmt.Println("Error,token not exist")
		}
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
		//写文件
		bytesWritten, err = file.Write([]byte(strings.Join(r.Form["username"], " ") + " "))
		check(err)
		fmt.Println("Wrote %d bytes.\n", bytesWritten)
		//写文件
		bytesWritten, err = file.Write([]byte(strings.Join(r.Form["password"], " ") + " "))
		check(err)
		fmt.Println("Wrote %d bytes.\n", bytesWritten)
		if len(r.Form["email"][0]) == 0 {
			//为空的处理
			fmt.Println("email add must input!")
		} else {
			if m, _ := regexp.MatchString(`^([\w\.\_]{2,10})@jd.com$`, r.Form.Get("email")); !m {
				//if m, _ := regexp.MatchString(`^([\w\.\_]{2,10})@(\w{1,}).([a-z]{2,4})$`, r.Form.Get("email")); !m {
				fmt.Println("email addr format not corret!")
			} else {
				fmt.Println("email:", r.Form["email"])
				//写文件
				bytesWritten, err = file.Write([]byte(strings.Join(r.Form["email"], " ") + " "))
				check(err)
				fmt.Println("Wrote %d bytes.\n", bytesWritten)
			}
		}
		if len(r.Form["tenant_id"][0]) == 0 {
			//为空的处理
			fmt.Println("tenant_id  must input!")
		} else {
			fmt.Println("tenant_id:", r.Form["tenant_id"])
			//写文件
			bytesWritten, err = file.Write([]byte(strings.Join(r.Form["tenant_id"], "") + " "))
			check(err)
			fmt.Println("Wrote %d bytes.\n", bytesWritten)
		}
		fmt.Println("ZBS Operation:", r.Form["operation"])
		//写文件
		bytesWritten, err = file.Write([]byte(strings.Join(r.Form["operation"], "")))
		check(err)
		fmt.Println("Wrote %d bytes.\n", bytesWritten)
		if len(r.Form["note"][0]) == 0 {
			//为空的处理
			fmt.Println("Operation reason  must input!")
		} else {
			fmt.Println("note:", r.Form["note"])
			//写文件
			bytesWritten, err = file.Write([]byte(strings.Join(r.Form["note"], "") + "\n"))
			check(err)
			fmt.Println("Wrote %d bytes.\n", bytesWritten)
		}
	}
}

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       //解析参数，默认是不会解析的
	fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
		out, err := Shell(k)
		if err != nil {
			fmt.Fprintf(w, "Error parameter not correct!") //这个写入到w的是输出到客户端的
		}
		fmt.Fprintf(w, "Hello astaxie!") //这个写入到w的是输出到客户端的
		fmt.Fprintf(w, out)              //这个写入到w的是输出到客户端的
	}
}

func sayhi(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       //解析参数，默认是不会解析的
	fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
		fmt.Fprintf(w, "Hi astaxie!") //这个写入到w的是输出到客户端的
	}
}

func Shell(bash string) (string, error) {
	var buf bytes.Buffer
	cmd := exec.Command("sh", "-c", bash)
	cmd.Stderr = &buf
	cmd.Stdout = &buf
	err := cmd.Run()
	out := buf.String()
	return out, err
}

func main() {
	http.HandleFunc("/bash", sayhelloName)   //设置访问的路由
	http.HandleFunc("/hi", sayhi)            //设置访问的路由
	http.HandleFunc("/login", login)         //设置访问的路由
	http.HandleFunc("/cleanup", cleanup)     //设置访问的路由
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
