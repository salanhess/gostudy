package real

import (
	"net/http"
	"net/http/httputil"
	"time"
)

type Retreiver struct {
	UserAgent string
	TimeOut   time.Duration
}

func (r Retreiver) Get(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	result, err := httputil.DumpResponse(resp, false) //true

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	return string(result)
}
