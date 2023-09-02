package main

import "fmt"

type Maybe[T any] interface {
	bind(func(val T) Maybe[T]) Maybe[T]
	fmt.Stringer
}

type Just[T any] struct {
	Data T
}

func NewJust[T any](val T) Just[T] {
	return Just[T]{Data: val}
}

func (instance Just[T]) bind(transform func(val T) Maybe[T]) Maybe[T] {
	return transform(instance.Data)
}

func (instance Just[T]) String() string {
	return fmt.Sprintf("%v", instance.Data)
}

type Nothing[T any] struct {
}

func (nothing Nothing[T]) bind(func(val T) Maybe[T]) Maybe[T] {
	return nothing
}

func (nothing Nothing[T]) String() string {
	return "Nothing"
}

func NewNothing[T any]() Nothing[T] {
	return Nothing[T]{}
}

func MaybeExample() {
	propagationResult := NewJust[int](6).bind(func(val int) Maybe[int] {
		return NewJust[int](5).bind(func(val2 int) Maybe[int] {
			return NewJust[int](val + val2)
		})
	})
	fmt.Println("Maybe Monad: Propagation with all just", propagationResult)

	nothingResult := NewJust[int](6).bind(func(val int) Maybe[int] {
		return NewJust[int](5).bind(func(val2 int) Maybe[int] {
			return NewNothing[int]().bind(func(val3 int) Maybe[int] {
				return NewJust[int](val + val2)
			})
		})
	})
	fmt.Println("Maybe monad: Propagation will stop:", nothingResult)
}
