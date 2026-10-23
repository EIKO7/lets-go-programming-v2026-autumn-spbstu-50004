package main

import "fmt"

type MaxHeap []int

func (h MaxHeap) Len() int {
	return len(h)
}

func (h MaxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func main() {
	var n int
	fmt.Scan(&n)

	heap := make(MaxHeap, 0, n)
	for i := 0; i < n; i++ {
		var value int
		fmt.Scan(&value)
		heap = append(heap, value)
	}

	var k int
	fmt.Scan(&k)
}
