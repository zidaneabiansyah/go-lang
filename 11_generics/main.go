package main

import (
	"fmt"
)

type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64 |
		~string
}

func PrintValue[T any](value T) {
	fmt.Printf("Nilai: %v, Tipe: %T\n", value, value)
}

func Min[T Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func Max[T Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

type Stack[T any] struct {
	items []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item, true
}

func (s *Stack[T]) Peek() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}
	return s.items[len(s.items)-1], true
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

func FindIndex[T comparable](slice []T, target T) int {
	for i, v := range slice {
		if v == target {
			return i
		}
	}
	return -1
}

func Map[T, U any](input []T, fn func(T) U) []U {
	result := make([]U, len(input))
	for i, v := range input {
		result[i] = fn(v)
	}
	return result
}

type Numeric interface {
	~int | ~int64 | ~float64
}

func Sum[T Numeric](values []T) T {
	var total T
	for _, v := range values {
		total += v
	}
	return total
}

func main() {
	fmt.Println("GENERIC FUNCTION [T any]")
	PrintValue(42)
	PrintValue("halo")
	PrintValue(3.14)
	PrintValue(true)

	fmt.Println()
	fmt.Println("GENERIC DENGAN CONSTRAINT [T Ordered]")
	fmt.Printf("Min(10, 20) = %d\n", Min(10, 20))
	fmt.Printf("Min(3.5, 2.1) = %.1f\n", Min(3.5, 2.1))
	fmt.Printf("Min(\"abc\", \"xyz\") = %s\n", Min("abc", "xyz"))
	fmt.Printf("Max(10, 20) = %d\n", Max(10, 20))

	fmt.Println()
	fmt.Println("GENERIC STRUCT Stack[T any]")
	intStack := NewStack[int]()
	intStack.Push(10)
	intStack.Push(20)
	intStack.Push(30)
	fmt.Printf("Peek: %d\n", must(intStack.Peek()))
	fmt.Printf("Pop: %d\n", must(intStack.Pop()))
	fmt.Printf("Pop: %d\n", must(intStack.Pop()))
	fmt.Printf("IsEmpty? %v\n", intStack.IsEmpty())
	fmt.Printf("Pop: %d\n", must(intStack.Pop()))
	fmt.Printf("IsEmpty? %v\n", intStack.IsEmpty())

	stringStack := NewStack[string]()
	stringStack.Push("apple")
	stringStack.Push("banana")
	fmt.Printf("Pop string: %s\n", must(stringStack.Pop()))

	fmt.Println()
	fmt.Println("GENERIC DENGAN [T comparable]")
	numbers := []int{10, 20, 30, 40, 50}
	fmt.Printf("Index of 30: %d\n", FindIndex(numbers, 30))
	fmt.Printf("Index of 99: %d\n", FindIndex(numbers, 99))

	names := []string{"Alice", "Bob", "Charlie"}
	fmt.Printf("Index of Bob: %d\n", FindIndex(names, "Bob"))

	fmt.Println()
	fmt.Println("GENERIC MULTIPLE TYPE [T, U any]")
	doubled := Map(numbers, func(n int) int {
		return n * 2
	})
	fmt.Printf("Double: %v\n", doubled)

	toString := Map(numbers, func(n int) string {
		return fmt.Sprintf("angka-%d", n)
	})
	fmt.Printf("ToString: %v\n", toString)

	fmt.Println()
	fmt.Println("CUSTOM CONSTRAINT Numeric")
	fmt.Printf("Sum int: %d\n", Sum([]int{1, 2, 3, 4, 5}))
	fmt.Printf("Sum float: %.1f\n", Sum([]float64{1.5, 2.5, 3.0}))
}

func must[T any](val T, ok bool) T {
	if !ok {
		panic("gagal")
	}
	return val
}
