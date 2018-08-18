package queue

type Queue []int

func (q *Queue) Push(v int) {
	*q = append(*q, v)
}

func (q *Queue) Pop() int {
	if len(*q) == 0 {
		panic("error ,q is nil!")
	}
	head := (*q)[0]
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
