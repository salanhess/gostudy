package mock

import "fmt"

type Retreiver struct {
	Context string
}

//modify string default expression
func (r *Retreiver) String() string {
	return fmt.Sprintf("Retriever: {Contents=%s}", r.Context)
}

func (r *Retreiver) Get(url string) string {
	return r.Context
}

func (r *Retreiver) Post(url string, form map[string]string) string {
	r.Context = form["Context"]
	return "ok"
}
