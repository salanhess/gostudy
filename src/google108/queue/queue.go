package queue

type Queue []interface{}

func (q *Queue) Push(v interface{}) {
	*q = append(*q, v)
}

func (q *Queue) Pop() interface{} {
	if len(*q) == 0 {
		panic("error ,q is nil!")
	}
	head := (*q)[0].(string) //.(int) will convert to int type
	*q = (*q)[1:]
	return head
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
	//	if len(*q) == 0 {
	//		return true
	//	} else {
	//		return false
	//	}
}
