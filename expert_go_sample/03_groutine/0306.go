// go routine
package main

import (
	"log"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			do(n)
		}(i)
	}

	wg.Wait() // 全てのgoroutineが終了するまで待つ
}

func do(n int) {
	time.Sleep(2 * time.Second)
	log.Printf("%d is called\n", n)
}
