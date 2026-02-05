package store

import "time"

type Item struct {
	Key    string
	Expiry time.Time
	index  int
}

type MinHeap []*Item

func (h MinHeap) Len() int { return len(h) }

func (h MinHeap) Less(i, j int) bool {
	return h[i].Expiry.Before(h[j].Expiry)
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].index = i
	h[j].index = j
}

func (h *MinHeap) Push(x interface{}) {
	n := len(*h)
	item := x.(*Item)
	item.index = n
	*h = append(*h, item)
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*h = old[0 : n-1]
	return item
}

func (h *MinHeap) Peek() *Item {
	if len(*h) == 0 {
		return nil
	}
	return (*h)[0]
}