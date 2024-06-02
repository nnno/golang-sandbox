package main

import "fmt"

// Stack スタックの構造体を定義
type Stack struct {
	elements []int
}

// Push 要素をスタックに追加
func (s *Stack) Push(x int) {
	s.elements = append(s.elements, x)
}

// Top スタックの先頭要素を参照
func (s *Stack) Top() (int, error) {
	if len(s.elements) == 0 {
		return 0, fmt.Errorf("stack is empty")
	}
	return s.elements[len(s.elements)-1], nil
}

// Pop スタックの先頭要素を削除
func (s *Stack) Pop() error {
	if len(s.elements) == 0 {
		return fmt.Errorf("stack is empty")
	}
	s.elements = s.elements[:len(s.elements)-1]
	return nil
}

// Empty スタックが空かどうかをチェック
func (s *Stack) Empty() bool {
	return len(s.elements) == 0
}

func main() {
	// スタックの作成
	st := Stack{}

	// 要素をスタックに追加
	st.Push(3)
	st.Push(5)
	st.Push(7)
	fmt.Println("Stack after pushing elements:", st.elements)

	// スタックの先頭要素を参照
	top, err := st.Top()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Top element:", top)
	}

	// スタックの先頭要素を削除
	err = st.Pop()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Stack after popping:", st.elements)
	}

	// スタックが空かどうかをチェック
	if st.Empty() {
		fmt.Println("Stack is empty")
	} else {
		fmt.Println("Stack is not empty")
	}
}
