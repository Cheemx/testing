package main

// Writing functions that take interface{} (any)
// can be extremely challenging and bug-prone
// because we lost our constraints, and we have no
// information at compile time as to what kinds of data we're dealing with.

// Generics offer us a way to make abstractions (like interfaces)
// by letting us describe our constraints.
// They allow us to write functions that have a similar level of flexibility
// that interface{} offers but retain type-safety and
// provide a better developer experience for callers.

type Stack[T any] struct {
	values []T
}

func (s *Stack[T]) Push(value T) {
	s.values = append(s.values, value)
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.values) == 0
}

func (s *Stack[T]) Pop() (T, bool) {
	if s.IsEmpty() {
		var zero T
		return zero, false
	}

	index := len(s.values) - 1
	el := s.values[index]
	s.values = s.values[:index]
	return el, true
}

// In particular using interface{} makes your code
//
// 1. Less Safe: requires more error handling
// 2. Less Expressive: tells you nothing about data
// 3. More likely to rely on reflection!!!
