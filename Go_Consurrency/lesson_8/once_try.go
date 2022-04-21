package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"sync"
	"sync/atomic"
	"time"
)

// 一个功能更加强大的Once
type Once struct {
	m    sync.Mutex
	done uint32
}

// 传入的函数f有返回值error，如果初始化失败，则需返回失败的error
// Do方法会把这个error返回给调用者
func (o *Once) Do(f func() error) error {
	if atomic.LoadUint32(&o.done) == 1 { // fast path
		return nil
	}
	return o.slowDo(f)
}

// 如果还没有初始化
func (o *Once) slowDo(f func() error) error {
	o.m.Lock()
	defer o.m.Unlock()
	var err error
	if o.done == 0 { // 双检查，还没有初始化
		err = f()
		if err == nil { // 初始化成功才将标记置为已初始化
			atomic.StoreUint32(&o.done, 1)
		}
	}
	return err
}

func main() {
	urls := []string{
		"127.0.0.1:3453",
		"127.0.0.1:9002",
		"127.0.0.1:9003",
		"baidu.com:80",
	}
	var conn net.Conn
	var o Once
	count := 0
	var err error
	for _, url := range urls {
		err := o.Do(func() error {
			count++
			fmt.Printf("初始化%d次\n", count)
			conn, err = net.DialTimeout("tcp", url, time.Second)
			fmt.Println(err)
			return err
		})
		if err == nil {
			break
		}
		if count == 3 {
			fmt.Println("初始化失败，不再重试")
			break
		}
	}

	if conn != nil {
		_, _ = conn.Write([]byte("GET / HTTP/1.1\r\nHost: google.com\r\n Accept: */*\r\n\r\n"))
		_, _ = io.Copy(os.Stdout, conn)
	}
}
