package main

import (
	"reflect"
	"testing"
)

func Test_changeName(t *testing.T) {
	type args struct {
		p       *Person
		newName string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{"1", args{&Person{Name: "Test"}, "Test1"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			changeName(tt.args.p, tt.args.newName)
		})
	}
}

func Test_changeNamePure(t *testing.T) {
	type args struct {
		p       Person
		newName string
	}
	tests := []struct {
		name string
		args args
		want Person
	}{
		// TODO: Add test cases.
		{"1", args{Person{
			Name: "Test",
		}, "Test1"}, Person{Name: "Test1"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := changeNamePure(tt.args.p, tt.args.newName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("changeNamePure() = %v, want %v", got, tt.want)
			}
		})
	}
}
