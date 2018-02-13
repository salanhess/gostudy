package main

//refer to https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/03.2.md
//refer to http://wiki.jikexueyuan.com/project/the-way-to-go/15.3.html
//Error handler refer to http://rgyq.blog.163.com/blog/static/3161253820139303344753/
//env in
// [root@hc172-80 delvol]# pwd
// /export/baihao/jcloud-zbs/src/jd.com/zbs/zbs-test/tools/delvol
// write file: https://gocn.io/article/40
//todo: add email and tenant_id white list via json?

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Tenant struct {
	Tenant_id string
}

type Mail struct {
	Email string
}

type Tidslice struct {
	Tenants []Tenant
	Emails  []Mail
}

func validateroles(cfg string, checkField string, target string) bool {
	var s Tidslice
	var pwd, _ = os.Getwd()
	var CfgPath string = pwd + string(os.PathSeparator) + cfg
	fmt.Println(CfgPath)
	fmt.Println("target string is: ", target)
	raw, err := ioutil.ReadFile(CfgPath)
	check(err)
	err = json.Unmarshal(raw, &s)
	check(err)
	switch {
	case checkField == "email":
		for i, v := range s.Emails {
			fmt.Println(checkField+"[", i, "] =", v)
			if target == v.Email {
				return true
			}
		}
		return false
	case checkField == "tid":
		for i, v := range s.Tenants {
			fmt.Println(checkField+"[", i, "] =", v)
			if target == v.Tenant_id {
				return true
			}
		}
		return false
	default:
		fmt.Println("not check")
		return false
	}

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

//out put err in ser console and web
func errhandle(w http.ResponseWriter, out string) {
	fmt.Println(out)
	fmt.Fprintf(w, out)
	time.Sleep(10 * time.Microsecond)
}

func wfile(file *os.File, str string) {
	//写文件
	byteSlice := []byte(str + " ")
	bytesWritten, err := file.Write(byteSlice)
	check(err)
	fmt.Println("Wrote %d bytes.\n", bytesWritten)
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
		wfile(file, time.Now().String()+" Start to write Bytes log!\n")
		//请求的是登录数据，那么执行登录的逻辑判断
		r.ParseForm()
		token := r.Form.Get("token")
		if token != "" {
			//验证token的合法性
			fmt.Println("Verify token wether correct...")
			wfile(file, "token: "+token+" ")
		} else {
			//不存在token报错
			fmt.Println("Error,token not exist")
			http.Error(w, err.Error(), 500)
		}
		//		fmt.Println("username:", r.Form["username"])
		//		fmt.Println("password:", r.Form["password"])
		//		wfile(file, strings.Join(r.Form["username"], " "))
		//		wfile(file, strings.Join(r.Form["password"], " "))
		if r.Form["email"][0] == "" {
			fmt.Println("email add must input!")
		} else {
			if m, _ := regexp.MatchString(`^([\w\.\_]{2,10})@jd.com$`, r.Form.Get("email")); !m {
				//if m, _ := regexp.MatchString(`^([\w\.\_]{2,10})@(\w{1,}).([a-z]{2,4})$`, r.Form.Get("email")); !m {
				errhandle(w, "email addr format not corret!")
				return
			} else {
				if validateroles("cfg.json", "email", r.Form["email"][0]) {
					fmt.Println("email:", r.Form["email"])
					wfile(file, strings.Join(r.Form["email"], " "))
				} else {
					errhandle(w, "email addr not in white list!")
					return
				}
			}
		}

		if r.Form["tenant_id"][0] == "" {
			errhandle(w, "tenant_id  must input!")
			return
		} else {
			if validateroles("cfg.json", "tid", r.Form["tenant_id"][0]) {
				fmt.Println("tenant_id:", r.Form["tenant_id"])
				wfile(file, strings.Join(r.Form["tenant_id"], " "))
			} else {
				errhandle(w, "tenant_id not in white list!")
				return
			}
		}
		if r.Form["note"][0] == "" {
			errhandle(w, "Operation reason  must input!")
			return
		} else {
			fmt.Println("note:", r.Form["note"])
			wfile(file, strings.Join(r.Form["note"], ""))
		}
		fmt.Println("ZBS Operation:", r.Form["operation"])
		if r.Form["operation"][0] == "delvolume" {
			fmt.Println("======ZBS Operation delvolume:")
			out, err := Shell("sh del_vol.sh " + r.Form["tenant_id"][0])
			if err != nil {
				fmt.Fprintf(w, "Error parameter not correct!") //这个写入到w的是输出到客户端的
			}
			wfile(file, "ZBS Operation:"+"sh del_vol.sh "+r.Form["tenant_id"][0])
			fmt.Fprintf(w, "ZBS Operation output: "+out) //这个写入到w的是输出到客户端的
		}
		wfile(file, strings.Join(r.Form["operation"], "")+"\n")
		fmt.Fprintf(w, "Cleanup finished...!")                 //这个写入到w的是输出到客户端的
		fmt.Fprintf(w, "Welcome to ZBS disk cleanup process!") //这个写入到w的是输出到客户端的
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
