package main

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/petermattis/goid"
)

// 简单获取 gid
func GoID() int {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	// 得到 id 字符串
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine"))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id :%v", err))
	}
	return id
}

// RecursiveMutex 包装一个Mutex，实现可重入
type RecursiveMutex struct {
	sync.Mutex
	owner 		int64 // 当前持有锁的goroutine id
	recursion 	int32 // 这个goroutine 重入的次数
}

func (m *RecursiveMutex) Lock() {
	gid := goid.Get()
	// 如果当前持有锁的goroutine就是这次调用的goroutine，说明是重入
	if atomic.LoadInt64(&m.owner) == gid {
		m.recursion++
		return
	}
	m.Mutex.Lock()
	// 获得锁的goroutine第一次调用，记录下它的goroutine id，调用次数加1
	atomic.StoreInt64(&m.owner, gid)
	m.recursion = 1
}

func (m *RecursiveMutex) Unlock() {
	gid := goid.Get()
	// 非持有锁的goroutine尝试释放锁，错误的使用
	if atomic.LoadInt64(&m.owner) != gid {
		panic(fmt.Sprintf("wrong the owner (%d):%d!", m.owner, gid))
	}
	// 调用次数减1
	m.recursion--
	if m.recursion != 0 { // 如果这个goroutine还没完全释放，则直接返回
		return
	}
	// 此goroutine最后一次调用，需要释放锁
	atomic.StoreInt64(&m.owner, -1)
	m.Mutex.Unlock()
}

// 通过runtime.Stack 获取运行gID
func GoId() int {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	// 得到id字符串
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get g id:%v", err))
	}
	return id
}

func main() {
	r := &RecursiveMutex{}
	StartLayer(r)
}

func StartLayer(r *RecursiveMutex) {
	r.Lock()
	fmt.Println("开始")
	TwoLayer(r)
	r.Unlock()
}

func TwoLayer(r *RecursiveMutex) {
	r.Lock()
	fmt.Println("进入第二层")
	ThreeLayer(r)
	r.Unlock()
}

func ThreeLayer(r *RecursiveMutex) {
	r.Lock()
	fmt.Println("进入第三层")
	r.Unlock()
}