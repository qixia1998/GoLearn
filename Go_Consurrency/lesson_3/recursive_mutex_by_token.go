package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Token方式的递归锁
type TokenRecursiveMutex struct {
	sync.Mutex
	token 		int64 // 当前持有锁的goroutine token
	recursion 	int32 // 这个goroutine 重入的次数
}

// 请求锁，需要传入token
func (m *TokenRecursiveMutex) Lock(token int64) {
	// 如果传入的token和持有锁的token一致，此时递归加锁，重入次数加1
	if atomic.LoadInt64(&m.token) == token {
		m.recursion++
		return
	}
	m.Mutex.Lock() // 传入的token不一致，说明不是递归调用
	// 抢到锁之后记录这个token
	atomic.StoreInt64(&m.token, token)
	m.recursion = 1
}

// 释放锁
func (m *TokenRecursiveMutex) Unlock(token int64) {
	if atomic.LoadInt64(&m.token) != token { // 释放其它token持有的锁
		panic(fmt.Sprintf("wrong the owner(%d):%d", m.token, token))
	}
	m.recursion-- // 当前持有这个锁的token释放锁
	if m.recursion != 0 { // 还没有回退到最初的递归调用
		return
	}
	atomic.StoreInt64(&m.token, 0) // 没有递归调用了，释放锁
	m.Mutex.Unlock()	
}

func main() {
	r := &TokenRecursiveMutex{}
	StartTokenLayer(r, 1024)
}

func StartTokenLayer(r *TokenRecursiveMutex, token int64) {
	r.Lock(token)
	fmt.Println("开始")
	TwoTokenLayer(r, token)
	r.Unlock(token)
}

func TwoTokenLayer(r *TokenRecursiveMutex, token int64) {
	r.Lock(token)
	fmt.Println("进入第二层")
	ThreeTokenLayer(r, token)
	r.Unlock(token)
}

func ThreeTokenLayer(r *TokenRecursiveMutex, token int64) {
	r.Lock(token)
	fmt.Println("进入第三层")
	r.Unlock(token)
}