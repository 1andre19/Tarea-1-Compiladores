package queue

import "errors"

const MAX_QUEUE_SIZE = 1024

type Queue struct {
	queue []int
	front int
	back  int
}

func NewQueue() *Queue {
	return &Queue{
		queue: make([]int, MAX_QUEUE_SIZE),
		front: -1,
		back:  -1,
	}
}

func (q *Queue) IsEmpty() bool {
	return q.back < 0
}

func (q *Queue) Size() int {
	if q.IsEmpty() {
		return 0
	}
	return q.back - q.front + 1
}

func (q *Queue) Enqueue(val int) error {
	if q.back >= MAX_QUEUE_SIZE-1 {
		return errors.New("Queue is full")
	}

	if q.IsEmpty() {
		q.front = 0
	}

	q.back += 1
	q.queue[q.back] = val
	return nil
}

func (q *Queue) Dequeue() error {
	if q.IsEmpty() {
		return errors.New("Queue is empty")
	}

	if q.front == q.back {
		q.front = -1
		q.back = -1
	} else {
		q.front += 1
	}
	return nil
}

func (q *Queue) Peek() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("Queue is empty")
	}
	return q.queue[q.front], nil
}
