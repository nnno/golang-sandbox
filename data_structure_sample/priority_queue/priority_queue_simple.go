// container/heapを用いない実装
package main

import (
	"fmt"
)

type MaxHeap struct {
	elements []int
}

func NewMaxHeap() *MaxHeap {
	return &MaxHeap{
		elements: []int{},
	}
}

func parentIndex(i int) int {
	return (i - 1) / 2
}

func leftChildIndex(i int) int {
	return 2*i + 1
}

func rightChildIndex(i int) int {
	return 2*i + 2
}

func (h *MaxHeap) swap(i, j int) {
	h.elements[i], h.elements[j] = h.elements[j], h.elements[i]
}

// ヒープの最大値を取得
func (h *MaxHeap) top() (int, error) {
	if len(h.elements) == 0 {
		return 0, fmt.Errorf("heap is empty")
	}
	return h.elements[0], nil
}

func (h *MaxHeap) push(x int) {
	h.elements = append(h.elements, x)
	h.up(len(h.elements) - 1)
}

// 上方向にヒープを修正
// (push時に要素数-1を引数に呼び出される)
func (h *MaxHeap) up(index int) {
	for index != 0 {
		parent := parentIndex(index)
		if h.elements[index] > h.elements[parent] {
			h.swap(index, parent)
			index = parent
		} else {
			break
		}
	}
}

// ヒープから最大値を取り出し、削除する
func (h *MaxHeap) pop() (int, error) {
	if len(h.elements) == 0 {
		return 0, fmt.Errorf("heap is empty")
	}
	root := h.elements[0]
	last := h.elements[len(h.elements)-1]
	h.elements = h.elements[:len(h.elements)-1]
	if len(h.elements) > 0 {
		h.elements[0] = last
		h.down(0)
	}
	return root, nil
}

// 下方向にヒープを修正
func (h *MaxHeap) down(index int) {
	lastIndex := len(h.elements) - 1
	for leftChildIndex(index) <= lastIndex {
		left := leftChildIndex(index)
		right := rightChildIndex(index)
		larger := left
		if right <= lastIndex && h.elements[right] > h.elements[left] {
			larger = right
		}
		if h.elements[index] < h.elements[larger] {
			h.swap(index, larger)
			index = larger
		} else {
			break
		}
	}
}

func main() {
	heap := NewMaxHeap()

	heap.push(3)
	heap.push(1)
	heap.push(5)
	heap.push(3)
	heap.push(2)
	heap.push(0)
	fmt.Println("Heap elements:", heap.elements)
	fmt.Println("----")

	maxValue, _ := heap.top()
	fmt.Println("Maximum value:", maxValue)
	fmt.Println("----")

	for len(heap.elements) > 0 {
		maxValue, _ := heap.pop()
		fmt.Println("Popped maximum value:", maxValue)
		fmt.Println("Heap elements:", heap.elements)
	}
}
