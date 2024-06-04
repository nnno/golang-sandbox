// 双方向リスト(Doubly Linked List)
//
// 双方向リストは、各ノードが次と前のノードへの参照（ポインタ）を持つリストです。この構造により、前後両方向にリストをたどることができます。
//
// 特徴
//
//	各ノード: データと次および前のノードへのポインタを持ちます。
//	前後両方向のトラバース: ノードを前後両方向にたどることができます。
//	任意の位置への挿入・削除が効率的: 挿入および削除操作はO(1)で行えますが、これは対象のノードへのポインタを持っている場合です。
//
// 操作の計算量
//
//	挿入・削除: O(1)（対象ノードへのポインタを持っている場合）
//	検索・ランダムアクセス: O(n)（ノードを順にたどる必要があるため）
package main

import (
	"container/list"
	"fmt"
)

func main() {
	// 新しいリストの作成
	ls := list.New()

	// 末尾に要素を追加
	ls.PushBack(3)
	fmt.Println("List after PushBack(3):")
	for e := ls.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value, " ")
	}
	fmt.Println()

	// 先頭に要素を追加
	it := ls.PushFront(2)
	fmt.Println("List after PushFront(2):")
	for e := ls.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value, " ")
	}
	fmt.Println()

	// 任意の位置に要素を挿入
	ls.InsertAfter(1, it)
	fmt.Println("List after InsertAfter(1, it):")
	for e := ls.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value, " ")
	}
	fmt.Println()

	// 任意の位置の要素を削除
	ls.Remove(it)
	fmt.Println("List after Remove(it):")
	for e := ls.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value, " ")
	}
	fmt.Println()
}
