package main

type calculateFunc func(int, int) int

var (
	operations = map[string]calculateFunc{
		"+":  add,
		"-":  sub,
		"*":  multi,
		"/":  div,
		"<<": func(a, b int) int { return a << b },
		">>": func(a, b int) int { return a >> b },
	}
)

func calculate(a, b int, operation string) int {
	switch operation {
	case "+":
		return add(a, b)
	case "-":
		return sub(a, b)
	case "*":
		return multi(a, b)
	case "/":
		return div(a, b)
	default:
		panic("operation not supported")
	}
}

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func multi(a, b int) int {
	return a * b
}

func div(a, b int) int {
	if b == 0 {
		panic("divide by zero")
	}
	return a / b
}

func calculateWithMap(a, b int, opString string) int {
	if operation, ok := operations[opString]; ok {
		return operation(a, b)
	}
	panic("operation not supported")
}
