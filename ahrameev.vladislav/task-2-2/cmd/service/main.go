package main

import (
	"container/heap"
	"fmt"
)

// MaxHeap реализует интерфейс heap.Interface для среза int.
// Элемент с наибольшим значением всегда находится на вершине.
type MaxHeap []int

func (h *MaxHeap) Len() int {
	return len(*h)
}

func (h *MaxHeap) Less(i, j int) bool {
	return (*h)[i] > (*h)[j]
}

func (h *MaxHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *MaxHeap) Push(x interface{}) {
	val, ok := x.(int)
	if !ok {
		return
	}

	*h = append(*h, val)
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]

	return x
}

func main() {
	var dishesCount int
	if _, err := fmt.Scan(&dishesCount); err != nil {
		return
	}

	maxHeap := make(MaxHeap, 0, dishesCount)

	for range dishesCount {
		var value int
		if _, err := fmt.Scan(&value); err != nil {
			return
		}

		maxHeap = append(maxHeap, value)
	}

	var position int
	if _, err := fmt.Scan(&position); err != nil {
		return
	}

	heap.Init(&maxHeap)

	var result int
	targetIndex := position - 1
	for currentIndex := range position {
		val, ok := heap.Pop(&maxHeap).(int)
		if !ok {
			return
		}

		if currentIndex == targetIndex {
			result = val
		}
	}

	fmt.Println(result)
}
