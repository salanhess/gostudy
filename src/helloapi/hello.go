// restAPItest

package main

import (
	"encoding/json"
	"fmt"
<<<<<<< HEAD
)

func checkSlice(s []int) {
	fmt.Printf("s=%v,len(s)=%d,cap(s)=%d\n", s, len(s), cap(s))
}

func lenOfNonRepeat(s string) int {
	lastOccured := make(map[byte]int)
	start := 0
	maxLenth := 0

	for i, ch := range []byte(s) {
		lastI, ok := lastOccured[ch]
		if ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLenth {
			maxLenth = i - start + 1
		}
		lastOccured[ch] = i
	}
	return maxLenth
}

func stringview(s string) {
	for i, v := range s {
		fmt.Printf("Not convert,str=%s,s=%d,v=%c\n", s, i, v)
	}

	for i, v := range []byte(s) {
		fmt.Printf("Convert to []byte,str=%s,s=%d,v=%c\n", s, i, v)
	}

	for i, v := range []rune(s) {
		fmt.Printf("Convert to []rune,str=%s,s=%d,v=%c\n", s, i, v)
	}

}
func main() {
	var s1 = []int{}
	for i := 1; i < 10; i++ {
		s1 = append(s1, i)
		checkSlice(s1)
	}
	s2 := "aabbcc"
	s3 := "我爱吃肉"
	fmt.Println(lenOfNonRepeat(s2))
	fmt.Println(lenOfNonRepeat(s3))
	stringview(s2)

	stringview(s3)
=======
	"lib"
	"log"
	"model"
	"net/http"
)

type volume struct {
	ID            uint64 `json:"card_number"`
	Balance       string `json:"card_balance"`
	Balance_time  string `json:"balance_time"`
	Validity_time string `json:"card_validity"`
	Server_time   string `json:"current_time"`
}

const openapiUrlFmt = "http://%s/v1/regions/%s"

func main() {
	fmt.Println("Hello World!")
	ip_port := "192.168.180.116:9200"
	region_disk := "cn-north-1/disks/vol-5v82jakkj3"
	url := fmt.Sprintf(openapiUrlFmt, ip_port, region_disk)
	var headers = make(map[string]string)
	headers["x-jcloud-pin"] = "SmFuZUNsb3VkMDA3"
	resp := &http.Response{}
	var raw []byte
	resp, raw, _ = lib.GetHttpRespWithHeaders(url, headers)
	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	log.Println("=====\n", string(raw))

	if resp.Status == "200 OK" {
		volumedes := &model.DescVolume{}
		_ = json.Unmarshal(raw, volumedes) // JSON to Struct
		log.Println("=====\n", volumedes.RequestId)
		log.Println("=====\n", volumedes.Result.Id)
	} else {
		if resp.Status == "505 HTTP Version Not Supported" {
			Errdes := &model.DescVolumeErr{}
			_ = json.Unmarshal(raw, Errdes) // JSON to Struct
			log.Println("*****\n", Errdes.RequestId)
			log.Println("*****\n", Errdes.Error.Message)
		} else {
			fmt.Println("unknonw handler")
		}
	}
>>>>>>> 172c69058f800f09329c4d3e52b834521eaa1bdc

}
