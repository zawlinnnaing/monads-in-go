package main

import "fmt"

type List[T any] []T

func (list List[T]) bind(transform func(val T) List[T]) List[T] {
	var transformedList []T
	for _, item := range list {
		transformedList = append(transformedList, transform(item)...)
	}
	return transformedList
}

func ListExample() {
	result := List[int]{
		1, 2, 3,
	}.bind(func(val int) List[int] {
		return List[int]{
			4, 5, 6,
		}.bind(func(val2 int) List[int] {
			return List[int]{val + val2}
		})
	})

	fmt.Println("List monad example", result)
}
