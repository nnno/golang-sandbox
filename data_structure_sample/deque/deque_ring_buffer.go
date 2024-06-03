/*
リングバッファ(循環バッファ)を使ったDeque(両端キュー)の実装

* O(1)で先頭/末尾への追加, 削除
* リングバッファを使っているため, メモリの再確保が発生しない
* 途中への挿入はO(n)のため、基本的に行わない
*/
package main

import "fmt"

type DequeRing struct {
	elements []int
	front    int
	rear     int
	cnt      int // データの数
	capacity int // 配列のサイズ
}

// NewDequeRing Dequeの生成
func NewDequeRing(capacity int) *DequeRing {
	return &DequeRing{
		elements: make([]int, capacity),
		front:    0,
		rear:     0,
		cnt:      0,
		capacity: capacity,
	}
}

// IsFull データが満杯かどうか
func (d *DequeRing) IsFull() bool {
	return d.cnt == d.capacity
}

// IsEmpty データが空かどうか
func (d *DequeRing) isEmpty() bool {
	return d.cnt == 0
}

// PushFront 先頭に要素を追加
func (d *DequeRing) PushFront(x int) error {
	if d.IsFull() {
		return fmt.Errorf("deque is full")
	}
	d.front = (d.front - 1 + d.capacity) % d.capacity
	d.elements[d.front] = x
	d.cnt++
	return nil
}

// PushBack 末尾に要素を追加
func (d *DequeRing) PushBack(x int) error {
	if d.IsFull() {
		return fmt.Errorf("deque is full")
	}
	d.elements[d.rear] = x
	d.rear = (d.rear + 1) % d.capacity
	d.cnt++
	return nil
}

// PopFront 先頭要素を削除
func (d *DequeRing) PopFront() (int, error) {
	if d.isEmpty() {
		return 0, fmt.Errorf("deque is empty")
	}
	x := d.elements[d.front]
	d.front = (d.front + 1) % d.capacity
	d.cnt--
	return x, nil
}

// PopBack 末尾要素を削除
func (d *DequeRing) PopBack() (int, error) {
	if d.isEmpty() {
		return 0, fmt.Errorf("deque is empty")
	}
	d.rear = (d.rear - 1 + d.capacity) % d.capacity
	x := d.elements[d.rear]
	d.cnt--
	return x, nil
}

// Front 先頭要素を取得
func (d *DequeRing) Front() (int, error) {
	if d.isEmpty() {
		return 0, fmt.Errorf("deque is empty")
	}
	return d.elements[d.front], nil
}

// Back 末尾要素を取得
func (d *DequeRing) Back() (int, error) {
	if d.isEmpty() {
		return 0, fmt.Errorf("deque is empty")
	}
	return d.elements[(d.rear-1+d.capacity)%d.capacity], nil
}

// At ランダムアクセス
func (d *DequeRing) At(index int) (int, error) {
	if index < 0 || index >= d.cnt {
		return 0, fmt.Errorf("index out of range")
	}
	actualIndex := (d.front + index) % d.capacity
	return d.elements[actualIndex], nil
}

func main() {
	deq := NewDequeRing(5)

	_ = deq.PushBack(1)
	_ = deq.PushBack(2)
	_ = deq.PushFront(0)
	_ = deq.PushFront(8)
	fmt.Println("Deque:", deq.elements)
	for i := 0; i < deq.cnt; i++ {
		x, _ := deq.At(i)
		fmt.Println("At:", x)
	}

	front, _ := deq.Front()
	fmt.Println("Front element:", front)

	back, _ := deq.Back()
	fmt.Println("Back element:", back)

	x, _ := deq.PopFront()
	fmt.Println("Deque PopFront:", x)
	fmt.Println("Deque after PopFront:", deq.elements, " front:", deq.front, " rear:", deq.rear)

	x, _ = deq.PopBack()
	fmt.Println("Deque PopBack:", x)
	fmt.Println("Deque after PopFront:", deq.elements, " front:", deq.front, " rear:", deq.rear)

	for i := 0; i < deq.cnt; i++ {
		x, _ = deq.At(i)
		fmt.Println("At:", x)
	}
}
