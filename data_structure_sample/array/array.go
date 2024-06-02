package main

import "fmt"

func main() {
	// 要素数nの1次元スライス
	n := 5
	v := make([]int, n)

	// 要素のランダムアクセス
	i := 2
	value := v[i]
	fmt.Println("Access element at index", i, ":", value)

	// 末尾への挿入
	x := 10
	v = append(v, x)
	fmt.Println("After appending", x, ":", v)

	// 末尾要素の削除
	v = v[:len(v)-1]
	fmt.Println("After removing last element:", v)

	// 末尾以外への値の挿入
	// スライス操作で実現できる, 挿入する位置の前後でスライスを分割し、新しい要素を追加してから結合する
	v = []int{1, 2, 3, 4, 5}
	// 挿入位置と挿入する値
	i = 2
	x = 10

	// 挿入操作
	v = append(v[:i], append([]int{x}, v[i:]...)...)

	// 結果表示
	fmt.Println("After inserting", x, "at index", i, ":", v)
}
