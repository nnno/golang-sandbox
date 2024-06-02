package main

import "fmt"

func main() {
	// キューの作成 (Queue構造体を使わずにスライスを使ってキューを実現)
	q := make([]int, 0)

	// 末尾に要素を追加
	q = append(q, 3)
	fmt.Println("Queue after pushing 3:", q)

	q = append(q, 6)
	q = append(q, 10)
	q = append(q, 1)
	fmt.Println("Queue after pushing 6, 10, 1:", q)

	// 先頭要素を参照
	front := q[0]
	fmt.Println("Front element:", front)

	// 先頭要素を削除
	x := q[0]
	q = q[1:] // 先頭要素を削除
	fmt.Println("Pop element:", x)
	fmt.Println("Queue after popping:", q)
}
