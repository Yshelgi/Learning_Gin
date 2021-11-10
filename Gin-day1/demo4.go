package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var x, y int64
var l sync.Mutex
var wg sync.WaitGroup

func mutexAdd() {
	l.Lock()
	x++
	l.Unlock()
	wg.Done()
}

func atomicAdd() {
	atomic.AddInt64(&y, 1)
	wg.Done()
}

func test1() {
	start := time.Now()
	for i := 0; i < 1000000; i++ {
		wg.Add(1)
		// go add()       // 普通版add函数 不是并发安全的
		go mutexAdd() // 加锁版add函数 是并发安全的，但是加锁性能开销大
		//go atomicAdd() // 原子操作版add函数 是并发安全，性能优于加锁版
	}
	wg.Wait()
	end := time.Now()
	fmt.Println(x)
	fmt.Println(end.Sub(start))
}

func test2() {
	start := time.Now()
	for i := 0; i < 1000000; i++ {
		wg.Add(1)
		// go add()       // 普通版add函数 不是并发安全的
		//go mutexAdd()  // 加锁版add函数 是并发安全的，但是加锁性能开销大
		go atomicAdd() // 原子操作版add函数 是并发安全，性能优于加锁版
	}
	wg.Wait()
	end := time.Now()
	fmt.Println(y)
	fmt.Println(end.Sub(start))
}

func main() {
	test2()
	test1()
}
