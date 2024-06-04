package main

import (
	"fmt"
	"golang.org/x/exp/slices"
)

func main() {
	// 集合の作成
	var s []int

	// 2を挿入
	s = insert(s, 2)
	fmt.Println("After insertion:", s)

	s = insert(s, 100)
	s = insert(s, 3)
	s = insert(s, 100) // 重複しているので挿入されない
	fmt.Println("After insertion:", s)

	// 2を探す
	found := find(s, 2)
	fmt.Println("Find 2:", found)

	// 2を削除
	s = erase(s, 2)
	fmt.Println("After deletion:", s)
}

// 挿入操作
func insert(s []int, x int) []int {
	if !slices.Contains(s, x) {
		s = append(s, x)
		slices.Sort(s)
	}
	return s
}

// 検索操作
func find(s []int, x int) bool {
	return slices.Contains(s, x)
}

// 削除操作
func erase(s []int, x int) []int {
	index := slices.Index(s, x)
	if index != -1 {
		s = append(s[:index], s[index+1:]...)
	}
	return s
}
