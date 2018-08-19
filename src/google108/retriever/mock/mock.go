package mock

type Retreiver struct {
	Context string
}

func (r Retreiver) Get(url string) string {
	return r.Context
}
