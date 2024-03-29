package main

import (
	"fmt"
	"sync"
	"time"
)

// SafeCounter 的并发使用是安全的。
type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

// Inc 增加给定 key 的计数器的值。
func (c *SafeCounter) Inc(key string) {
	c.mux.Lock()
	// Lock 之后同一时刻只有一个 goroutine 能访问 c.v]
  //unlockをするまでに keyを変えられるのは1個のgoroutineだけ
	c.v[key]++
  //終わったら、unlockをして、次のgoroutineがvにアクセスできる
	c.mux.Unlock()
}

// Value 返回给定 key 的计数器的当前值。
func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()
	// Lock 之后同一时刻只有一个 goroutine 能访问 c.v
	defer c.mux.Unlock()
	return c.v[key]
}

func main() {
  //マップを初期化して、SafeCounter型のcを作成
	c := SafeCounter{v: make(map[string]int)}

  //1000回ループする
	for i := 0; i < 1000; i++ {
    //go goroutineを呼び出す
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}
