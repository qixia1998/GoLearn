package main

import "fmt"

type phoneNumber string
type age uint
type Person struct {
	name        string
	age         age
	phonenumber phoneNumber
}

func (p *Person) setPhoneNumber(s phoneNumber) {
	p.phonenumber = s
}

func (p *Person) update(name string, phonenumber phoneNumber) {
	p.name = name
	p.phonenumber = phonenumber
}

func (a age) valid() bool {
	return a < 120
}

func isValidPerson(p Person) bool {
	return p.age.valid() && p.name != ""
}

type predicate func(int) bool

//func largerThanTwo(i int) bool {
//	return i > 2
//}

func filter(is []int, p predicate) []int {
	out := []int{}
	for _, i := range is {
		if p(i) {
			out = append(out, i)
		}
	}
	return out
}

func createLargerThanPredicate(threshold int) predicate {
	return func(i int) bool {
		return i > threshold
	}
}

var (
	largerThanTwo     = createLargerThanPredicate(2)
	largerThanFive    = createLargerThanPredicate(5)
	largerThanHundred = createLargerThanPredicate(100)
)

type ConstraintChecker struct {
	largerThan  predicate
	smallerThan predicate
}

func (c ConstraintChecker) check(input int) bool {
	return c.largerThan(input) && c.smallerThan(input)
}

func main() {
	p := Person{
		name:        "John",
		phonenumber: "123",
	}

	fmt.Printf("%v\n", p)

	// functions in variables
	inlinePersonStruct := struct {
		name string
	}{
		name: "John",
	}
	fmt.Printf("%v\n", inlinePersonStruct)

	//ints := []int{1, 2, 3}
	//inlineFunction := func(i int) bool {return i > 2}
	//filter(ints, inlineFunction)

	filter([]int{1, 2, 3}, func(i int) bool {
		return i > 2
	})

	//largerThanTwo := createLargerThanPredicate(2)
	//filter(ints, largerThanTwo)
	//largerThanFive := createLargerThanPredicate(5)
	//largerThanHundred := createLargerThanPredicate(100)

	ints := []int{1, 2, 3, 6, 101}
	//predicates := []predicate{largerThanTwo, largerThanFive,
	//	largerThanHundred}
	//
	//for _, predicate := range predicates {
	//	fmt.Printf("%v\n", filter(ints, predicate))
	//}

	dispatcher := map[string]predicate{
		"2": largerThanTwo,
		"5": largerThanFive,
	}
	fmt.Printf("%v\n", filter(ints, dispatcher["2"]))

	checker := ConstraintChecker{
		largerThan: createLargerThanPredicate(2),
		smallerThan: func(i int) bool {
			return i < 10
		},
	}
	fmt.Printf("%v\n", checker.check(5))
}
