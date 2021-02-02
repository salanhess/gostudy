package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"syscall"
)

var (
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

func init() {
	errFile, err := os.OpenFile("errors.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("打开日志文件失败：", err)
	}

	Info = log.New(os.Stdout, "Info:", log.Ldate|log.Ltime|log.Lshortfile)
	Warning = log.New(os.Stdout, "Warning:", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(io.MultiWriter(os.Stderr, errFile), "Error:", log.Ldate|log.Ltime|log.Lshortfile)

}

func main() {
	fmt.Println("vim-go")
	threshold := 10
	path := "/tmp/"
	isFull, err := IsPathUsageFull(path, threshold, Info)
	fmt.Printf("isFull:%v, err:%v \n", isFull, err)
	volDirs, _ := ioutil.ReadDir(path)
	fmt.Printf("varDirs:%#v\n", volDirs)
	for _, vol := range volDirs {
		fmt.Printf("volname %v Size %v\n", (vol.Name), (vol.Size))
	}
	err = clearOldVolumeInfo(path, volDirs, Info)
	fmt.Printf("err:%#v\n", err)
}

func IsPathUsageFull(path string, threshold int, log *log.Logger) (bool, error) {
	//check the path is ready by the cache-ssd.
	d, err := DiskUsage(path, log)
	if err != nil {
		return false, err
	}
	fmt.Printf("float64(d.Used*100/d.All) %v, float64(threshold) %v \n", float64(d.Used*100/d.All), float64(threshold))
	if float64(d.Used*100/d.All) > float64(threshold) {
		return true, nil
	}
	return false, nil
}

type DiskStatus struct {
	All  uint64 `json:"all"`
	Used uint64 `json:"used"`
	Free uint64 `json:"free"`
}

func DiskUsage(path string, log *log.Logger) (*DiskStatus, error) {
	fs := syscall.Statfs_t{}
	err := syscall.Statfs(path, &fs)
	if err != nil {
		Error.Println("DiskUsage Failed, fs:", path, "err:", err)
		return nil, err
	}
	d := &DiskStatus{}
	d.All = fs.Blocks * uint64(fs.Bsize)
	d.Free = fs.Bfree * uint64(fs.Bsize)
	d.Used = d.All - d.Free
	return d, nil
}

func clearOldVolumeInfo(path string, vols []os.FileInfo, log *log.Logger) error {
	var oldVolModifyTime int64
	var oldVol string
	for _, vol := range vols {
		if oldVolModifyTime == 0 {
			oldVolModifyTime = vol.ModTime().Unix()
			oldVol = vol.Name()
			continue
		}
		if vol.ModTime().Unix() < oldVolModifyTime {
			oldVolModifyTime = vol.ModTime().Unix()
			oldVol = vol.Name()
		}
	}

	if oldVol != "" {
		err := os.RemoveAll(path + "/" + oldVol)
		if err != nil {
			Error.Println("os.Remove, vol:", oldVol, "err:", err)
		}
	}

	return nil
}
