package utils

type Stack[T any] struct {
  array []T
}

const STACK_SIZE_CHUNKS = 16

func StackFactory[T any]() Stack[T] {
  array := make([]T, STACK_SIZE_CHUNKS)

  return Stack[T]{
    array: array,
  }
}

func (s *Stack[T]) IsEmpty() bool {
  return len(s.array) == 0
}

func (s *Stack[T]) Push(item T) {
  s.array = append(s.array, item)
}

func (s *Stack[T]) Pop() T {
  topElement := s.array[len(s.array) - 1]
  s.array = s.array[:len(s.array) - 1]
  return topElement
}

func (s *Stack[T]) Peek() T {
  return s.array[len(s.array) - 1]
}
