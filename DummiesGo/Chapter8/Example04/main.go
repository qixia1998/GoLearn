package main

import (
	"fmt"
	"sort"
)

// 在Go中使用结构体和Maps
// 创建结构映射
type dob struct {
	day   int
	month int
	year  int
}

type people struct {
	name  string
	email string
	dob   dob
}

var members map[int]people

func main() {
	members = make(map[int]people)

	members[1] = people{
		name:  "Mary Smith",
		email: "marysmith@example.com",
		dob: dob{
			day:   17,
			month: 3,
			year:  1990,
		},
	}
	members[2] = people{
		name:  "John Smith",
		email: "johnsmith@example.com",
		dob: dob{
			day:   9,
			month: 12,
			year:  1988,
		},
	}
	members[3] = people{
		name:  "Janet Doe",
		email: "janetdoe@example.com",
		dob: dob{
			day:   1,
			month: 12,
			year:  1988,
		},
	}
	members[4] = people{
		name:  "Adam Jones",
		email: "adamjones@example.com",
		dob: dob{
			day:   19,
			month: 8,
			year:  2001,
		},
	}

	for k, v := range members {
		fmt.Println(k, v.name, v.email, v.dob)
	}

	// 获取成员中的所有keys
	var keys []int
	for k := range members {
		keys = append(keys, k)
	}

	// 按升序对keys 进行排序
	sort.Ints(keys)

	// 将成员中的值复制到切片中
	var sliceMembers []people
	for _, k := range keys {
		sliceMembers = append(sliceMembers, members[k])
	}

	// func SliceStable(slice interface{},
	//                 less func(i, j int) bool)

	sort.SliceStable(sliceMembers, func(i, j int) bool {
		// compare year
		if sliceMembers[i].dob.year !=
			sliceMembers[j].dob.year {
			return sliceMembers[i].dob.year <
				sliceMembers[j].dob.year
		}

		// compare month
		if sliceMembers[i].dob.month !=
			sliceMembers[j].dob.month {
			return sliceMembers[i].dob.month <
				sliceMembers[j].dob.month
		}

		// compare day
		return sliceMembers[i].dob.day <
			sliceMembers[j].dob.day
	})

	for _, v := range sliceMembers {
		fmt.Println(v.name, v.email, v.dob)
	}

	sort.SliceStable(sliceMembers, func(i, j int) bool {
		return sliceMembers[i].dob.year <
			sliceMembers[j].dob.year
	})
	for _, v := range sliceMembers {
		fmt.Println(v.name, v.email, v.dob)
	}

	sort.Slice(sliceMembers, func(i, j int) bool {
		return sliceMembers[i].dob.year <
			sliceMembers[j].dob.year
	})
	for _, v := range sliceMembers {
		fmt.Println(v.name, v.email, v.dob)
	}
}
