package go_timewheel

import "context"

type Task func(ctx context.Context)

// todo 实现插入，删除以及，操作加锁
type twoLinkedList struct {
	next, previous *twoLinkedList
	currentTask    Task
	totalSize      int
}

type ListWithRoot struct {
	root *twoLinkedList
}
