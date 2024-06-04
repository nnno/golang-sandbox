// 集合 (Set)
// 重複しない要素を格納し、要素の挿入、削除、検索を効率的に行うデータ構造
//
// golang.org/x/exp/slicesパッケージ等でスライスを常にソートしながらの実装と比べると、
// mapの場合、挿入、削除、検索がO(1)で行えるが、メモリ効率が悪い
// sliceの場合, メモリ効率は良いが、時間効率が低くなる
package main

import "fmt"

type Set struct {
	elements map[int]struct{}
}

func NewSet() *Set {
	return &Set{
		elements: make(map[int]struct{}),
	}
}

// Insert 要素の追加
func (s *Set) Insert(x int) {
	s.elements[x] = struct{}{}
}

// Erase 要素の削除
func (s *Set) Erase(x int) {
	delete(s.elements, x)
}

// Find 要素の検索
func (s *Set) Find(x int) bool {
	_, exists := s.elements[x]
	return exists
}

// Size 要素数の取得
func (s *Set) Size() int {
	return len(s.elements)
}

func main() {
	s := NewSet()

	s.Insert(100)
	fmt.Println(s)
	s.Insert(2)
	fmt.Println(s)
	s.Insert(3)
	fmt.Println(s)
	fmt.Println("Set contains 2:", s.Find(2))
	fmt.Println("Set contains 3:", s.Find(3))
	fmt.Println("Set contains 4:", s.Find(4))
	s.Erase(2)
	fmt.Println("Set contains 2 after erase:", s.Find(2))
	fmt.Println("Set size:", s.Size())
}
