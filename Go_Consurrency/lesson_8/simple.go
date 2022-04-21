package main

import (
	"sync"
)

type OnceTest struct {
	m sync.Mutex
}

func (o *OnceTest) doSlow() {
	o.m.Lock()
	defer o.m.Unlock()

	// 这里更新的o指针的值！！！，会导致上一行Unlock出错
	*o = OnceTest{}
}

func main() {
	var once OnceTest
	once.doSlow()
}
