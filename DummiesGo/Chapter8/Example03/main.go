package main

import (
	"fmt"
	"sort"
)

// 按值对map 中的项进行排序

type kv struct {
	key   string
	value int
}

type kvPairs []kv

var heights map[string]int

func (p kvPairs) Len() int {
	// 返回集合的长度
	return len(p)
}

func (p kvPairs) Less(i, j int) bool {
	// 表示第一个值（height）必须比第二个值更小
	return p[i].value < p[j].value
}

func (p kvPairs) Swap(i, j int) {
	// 交换集合中的项
	p[i], p[j] = p[j], p[i]
}

func main() {

	heights := make(map[string]int)
	heights["Peter"] = 170
	heights["Joan"] = 168
	heights["Jan"] = 175

	p := make(kvPairs, len(heights))
	i := 0
	for k, v := range heights {
		p[i] = kv{k, v}
		i++
	}

	fmt.Println(p)

	sort.Sort(p)

	fmt.Println(p)

	for _, v := range p {
		fmt.Println(v)
	}
}
