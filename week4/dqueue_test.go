package week4

import "testing"

func TestQueue(t *testing.T) {
	queue := Constructor()
	queue.Push_back(15)
	queue.Push_back(9)
	t.Log(queue.Max_value())
}

type MaxQueue struct {
	deque []int //双端队列 递减
	elems []int
}

func Constructor() MaxQueue {
	return MaxQueue{
		deque: make([]int, 0),
		elems: make([]int, 0),
	}
}

func (this *MaxQueue) Max_value() int {
	if len(this.elems) == 0 {
		return -1
	}
	return this.deque[0]
}

func (this *MaxQueue) Push_back(value int) {
	//插入元素
	this.elems = append(this.elems, value)
	//更新双端队列
	var point int
	for k, v := range this.deque {
		if v < value {
			point = k
			this.deque = this.deque[:point]
			break
		}
	}
	this.deque = append(this.deque, value)
	return
}

func (this *MaxQueue) Pop_front() int {
	if len(this.elems) == 0 {
		return -1
	}
	var res int
	if this.deque[0] == this.elems[0] {
		this.deque = this.deque[1:]
	}
	res = this.elems[0]
	this.elems = this.elems[1:]
	return res
}
