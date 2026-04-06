package stack

import "errors"

const MAX_STACK_SIZE = 1024

type Stack struct {
	stack []int
	top   int
}

func NewStack() *Stack {
	return &Stack{
		stack: make([]int, MAX_STACK_SIZE),
		top:   -1,
	}
}

func (s *Stack) Push(val int) error {
	if s.top >= MAX_STACK_SIZE-1 {
		return errors.New("Stack is full")
	}
	s.top += 1
	s.stack[s.top] = val
	return nil
}

func (s *Stack) Pop() error {
	if s.IsEmpty() {
		return errors.New("Stack is empty")
	}
	s.top -= 1
	return nil
}

func (s *Stack) Peek() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("Stack is empty")
	}
	return s.stack[s.top], nil
}

func (s *Stack) Size() int {
	return s.top + 1
}

func (s *Stack) IsEmpty() bool {
	return s.top == -1
}
