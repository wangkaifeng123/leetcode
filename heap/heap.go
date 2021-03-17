package heap

type Heap struct {
	Elements []int
}

//堆的特性
//1.堆节点是i  父节点是i/2
//2.堆节点是i  子节电是 2i 和 2i+1
func NewHeap() *Heap {
	var h Heap
	h.Elements = make([]int, 0)
	return &h
}

func (h *Heap) Insert(ele int) {
	h.Elements = append(h.Elements, ele)
	i := len(h.Elements) - 1
	for h.Elements[i/2] < ele {
		h.Elements[i/2], h.Elements[i] = h.Elements[i], h.Elements[i/2]
		i = i / 2
	}
	return
}

//堆删除最大值
func (h *Heap) RemoveMax() {
	if len(h.Elements) == 0 {
		return
	}
	h.Elements[0] = h.Elements[len(h.Elements)-1]
	h.Elements = h.Elements[0 : len(h.Elements)-1]
	i := 0
	for 2*i < len(h.Elements)-1 {
		//和左右节点比较
		temp := i
		i = swap(h.Elements, i, 2*i, 2*i+1)
		if temp == i {
			return
		}
	}
	return
}

func swap(elements []int, i, j, k int) int {
	if elements[i] >= elements[j] && elements[i] >= elements[k] {
		return i
	}
	if elements[i] < elements[j] && elements[j] >= elements[k] {
		elements[j], elements[i] = elements[i], elements[j]
		return j
	}
	if elements[k] >= elements[i] && elements[k] >= elements[j] {
		elements[k], elements[i] = elements[i], elements[k]
		return k
	}
	return 0
}
