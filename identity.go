// Reference: https://curiosity-driven.org/monads-in-javascript#:~:text=version%2036).-,Identity%20monad,-The%20identity%20monad

package main

import "fmt"

type Identity[T any] struct {
	data T
}

func NewIdentity[T any](val T) Identity[T] {
	return Identity[T]{data: val}
}

func (instance Identity[T]) bind(transform func(val T) Identity[T]) Identity[T] {
	return transform(instance.data)
}

func IdentityExample() {
	result := NewIdentity[int](6).bind(func(val int) Identity[int] {
		return NewIdentity[int](5).bind(func(val2 int) Identity[int] {
			return NewIdentity[int](val2 + val)
		})
	})
	fmt.Println("Identity example result:", result.data)
}
