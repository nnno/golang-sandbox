package main

import "time"

func main() {
	// ミューテックスを使わずに共有された変数aにアクセスしているため、データ競合が発生する
	a := 100
	go func() { a += 50 }()
	go func() { a -= 50 }()

	time.Sleep(100 * time.Millisecond)
}
