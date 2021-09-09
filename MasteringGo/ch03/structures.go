package main

import "fmt"

type Entry struct {
	Name 	string
	Surname string
	Year 	int
}

// 由Go初始化
func zeroS() Entry {
	return Entry{}
}

// 由 user 初始化
func initS(N, S string, Y int) Entry {
	if Y < 2000 {
		return Entry{Name: N, Surname: S, Year: 2000}
	}
	return Entry{Name: N, Surname: S, Year: Y}
}

// 由 Go 初始化 - 返回指针
func zeroPtoS() *Entry {
	t := &Entry{}
	return t
}

// 由 user 初始化 - 返回指针
func initPtoS(N, S string, Y int) *Entry {
	if len(S) == 0 {
		return &Entry{Name: N, Surname: "Unknown", Year: Y}
	}
	return &Entry{Name: N, Surname: S, Year: Y}
}

func main() {
	s1 := zeroS()
	p1 := zeroPtoS()
	fmt.Println("s1:", s1, "p1:", *p1)

	s2 := initS("Mihalis", "Tsoukalos", 2000)
	p2 := initPtoS("Mihalis", "Tsoukalos", 2000)
	fmt.Println("s2:", s2, "p2:", *p2)

	fmt.Println("Year:", s1.Year, s2.Year, p1.Year, p2.Year)

	pS := new(Entry)
	fmt.Println("pS:", pS)
}
