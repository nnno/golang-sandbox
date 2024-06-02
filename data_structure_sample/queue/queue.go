package main

// Go言語でのキュー操作
// Go言語には標準でキューが実装されていないため、スライスを使ってキューを実現する
// 但し、スライスは先頭に要素を追加する場合、メモリの再確保が発生して処理速度が遅くなってしまう

import "fmt"

type Queue struct {
	elements []int
}

// Push キューの末尾に要素を追加
func (q *Queue) Push(x int) {
	q.elements = append(q.elements, x)
}

// Front キューの先頭要素を参照
func (q *Queue) Front() (int, error) {
	if len(q.elements) == 0 {
		return 0, fmt.Errorf("queue is empty")
	}
	return q.elements[0], nil
}

// Pop キューの先頭要素を削除
func (q *Queue) Pop() error {
	if len(q.elements) == 0 {
		return fmt.Errorf("queue is empty")
	}
	q.elements = q.elements[1:]
	return nil
}

// Empty キューが空かどうかをチェック
func (q *Queue) Empty() bool {
	return len(q.elements) == 0
}

func main() {
	// キューの作成
	q := Queue{}

	// 末尾に要素を追加
	q.Push(3)
	fmt.Println("Queue after pushing 3:", q.elements)

	q.Push(6)
	q.Push(10)
	q.Push(1)
	fmt.Println("Queue after pushing 6, 10, 1:", q.elements)

	// 先頭要素を参照
	front, err := q.Front()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Front element:", front)
		fmt.Println("    Queue:", q.elements)
	}

	// 先頭要素を削除
	err = q.Pop()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Queue after popping:", q.elements)
	}

	// キューが空かどうかをチェック
	if q.Empty() {
		fmt.Println("Queue is empty")
	} else {
		fmt.Println("Queue is not empty")
	}

	// 空のキューの場合
	q2 := Queue{}
	if q2.Empty() {
		fmt.Println("Queue(q2) is empty")
	} else {
		fmt.Println("Queue(q2) is not empty")
	}
}
