package pkg

import "testing"

func TestTodoWrite(t *testing.T) {
	todo := Todo{
		Db: &Db{
			AuthorizationFn: func() bool {
				return true
			},
		},
	}
	todo.Write("Hello")
	if todo.Text != "Hello" {
		t.Errorf("Expected 'Hello' but got %v\n", todo.Text)
	}
	todo.Append(" World")
	if todo.Text != "Hello World" {
		t.Errorf("Expected 'Hello World' but got %v\n", todo.Text)
	}
}
