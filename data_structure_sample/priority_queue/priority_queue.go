// priority queue: 優先度付きキュー
// まずは、container/heapを用いた実装
package main

import (
	"container/heap"
	"fmt"
)

// IntHeap 定義
type IntHeap []int

func (h *IntHeap) Len() int           { return len(*h) }
func (h *IntHeap) Less(i, j int) bool { return (*h)[i] > (*h)[j] } // 最大ヒープの場合
func (h *IntHeap) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	h := &IntHeap{}
	heap.Init(h)

	heap.Push(h, 3)
	heap.Push(h, 1)
	heap.Push(h, 5)

	fmt.Printf("maximum: %d\n", (*h)[0]) // 最大値を取得

	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h))
	}
	fmt.Println()
}
