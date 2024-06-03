// container/listパッケージを利用したDeque

package main

import (
	"container/list"
	"fmt"
)

func main() {
	// 新しいリストの作成
	deq := list.New()

	// 先頭に要素を追加
	deq.PushFront(3)
	fmt.Println("Deque after pushing 3 at front:")
	for e := deq.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value, " ")
	}
	fmt.Println()

	// 末尾に要素を追加
	deq.PushBack(0)
	fmt.Println("Deque after pushing 0 at back:")
	for e := deq.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value, " ")
	}
	fmt.Println()

	// 先頭要素を取得
	front := deq.Front().Value
	fmt.Println("Front element:", front)

	// 先頭要素を削除
	deq.Remove(deq.Front())
	fmt.Println("Deque after popping front:")
	for e := deq.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value, " ")
	}
	fmt.Println()

	// 末尾要素を取得
	back := deq.Back().Value
	fmt.Println("Back element:", back)

	// 末尾要素を削除
	deq.Remove(deq.Back())
	fmt.Println("Deque after popping back:")
	for e := deq.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value, " ")
	}
	fmt.Println()
}
