package main

import "fmt"

// Deque Deque構造体を定義
type Deque struct {
	elements []int
}

// PushFront 先頭に要素を追加
func (d *Deque) PushFront(x int) {
	d.elements = append([]int{x}, d.elements...)
}

// PushBack 末尾に要素を追加
func (d *Deque) PushBack(x int) {
	d.elements = append(d.elements, x)
}

// Front 先頭要素を取得
func (d *Deque) Front() (int, error) {
	if len(d.elements) == 0 {
		return 0, fmt.Errorf("deque is empty")
	}
	return d.elements[0], nil
}

// Back 末尾要素を取得
func (d *Deque) Back() (int, error) {
	if len(d.elements) == 0 {
		return 0, fmt.Errorf("deque is empty")
	}
	return d.elements[len(d.elements)-1], nil
}

// PopFront 先頭要素を削除
func (d *Deque) PopFront() error {
	if len(d.elements) == 0 {
		return fmt.Errorf("deque is empty")
	}
	d.elements = d.elements[1:]
	return nil
}

// PopBack 末尾要素を削除
func (d *Deque) PopBack() error {
	if len(d.elements) == 0 {
		return fmt.Errorf("deque is empty")
	}
	d.elements = d.elements[:len(d.elements)-1]
	return nil
}

// At ランダムアクセス
func (d *Deque) At(index int) (int, error) {
	if index < 0 || index >= len(d.elements) {
		return 0, fmt.Errorf("index out of range")
	}
	return d.elements[index], nil
}

// Empty Dequeが空かどうかをチェック
func (d *Deque) Empty() bool {
	return len(d.elements) == 0
}

func main() {
	// Dequeの作成
	deq := Deque{}

	// 先頭に要素を追加
	deq.PushFront(3)
	deq.PushBack(0)
	deq.PushBack(5)
	fmt.Println("Deque:", deq.elements)

	// ランダムアクセス
	value, err := deq.At(1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Element at index 1:", value)
	}
}
