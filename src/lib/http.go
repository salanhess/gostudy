package lib

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func PostHttpResp(req string, data []byte) (*http.Response, []byte, error) {
	// func PostHttpResp(req string, data []byte) (*http.Response, *json.RawMessage, error) {
	client := &http.Client{Timeout: 10 * time.Second}
	//log.Println("######")
	//log.Println(req, string(data))
	if resp, err := client.Post(req, "application/json",
		bytes.NewBuffer(data)); err != nil {
		return nil, nil, err
	} else {
		defer resp.Body.Close()
		if raw, err := ioutil.ReadAll(resp.Body); err != nil {
			return resp, nil, err
			// } else if err = json.Unmarshal(raw, objmap); err != nil {
			// return resp, nil, err
		} else {
			return resp, raw, nil
		}
	}
}

func PostHttpRespWithHeaders(req string, headers map[string]string, data []byte) (*http.Response, []byte, error) {
	// data := []byte(fmt.Sprintf(`{"reboot" : {"type" : "SOFT"}}`))
	client := &http.Client{Timeout: 10 * time.Second}
	log.Println("[INFO][PostHttpRespWithHeaders] url: " + req + ", post_data: " + string(data))
	if req, err := http.NewRequest("POST", req, bytes.NewReader(data)); err != nil {
		log.Println("[ERROR][PostHttpRespWithHeaders] http.NewRequest err ", err)
		return nil, nil, err
	} else {
		for k, v := range headers {
			req.Header.Add(k, string(v))
		}

		if resp, err := client.Do(req); err != nil {
			log.Println("[ERROR][PostHttpRespWithHeaders] client.Do err ", err)
			return nil, nil, err
		} else {
			defer resp.Body.Close()
			if raw, err := ioutil.ReadAll(resp.Body); err != nil {
				log.Println("[ERROR][PostHttpRespWithHeaders] ioutil.ReadAll err ", err)
				return resp, nil, err
			} else {
				log.Println("INFO][PostHttpRespWithHeaders] raw: " + string(raw))
				return resp, raw, nil
			}
		}
	}
}

func GetHttpRespWithHeaders(req string, headers map[string]string) (*http.Response, []byte, error) {
	client := &http.Client{Timeout: 10 * time.Second}
	log.Println("[INFO][GetHttpRespWithHeaders] url: ", req)
	if req, err := http.NewRequest("GET", req, nil); err != nil {
		log.Println("[ERROR][GetHttpRespWithHeaders] http.NewRequest err ", err)
		return nil, nil, err
	} else {
		for k, v := range headers {
			req.Header.Add(k, string(v))
		}

		if resp, err := client.Do(req); err != nil {
			log.Println("[ERROR][GetHttpRespWithHeaders] client.Do err ", err)
			return nil, nil, err
		} else {
			defer resp.Body.Close()
			if raw, err := ioutil.ReadAll(resp.Body); err != nil {
				return resp, nil, err
			} else {
				//log.Println("[INFO][GetHttpRespWithHeaders] raw: " + string(raw))
				return resp, raw, nil
			}
		}
	}
}

func DeleteHttpRespWithHeaders(req string, headers map[string]string) (*http.Response, []byte, error) {
	log.Println("[INFO][DeleteHttpRespWithHeaders] url: ", req)
	if req, err := http.NewRequest("DELETE", req, nil); err != nil {
		log.Println("[ERROR][DeleteHttpRespWithHeaders] http.NewRequest err ", err)
		return nil, nil, err
	} else {
		for k, v := range headers {
			req.Header.Add(k, string(v))
		}

		if resp, err := http.DefaultClient.Do(req); err != nil {
			log.Println("[ERROR][DeleteHttpRespWithHeaders] client.Do err ", err)
			return resp, nil, err
		} else {
			defer resp.Body.Close()
			if raw, err := ioutil.ReadAll(resp.Body); err != nil {
				return resp, nil, err
			} else {
				return resp, raw, nil
			}
		}
	}
}

