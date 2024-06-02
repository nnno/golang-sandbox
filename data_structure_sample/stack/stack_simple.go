package main

import "fmt"

// stack.goの用に構造体を定義しても良いが、
// Last In First Out (LIFO)するだけなら標準のスライスでも可能

func main() {
	// スタックの作成
	st := make([]int, 0)

	// 要素をスタックに追加
	st = append(st, 3)
	st = append(st, 5)
	st = append(st, 7)
	fmt.Println("Stack after pushing elements:", st)

	// スタックの先頭要素を参照
	top := st[len(st)-1]
	fmt.Println("Top element:", top)

	// スタックの先頭要素を削除
	x := st[len(st)-1]
	st = st[:len(st)-1]
	fmt.Println("Pop value from Stack:", x)
	fmt.Println("Stack after popping:", st)

}
