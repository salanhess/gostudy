// restAPItest
package main

import (
	"encoding/json"
	"fmt"
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

}
