package main

import (
	"fmt"
	"hw1/dict"
	"hw1/queue"
	"hw1/stack"
)

func main() {
	fmt.Println("STACK")
	s := stack.NewStack()

	fmt.Printf("stack is empty: %t\n", s.IsEmpty())
	s.Push(10)
	s.Push(20)
	s.Push(30)
	fmt.Printf("pushed 10, 20, 30\n")

	top, _ := s.Peek()
	fmt.Printf("peek: %d,  size: %d\n", top, s.Size())

	s.Pop()
	top, _ = s.Peek()
	fmt.Printf("after pop — peek: %d, size: %d\n", top, s.Size())

	// test stack is empty after popping
	s.Pop()
	s.Pop()
	fmt.Printf("after 3 pops - empty: %t\n", s.IsEmpty())

	fmt.Printf("\nQUEUE\n")
	q := queue.NewQueue()

	fmt.Printf("empty: %t, size: %d\n", q.IsEmpty(), q.Size())

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	fmt.Println("queued: 1, 2, 3")
	fmt.Printf("size: %d\n", q.Size())

	q.Dequeue()
	fmt.Printf("one dequeue — size: %d\n", q.Size())

	q.Dequeue()
	q.Dequeue()
	fmt.Printf("after dequeueing all — empty: %t\n", q.IsEmpty())

	fmt.Printf("\nHashMap\n")
	hm := dict.NewDict()

	hm.Set("alice", 30)
	hm.Set("bob", 25)
	hm.Set("carol", 28)
	fmt.Println("inserted: alice=30, bob=25, carol=28")

	if val, ok := hm.Get("bob"); ok {
		fmt.Printf("get bob: %d\n", val)
	}

	hm.Set("bob", 99)
	if val, ok := hm.Get("bob"); ok {
		fmt.Printf("overwrite bob: %d\n", val)
	}

	if _, ok := hm.Get("dave"); !ok {
		fmt.Println("get dave: not found")
	}

	hm.Delete("alice")
	if _, ok := hm.Get("alice"); !ok {
		fmt.Println("after delete — alice: not found")
	}
}
