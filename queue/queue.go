package queue

// 定义队列
type Queue []interface{}

// 添加元素的方法
func (q *Queue) Push(v interface{}) {
	*q = append(*q, v.(int)) // 会出现运行时错误
}

// 弹出元素方法
func (q *Queue) Pop() interface{} {
	head := (*q)[0]
	*q = (*q)[1:]
	return head.(int)
}

// 空队列
func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
